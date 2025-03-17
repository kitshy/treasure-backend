package eventlog

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type GrantRewardTokens struct {
	GUID         uuid.UUID      `gorm:"primaryKey" json:"guid"`
	BlockNumber  *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	Granter      common.Address `json:"granter" gorm:"serializer:bytes"`
	Amount       *big.Int       `json:"amount" gorm:"serializer:u256"`
	Timestamp    uint64
}

type grantRewardTokensDB struct {
	gorm *gorm.DB
}

type GrantRewardTokensViews interface {
	QueryGrantRewardTokensByID(uuid.UUID) ([]GrantRewardTokens, error)
}

type GrantRewardTokensDB interface {
	SaveGrantRewardTokens([]*GrantRewardTokens) error
}

func NewGrantRewardTokensDB(db *gorm.DB) GrantRewardTokensDB {
	return &grantRewardTokensDB{gorm: db}
}

func (*grantRewardTokensDB) QueryGrantRewardTokensByID(id uuid.UUID) ([]GrantRewardTokens, error) {
	result := []GrantRewardTokens{}
	return result, nil
}

func (*grantRewardTokensDB) SaveGrantRewardTokens(grantRewardTokens []*GrantRewardTokens) error {
	return nil
}
