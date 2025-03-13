package event

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractEvents struct {
	GUID            uuid.UUID      `gorm:"primaryKey"`
	BlockHash       common.Hash    `gorm:"serializer:bytes"`
	ContractAddress common.Address `gorm:"serializer:bytes"`
	TransactionHash common.Hash    `gorm:"serializer:bytes"`
	LogIndex        uint64
	EventSignature  common.Hash `gorm:"serializer:bytes"`
	Timestamp       uint64
	RlpLogs         *types.Log `gorm:"serializer:rlp;column:rlp_bytes"`
}

type contractEventsDB struct {
	gorm *gorm.DB
}

type ContractEventsViews interface {
	QueryContractEventsByID(uuid.UUID) (*ContractEvents, error)
}

type ContractEventsDB interface {
	ContractEventsViews
	SaveContractEvents(*[]ContractEvents) error
}

func NewContractEventsDB(db *gorm.DB) ContractEventsDB {
	return &contractEventsDB{gorm: db}
}

func (c *ContractEvents) AfterFind(db *gorm.DB) error {
	c.RlpLogs.BlockHash = c.BlockHash
	c.RlpLogs.TxHash = c.TransactionHash
	c.RlpLogs.Index = uint(c.LogIndex)
	return nil
}

func ContractEventsFromLog(log *types.Log, time uint64) *ContractEvents {
	eventSig := common.Hash{}
	if len(log.Topics) > 0 {
		eventSig = log.Topics[0]
	}
	return &ContractEvents{
		GUID:            uuid.New(),
		BlockHash:       log.BlockHash,
		ContractAddress: log.Address,
		TransactionHash: log.TxHash,
		LogIndex:        uint64(log.Index),
		EventSignature:  eventSig,
		Timestamp:       time,
		RlpLogs:         log,
	}
}

func (c *contractEventsDB) SaveContractEvents(events *[]ContractEvents) error {
	result := c.gorm.CreateInBatches(&events, len(*events))
	return result.Error
}

func (c *contractEventsDB) QueryContractEventsByID(id uuid.UUID) (*ContractEvents, error) {
	contractEvents := &ContractEvents{}
	result := c.gorm.Where(&ContractEvents{GUID: id}).Take(&contractEvents)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return contractEvents, nil
}
