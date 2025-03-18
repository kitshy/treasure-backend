package config

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Migrations          string       `json:"migrations"`
	HttpService         ServerConfig `json:"http_service"`
	Chain               ChainConfig  `json:"chain"`
	MasterDB            DBConfig     `json:"master_db"`
	SlaveDB             DBConfig     `json:"slave_db"`
	SlaveDbEnable       bool
	CallContractsConfig CallContractsConfig
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DBConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type ChainConfig struct {
	ChainRpcUrl string           `json:"chain_rpc_url"`
	ChainId     uint             `json:"chain_id"`
	Contracts   []common.Address `json:"contracts"`
	StartHeight uint64           `json:"start_height"`
	BlockStep   uint64           `json:"block_step"`
}

type CallContractsConfig struct {
	// chain info same as chain config
	ChainRpcUrl               string           `json:"chain_rpc_url"`
	ChainId                   uint             `json:"chain_id"`
	Contracts                 []common.Address `json:"contracts"`
	PrivateKey                string
	Mnemonic                  string
	SequencerHDPath           string
	Password                  string
	NumConfirmations          uint64 `json:"num_confirmations"`
	SafeAbortNonceTooLowCount uint64 `json:"safe_abort_nonce_too_low_count"`
	EnableHsm                 bool
	HsmAPIName                string `json:"hsm_api_name"`
	HsmCreden                 string `json:"hsm_creden"`
	HsmAddress                common.Address
}

func NewConfig(ci *cli.Context) (*Config, error) {
	conf := Config{
		Migrations:    ci.String(MigrationsFlags.Name),
		SlaveDbEnable: ci.Bool(SlaveDbEnableFlags.Name),
		HttpService: ServerConfig{
			Host: ci.String(HttpHostFlags.Name),
			Port: ci.Int(HttpPortFlags.Name),
		},
		Chain: ChainConfig{
			ChainRpcUrl: ci.String(ChainRpcUrlFlags.Name),
			ChainId:     ci.Uint(ChainIdFlags.Name),
			Contracts:   loadContracts(ci.String(ContractsFlags.Name)),
			StartHeight: ci.Uint64(StartHeightFlags.Name),
			BlockStep:   ci.Uint64(BlockStepFlags.Name),
		},
		MasterDB: DBConfig{
			Host: ci.String(MasterDbHostFlags.Name),
			Port: ci.Int(MasterDbPortFlags.Name),
			User: ci.String(MasterDbUserFlags.Name),
			Pass: ci.String(MasterDbPasswordFlags.Name),
			Name: ci.String(MasterDbNameFlags.Name),
		},
		SlaveDB: DBConfig{
			Host: ci.String(SlaveDbHostFlags.Name),
			Port: ci.Int(SlaveDbPortFlags.Name),
			Name: ci.String(SlaveDbNameFlags.Name),
			User: ci.String(SlaveDbUserFlags.Name),
			Pass: ci.String(SlaveDbPasswordFlags.Name),
		},
		CallContractsConfig: CallContractsConfig{
			ChainRpcUrl:               ci.String(ChainRpcUrlFlags.Name),
			ChainId:                   ci.Uint(ChainIdFlags.Name),
			Contracts:                 loadContracts(ci.String(ContractsFlags.Name)),
			PrivateKey:                ci.String(PrivateKeyNameFlags.Name),
			Mnemonic:                  ci.String(MnemonicNameFlags.Name),
			SequencerHDPath:           ci.String(SequencerHDPathNameFlags.Name),
			Password:                  ci.String(PasswordNameFlags.Name),
			NumConfirmations:          ci.Uint64(NumConfirmationsNameFlags.Name),
			SafeAbortNonceTooLowCount: ci.Uint64(SafeAbortNonceTooLowCountNameFlags.Name),
			EnableHsm:                 ci.Bool(EnableHsmNameFlags.Name),
			HsmAPIName:                ci.String(HsmAPINameFlags.Name),
			HsmCreden:                 ci.String(HsmCredenNameFlags.Name),
			HsmAddress:                common.HexToAddress(ci.String(HsmAddressNameFlags.Name)),
		},
	}
	log.Info("loaded config success", "config")
	return &conf, nil
}

func loadContracts(value string) []common.Address {
	var Contracts []common.Address
	addresses := strings.Split(value, ",")
	for _, addr := range addresses {
		if !common.IsHexAddress(addr) {
			panic(errors.New("invalid contacts address"))
		}
		Contracts = append(Contracts, common.HexToAddress(addr))
	}
	return Contracts
}
