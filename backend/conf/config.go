package conf

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shinecloudnet/explorer/backend/logger"
)

const (
	KeyDbAddr      = "DB_ADDR"
	KeyDATABASE    = "DB_DATABASE"
	KeyDbUser      = "DB_USER"
	KeyDbPwd       = "DB_PASSWORD"
	KeyDbPoolLimit = "DB_POOL_LIMIT"

	KeyServerPort    = "PORT"
	KeyAddrHubLcd    = "ADDR_NODE_SERVER"
	KeyAddrHubNode   = "ADDR_HUB_RPC"
	KeyAddrFaucet    = "FAUCET_URL"
	KeyChainId       = "CHAIN_ID"
	KeyApiVersion    = "API_VERSION"
	KeyMaxDrawCnt    = "MAX_DRAW_CNT"
	KeyShowFaucet    = "SHOW_FAUCET"
	KeyCurEnv        = "CUR_ENV"
	KeyInitialSupply = "INITIAL_SUPPLY"

	KeyPrefixAccAddr  = "PrefixAccAddr"
	KeyPrefixAccPub   = "PrefixAccPub"
	KeyPrefixValAddr  = "PrefixValAddr"
	KeyPrefixValPub   = "PrefixValPub"
	KeyPrefixConsAddr = "PrefixConsAddr"
	KeyPrefixConsPub  = "PrefixConsPub"

	KeyCronTimeAssetGateways       = "CronTimeAssetGateways"
	KeyCronTimeAssetTokens         = "CronTimeAssetTokens"
	KeyCronTimeGovParams           = "CronTimeGovParams"
	KeyCronTimeTxNumByDay          = "CronTimeTxNumByDay"
	KeyCronTimeValidators          = "CronTimeValidators"
	KeyCronTimeValidatorIcons      = "CronTimeValidatorIcons"
	KeyCronTimeProposalVoters      = "CronTimeProposalVoters"
	KeyCronTimeValidatorStaticInfo = "CronTimeValidatorStaticInfo"

	EnvironmentDevelop = "dev"
	EnvironmentLocal   = "local"
	EnvironmentQa      = "qa"
	EnvironmentStage   = "stage"
	EnvironmentProd    = "prod"

	InitialSupply      = "2000000000" //IRIS
	DefaultEnvironment = EnvironmentDevelop
)

var (
	config        Config
	defaultConfig = map[string]map[string]string{}
	//IniSupply     string
)

func init() {
	logger.Info("==================================load config start==================================")
	loadDefault()
	//IniSupply = getEnv(KeyInitialSupply, DefaultEnvironment)
	addrs := strings.Split(getEnv(KeyDbAddr, DefaultEnvironment), ",")
	db := dbConf{
		Addrs:     addrs,
		Database:  getEnv(KeyDATABASE, DefaultEnvironment),
		UserName:  getEnv(KeyDbUser, DefaultEnvironment),
		Password:  getEnv(KeyDbPwd, DefaultEnvironment),
		PoolLimit: getEnvInt(KeyDbPoolLimit, DefaultEnvironment),
	}
	config.Db = db

	rand.Seed(time.Now().Unix())
	server := serverConf{
		InstanceNo:                  fmt.Sprintf("%d-%d", time.Now().Unix(), rand.Int63n(100)),
		ServerPort:                  getEnvInt(KeyServerPort, DefaultEnvironment),
		FaucetUrl:                   getEnv(KeyAddrFaucet, DefaultEnvironment),
		ApiVersion:                  getEnv(KeyApiVersion, DefaultEnvironment),
		MaxDrawCnt:                  getEnvInt(KeyMaxDrawCnt, DefaultEnvironment),
		ShowFaucet:                  getEnv(KeyShowFaucet, DefaultEnvironment),
		CurEnv:                      getEnv(KeyCurEnv, DefaultEnvironment),
		CronTimeAssetGateways:       getEnvInt(KeyCronTimeAssetGateways, DefaultEnvironment),
		CronTimeAssetTokens:         getEnvInt(KeyCronTimeAssetTokens, DefaultEnvironment),
		CronTimeGovParams:           getEnvInt(KeyCronTimeGovParams, DefaultEnvironment),
		CronTimeTxNumByDay:          getEnvInt(KeyCronTimeTxNumByDay, DefaultEnvironment),
		CronTimeValidators:          getEnvInt(KeyCronTimeValidators, DefaultEnvironment),
		CronTimeValidatorIcons:      getEnvInt(KeyCronTimeValidatorIcons, DefaultEnvironment),
		CronTimeProposalVoters:      getEnvInt(KeyCronTimeProposalVoters, DefaultEnvironment),
		CronTimeValidatorStaticInfo: getEnvInt(KeyCronTimeValidatorStaticInfo, DefaultEnvironment),
	}
	logger.Info(fmt.Sprintf("serverInstanceNo: %s", server.InstanceNo))
	config.Server = server

	hubcf := hubConf{
		Prefix: bech32Prefix{
			AccAddr:  getEnv(KeyPrefixAccAddr, DefaultEnvironment),
			AccPub:   getEnv(KeyPrefixAccPub, DefaultEnvironment),
			ValAddr:  getEnv(KeyPrefixValAddr, DefaultEnvironment),
			ValPub:   getEnv(KeyPrefixValPub, DefaultEnvironment),
			ConsAddr: getEnv(KeyPrefixConsAddr, DefaultEnvironment),
			ConsPub:  getEnv(KeyPrefixConsPub, DefaultEnvironment),
		},
		LcdUrl:  getEnv(KeyAddrHubLcd, DefaultEnvironment),
		NodeUrl: getEnv(KeyAddrHubNode, DefaultEnvironment),
		ChainId: getEnv(KeyChainId, DefaultEnvironment),
	}
	config.Hub = hubcf

	logger.Info("==================================load config end==================================")
}

