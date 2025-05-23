package chain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/kitshy/treasure-backend/database/utils"
	"gorm.io/gorm"
	"math/big"
)

type BlockHeaders struct {
	GUID       uuid.UUID   `gorm:"primaryKey"`
	Hash       common.Hash `gorm:"serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RlpHeader  *utils.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
}

func (BlockHeaders) TableName() string {
	return "block_headers"
}

type blockHeadersDB struct {
	gorm *gorm.DB
}

type BlockHeadersViews interface {
	QueryBlockHeadersByID(uuid.UUID) (*BlockHeaders, error)
	LatestBlockHeader() (*BlockHeaders, error)
	QueryBlockHeaders(*BlockHeaders) ([]BlockHeaders, error)
	BlockHeaderWithScope(func(db *gorm.DB) *gorm.DB) (*BlockHeaders, error)
}

type BlockHeadersDB interface {
	BlockHeadersViews
	SaveBlockHeaders(*[]BlockHeaders) error
}

func NewBlockHeaders(db *gorm.DB) BlockHeadersDB {
	return &blockHeadersDB{gorm: db}
}

func (b blockHeadersDB) SaveBlockHeaders(headers *[]BlockHeaders) error {
	result := b.gorm.CreateInBatches(headers, len(*headers))
	return result.Error
}

func (b blockHeadersDB) QueryBlockHeadersByID(id uuid.UUID) (*BlockHeaders, error) {
	blockHeaders := &BlockHeaders{}
	result := b.gorm.Where(&BlockHeaders{GUID: id}).Take(blockHeaders)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return blockHeaders, nil
}

func (b blockHeadersDB) LatestBlockHeader() (*BlockHeaders, error) {
	var header BlockHeaders
	result := b.gorm.Table("block_headers").Order("number DESC").Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}

func (b blockHeadersDB) QueryBlockHeaders(param *BlockHeaders) ([]BlockHeaders, error) {
	var blockHeaders []BlockHeaders
	result := b.gorm.Where(&param).Take(&blockHeaders)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return blockHeaders, nil

}

func (b blockHeadersDB) BlockHeaderWithScope(f func(db *gorm.DB) *gorm.DB) (*BlockHeaders, error) {
	var header BlockHeaders
	result := b.gorm.Table("block_headers").Scopes(f).Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}
