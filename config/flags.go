package config

import "github.com/urfave/cli/v2"

const envVarPrefix = "TREASURE"

func prefixEnvVars(name string) []string {
	return []string{envVarPrefix + "_" + name}
}

var (
	MigrationsFlags = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "Directory to store migrations",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}

	// http service config
	HttpHostFlags = &cli.StringFlag{
		Name:     "http-host",
		Usage:    "",
		EnvVars:  prefixEnvVars("HTTP_HOST"),
		Required: true,
	}

	HttpPortFlags = &cli.IntFlag{
		Name:     "http-port",
		Usage:    "",
		EnvVars:  prefixEnvVars("HTTP_PORT"),
		Required: true,
	}

	// chain config
	ChainRpcUrlFlags = &cli.StringFlag{
		Name:    "chain-rpc-url",
		Value:   "",
		Usage:   "Chain RPC URL",
		EnvVars: prefixEnvVars("CHAIN_RPC_URL"),
	}
	ChainIdFlags = &cli.UintFlag{
		Name:     "chain-id",
		Usage:    "",
		EnvVars:  prefixEnvVars("CHAIN_ID"),
		Required: true,
	}
	ContractsFlags = &cli.StringFlag{
		Name:     "contracts",
		Usage:    "",
		EnvVars:  prefixEnvVars("CONTRACTS"),
		Required: true,
	}
	StartHeightFlags = &cli.Uint64Flag{
		Name:    "start-height",
		Usage:   "",
		EnvVars: prefixEnvVars("START_HEIGHT"),
		Value:   0,
	}
	BlockStepFlags = &cli.Uint64Flag{
		Name:     "block-step",
		Usage:    "",
		EnvVars:  prefixEnvVars("BLOCK_STEP"),
		Required: true,
		Value:    5,
	}

	// master db config
	MasterDbHostFlags = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlags = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlags = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlags = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlags = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}

	// slave db config
	SlaveDbEnableFlags = &cli.BoolFlag{
		Name:     "slave-db-enable",
		Usage:    "",
		EnvVars:  prefixEnvVars("SLAVE_DB_ENABLE"),
		Required: true,
	}
	SlaveDbHostFlags = &cli.StringFlag{
		Name:    "slave-db-host",
		Usage:   "",
		EnvVars: prefixEnvVars("SLAVE_DB_HOST"),
	}
	SlaveDbPortFlags = &cli.IntFlag{
		Name:    "slave-db-port",
		Usage:   "",
		EnvVars: prefixEnvVars("SLAVE_DB_PORT"),
	}
	SlaveDbUserFlags = &cli.StringFlag{
		Name:    "slave-db-user",
		Usage:   "",
		EnvVars: prefixEnvVars("SLAVE_DB_USER"),
	}
	SlaveDbPasswordFlags = &cli.StringFlag{
		Name:    "slave-db-password",
		Usage:   "",
		EnvVars: prefixEnvVars("SLAVE_DB_PASSWORD"),
	}
	SlaveDbNameFlags = &cli.StringFlag{
		Name:    "slave-db-name",
		Usage:   "",
		EnvVars: prefixEnvVars("SLAVE_DB_NAME"),
	}
)

var requiredFlags = []cli.Flag{
	MigrationsFlags,
	HttpHostFlags,
	HttpPortFlags,
	ContractsFlags,
	ChainRpcUrlFlags,
	ChainIdFlags,
	MasterDbHostFlags,
	MasterDbPortFlags,
	MasterDbUserFlags,
	MasterDbPasswordFlags,
	MasterDbNameFlags,
	SlaveDbEnableFlags,
}
var optionalFlags = []cli.Flag{
	StartHeightFlags,
	BlockStepFlags,
	SlaveDbHostFlags,
	SlaveDbPortFlags,
	SlaveDbUserFlags,
	SlaveDbPasswordFlags,
	SlaveDbNameFlags,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