func loadDefault() {
	defaultConfig[EnvironmentDevelop] = map[string]string{
		KeyDbAddr:         "127.0.0.1:27017",
		KeyDATABASE:       "shinecloudnet-explorer",
		KeyDbUser:         "shinecloudnet",
		KeyDbPwd:          "shinecloudnetpassword",
		KeyDbPoolLimit:    "4096",
		KeyServerPort:     "9080",
		KeyAddrHubLcd:     "http://localhost:1317",
		KeyAddrHubNode:    "http://localhost:26657",
		KeyAddrFaucet:     "http://localhost:30200",
		KeyChainId:        "shinecloudnet-test",
		KeyApiVersion:     "v0.6.5",
		KeyMaxDrawCnt:     "10",
		KeyPrefixAccAddr:  "scloud",
		KeyPrefixAccPub:   "scloudpub",
		KeyPrefixValAddr:  "scloudvaloper",
		KeyPrefixValPub:   "scloudvaloperpub",
		KeyPrefixConsAddr: "scloudvalcons",
		KeyPrefixConsPub:  "scloudvalconspub",
		KeyShowFaucet:     "1",
		KeyCurEnv:         "dev",
		KeyInitialSupply:  InitialSupply,

		KeyCronTimeAssetGateways:       "60",
		KeyCronTimeAssetTokens:         "60",
		KeyCronTimeGovParams:           "3600",
		KeyCronTimeTxNumByDay:          "43200",
		KeyCronTimeValidators:          "60",
		KeyCronTimeProposalVoters:      "60",
		KeyCronTimeValidatorIcons:      "43200",
		KeyCronTimeValidatorStaticInfo: "300",
	}
}

func Get() Config {
	return config
}

type Config struct {
	Db     dbConf
	Server serverConf
	Hub    hubConf
}

type dbConf struct {
	Addrs     []string
	Database  string
	UserName  string
	Password  string
	PoolLimit int
}

type serverConf struct {
	InstanceNo                  string
	ServerPort                  int
	FaucetUrl                   string
	ApiVersion                  string
	MaxDrawCnt                  int
	ShowFaucet                  string
	CurEnv                      string
	CronTimeAssetGateways       int
	CronTimeAssetTokens         int
	CronTimeGovParams           int
	CronTimeTxNumByDay          int
	CronTimeValidators          int
	CronTimeValidatorStaticInfo int
	CronTimeValidatorIcons      int
	CronTimeProposalVoters      int
}

type hubConf struct {
	Prefix  bech32Prefix
	LcdUrl  string
	NodeUrl string
	ChainId string
}

type bech32Prefix struct {
	AccAddr  string
	AccPub   string
	ValAddr  string
	ValPub   string
	ConsAddr string
	ConsPub  string
}

func getEnv(key string, environment string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}
	if value == "" {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	if key == KeyDbUser || key == KeyDbPwd {
		logger.Info("config", logger.Bool(key+" is empty", value == ""))
	} else {
		logger.Info("config", logger.String(key, value))
	}

	return value
}

func getEnvInt(key string, environment string) int {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		if DefaultEnvironment == EnvironmentStage || DefaultEnvironment == EnvironmentProd {
			logger.Panic("config is not able to use default config", logger.String("Environment", DefaultEnvironment))
		}
		value = defaultConfig[environment][key]
	}

	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		logger.Panic("config must be not empty", logger.String("key", key))
	}
	logger.Info("config", logger.Int64(key, i))
	return int(i)
}
