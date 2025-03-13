package config

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Migrations    string       `json:"migrations"`
	HttpService   ServerConfig `json:"http_service"`
	Chain         ChainConfig  `json:"chain"`
	MasterDB      DBConfig     `json:"master_db"`
	SlaveDB       DBConfig     `json:"slave_db"`
	SlaveDbEnable bool
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

func NewConfig(ci *cli.Context) (*Config, error) {
	conf := Config{
		Migrations: ci.String(MigrationsFlags.Name),
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
		SlaveDbEnable: ci.Bool(SlaveDbEnableFlags.Name),
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
