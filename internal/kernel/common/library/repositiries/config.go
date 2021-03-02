package repositiries

import (
	"os"
)

type Config = map[string]interface{}

type IConfigRepository interface {
	Set(key string, value interface{})
	Get(key string, def interface{}) interface{}
	Init()
}

type ConfigRepository struct {
	configs Config
}

func (cfg *ConfigRepository) Get(key string, def interface{}) interface{} {
	if val, ok := cfg.configs[key]; ok {
		return val
	}
	return def
}

func (cfg *ConfigRepository) Set(key string, value interface{}) {
	cfg.configs[key] = value
}

func (cfg *ConfigRepository) Init() {
	cfg.Set("app.version", os.Getenv("APP_VERSION"))
	cfg.Set("app.currency", os.Getenv("APP_CURRENCY"))
	cfg.Set("app.port", os.Getenv("APP_PORT"))
	cfg.Set("reader.node.eth", os.Getenv("NODE_ETH"))
	cfg.Set("reader.node.trx", os.Getenv("NODE_TRX"))
	cfg.Set("reader.node.xlm", os.Getenv("NODE_XLM"))
	cfg.Set("reader.node.xrp", os.Getenv("NODE_XRP"))
}

func NewConfigRepository(cfg Config) *ConfigRepository {
	instance := &ConfigRepository{configs: cfg}
	instance.Init()
	return instance
}
