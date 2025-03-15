package eventlog

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type DepositTokens struct {
	GUID         uuid.UUID      `gorm:"primaryKey" json:"guid"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	Sender       common.Address `json:"sender" gorm:"serializer:bytes"`
	Amount       *big.Int       `json:"amount" gorm:"serializer:u256"`
	BlockNumber  *big.Int       `json:"block_number" gorm:"serializer:u256"`
	Timestamp    uint64
}

type depositTokensDB struct {
	gorm *gorm.DB
}

type DepositTokensViews interface {
	QueryDepositTokensByPage(page int, pageSize int, order string) ([]DepositTokens, error)
}

type DepositTokensDB interface {
	DepositTokensViews
	SaveDepositTokens(*[]DepositTokens) error
}

func NewDepositTokensDB(db *gorm.DB) DepositTokensDB {
	return &depositTokensDB{
		gorm: db,
	}
}

func (d depositTokensDB) QueryDepositTokensByPage(page int, pageSize int, order string) ([]DepositTokens, error) {
	var depositTokens []DepositTokens
	// 计算 Offset
	offset := (page - 1) * pageSize
	// 执行分页查询
	if err := d.gorm.Order(order).Offset(offset).Limit(pageSize).Find(&depositTokens).Error; err != nil {
		return nil, err
	}
	return depositTokens, nil
}

func (d depositTokensDB) SaveDepositTokens(depositTokens *[]DepositTokens) error {
	result := d.gorm.CreateInBatches(&depositTokens, len(*depositTokens))
	return result.Error
}
