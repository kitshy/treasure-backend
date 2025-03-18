package call

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	common2 "github.com/kitshy/treasure-backend/common/crypto"
	"github.com/kitshy/treasure-backend/rpccall/ethcli"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/kitshy/treasure-backend/bindings"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/rpccall/txmgr"
)

type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type ContractCaller struct {
	ethClient                  *ethclient.Client
	Ctx                        context.Context
	Cfg                        *config.CallContractsConfig
	TreasureManagerContract    *bindings.TreasureManager
	RawTreasureManagerContract *bind.BoundContract
	TreasureManagerABI         *abi.ABI
	txMgr                      txmgr.TxManager
	SignerFn                   SignerFn
	PrivateKey                 *ecdsa.PrivateKey
	WalletAddress              common.Address

	cancel func()
	wg     sync.WaitGroup
	once   sync.Once
}

type ContractCallerHandler interface {
	setWithdrawManager(string) (*types.Transaction, error)
	setTokenWhiteList(string) error
	getTokenWhiteList() ([]string, error)
	UpdateGasPrice(context.Context, *types.Transaction) (*types.Transaction, error)
}

func NewContractCaller(ctx context.Context, conf *config.CallContractsConfig) (ContractCallerHandler, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	callerPrivateKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"ContractCaller", conf.Mnemonic, conf.SequencerHDPath, conf.PrivateKey, conf.Contracts[0].Hex(), conf.Password)

	var walletAddr = conf.HsmAddress
	if !conf.EnableHsm {
		walletAddr = crypto.PubkeyToAddress(callerPrivateKey.PublicKey)
	}

	ethClient, err := ethcli.EthClientWithTimeout(ctx, conf.ChainRpcUrl)
	if err != nil {
		return nil, err
	}

	treasureManagerContract, err := bindings.NewTreasureManager(conf.Contracts[0], ethClient)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.TreasureManagerMetaData.ABI,
	))
	if err != nil {
		log.Error("MtChallenger parse eigen layer contract abi fail", "err", err)
		return nil, err
	}
	treasureManagerABI, err := bindings.TreasureManagerMetaData.GetAbi()
	if err != nil {
		log.Error("MtChallenger get eigen layer contract abi fail", "err", err)
		return nil, err
	}
	rawTreasureManagerContract := bind.NewBoundContract(conf.Contracts[0], parsed, ethClient, ethClient, ethClient)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          conf.NumConfirmations,
		SafeAbortNonceTooLowCount: conf.SafeAbortNonceTooLowCount,
	}
	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, ethClient)
	return &ContractCaller{
		Cfg:                        conf,
		Ctx:                        ctx,
		TreasureManagerContract:    treasureManagerContract,
		RawTreasureManagerContract: rawTreasureManagerContract,
		TreasureManagerABI:         treasureManagerABI,
		txMgr:                      txMgr,
		PrivateKey:                 callerPrivateKey,
		WalletAddress:              walletAddr,
		cancel:                     cancel,
	}, nil
}

func (c *ContractCaller) setWithdrawManager(addr string) (*types.Transaction, error) {
	balance, err := c.ethClient.BalanceAt(c.Ctx, c.WalletAddress, nil)
	if err != nil {
		log.Error("Contract caller unable to get current balance", "err", err)
		return nil, err
	}
	log.Info("Contract wallet address balance", "balance", balance)

	nonce64, err := c.ethClient.NonceAt(c.Ctx, c.WalletAddress, nil)
	if err != nil {
		log.Error("Contract wallet unable to get current nonce", "err", err)
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)

	var opts *bind.TransactOpts
	if !c.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(c.PrivateKey, big.NewInt(int64(c.Cfg.ChainId)))
	} else {
		opts, err = common2.NewHSMTransactOpts(c.Ctx, c.Cfg.HsmAPIName, c.Cfg.HsmAddress.Hex(), big.NewInt(int64(c.Cfg.ChainId)), c.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = c.Ctx
	opts.Nonce = nonce
	opts.NoSend = true

	tx, err := c.TreasureManagerContract.SetWithdrawManager(opts, common.HexToAddress(addr))
	switch {
	case err == nil:
		return tx, nil
	case c.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("contract caller eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		tx, err1 := c.TreasureManagerContract.SetWithdrawManager(opts, common.HexToAddress(addr))
		if err1 != nil {
			return nil, err1
		}
		updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
			log.Info("Contract caller setWithdrawManager update gas price")
			return c.UpdateGasPrice(ctx, tx)
		}
		receipt, err := c.txMgr.Send(c.Ctx, updateGasPrice, c.SendTransaction)
		if err != nil {
			return nil, err
		}
		log.Info("Contract caller set withdraw manager success", "TxHash", receipt.TxHash)

		return nil, nil
	default:
		return nil, err
	}
}

func (c *ContractCaller) setTokenWhiteList(addr string) error {
	return nil
}

func (c *ContractCaller) getTokenWhiteList() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContractCaller) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	if !c.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(c.PrivateKey, big.NewInt(int64(c.Cfg.ChainId)))
	} else {
		opts, err = common2.NewHSMTransactOpts(ctx, c.Cfg.HsmAPIName, c.Cfg.HsmAddress.Hex(), big.NewInt(int64(c.Cfg.ChainId)), c.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true

	finalTx, err := c.RawTreasureManagerContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	case c.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Info("MtChallenger eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap", "txData", tx.Data())
		opts.GasTipCap = FallbackGasTipCap
		return c.RawTreasureManagerContract.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

func (c *ContractCaller) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}

func (c *ContractCaller) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), errMaxPriorityFeePerGasNotFound.Error(),
	)
}
