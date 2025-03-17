package contracts

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/kitshy/treasure-backend/bindings"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
	"github.com/kitshy/treasure-backend/database/event"
	"github.com/kitshy/treasure-backend/database/eventlog"
	"math/big"
)

func TreasureProcessor(db *database.DB, conf config.ChainConfig, fromHeight *big.Int, toHeight *big.Int) error {

	treasureManagerAbi, err := bindings.TreasureManagerMetaData.GetAbi()
	if err != nil {
		log.Error("treasure manager meta data", "err", err)
		return err
	}

	treasureManagerFilterer, err := bindings.NewTreasureManagerFilterer(conf.Contracts[0], nil)
	if err != nil {
		log.Error("new treasure manager fail", "err", err)
		return err
	}

	contractEventFilter := event.ContractEvents{ContractAddress: conf.Contracts[0]}
	contractEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		log.Error("get contract events fail", "err", err)
		return err
	}

	var txDepositTokens []eventlog.DepositTokens
	var txGrantRewardTokens = []eventlog.GrantRewardTokens{}

	for _, contractEvent := range *contractEvents {
		rlpLog := contractEvent.RlpLogs
		if contractEvent.EventSignature.String() == treasureManagerAbi.Events["DepositToken"].ID.String() {
			depositTokenEvent, err := treasureManagerFilterer.ParseDepositToken(*rlpLog)
			if err != nil {
				log.Error("parse deposit token fail", "err", err)
				return err
			}
			txDepositToken := depositTokenLogTransformDb(depositTokenEvent, big.NewInt(int64(rlpLog.BlockNumber)), contractEvent.Timestamp)
			txDepositTokens = append(txDepositTokens, txDepositToken)
		} else if contractEvent.EventSignature.String() == treasureManagerAbi.Events["WithdrawToken"].ID.String() {
			// todo
		} else if contractEvent.EventSignature.String() == treasureManagerAbi.Events["GrantRewardTokenAmount"].ID.String() {
			grantRewardTokenEvent, err := treasureManagerFilterer.ParseGrantRewardTokenAmount(*rlpLog)
			if err != nil {
				log.Error("parse withdraw token fail", "err", err)
				return err
			}
			txGrantRewardToken := grantRewardTokenLogTransformDb(grantRewardTokenEvent, big.NewInt(int64(rlpLog.BlockNumber)), contractEvent.Timestamp)
			txGrantRewardTokens = append(txGrantRewardTokens, txGrantRewardToken)
		} else if contractEvent.EventSignature.String() == treasureManagerAbi.Events["WithdrawManagerUpdate"].ID.String() {
			// todo
		}
	}
	return nil
}

func depositTokenLogTransformDb(depositTokenEvent *bindings.TreasureManagerDepositToken, blockNumber *big.Int, t uint64) eventlog.DepositTokens {
	txDepositToken := eventlog.DepositTokens{
		GUID:         uuid.New(),
		TokenAddress: depositTokenEvent.TokenAddress,
		Sender:       depositTokenEvent.Sender,
		Amount:       depositTokenEvent.Amount,
		BlockNumber:  blockNumber,
		Timestamp:    t,
	}
	return txDepositToken
}

func grantRewardTokenLogTransformDb(depositTokenEvent *bindings.TreasureManagerGrantRewardTokenAmount, blockNumber *big.Int, t uint64) eventlog.GrantRewardTokens {
	txGrantRewardToken := eventlog.GrantRewardTokens{
		GUID:         uuid.New(),
		BlockNumber:  blockNumber,
		TokenAddress: depositTokenEvent.TokenAddress,
		Granter:      depositTokenEvent.Granter,
		Amount:       depositTokenEvent.Amount,
		Timestamp:    t,
	}
	return txGrantRewardToken
}
