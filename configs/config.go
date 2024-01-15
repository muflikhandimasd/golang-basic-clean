package configs

import (
	"database/sql"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type config struct {
	dbMaster   *sql.DB
	log        *logrus.Entry
	ctxTimeout time.Duration
	port       string
}

type Config interface {
	DB() *sql.DB
	Log() *logrus.Entry
	Port() string
	ContextTimeout() time.Duration
}

func NewConfig() Config {
	cfg := new(config)
	cfg.initService()
	return cfg
}

func (c *config) initService() {
	var loadonce sync.Once
	loadonce.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
		dbMasterCfg := SQLConfig{
			DB_HOST: os.Getenv("DB_HOST"),
			DB_USER: os.Getenv("DB_USER"),
			DB_PASS: os.Getenv("DB_PASS"),
			DB_NAME: os.Getenv("DB_NAME"),
			DB_PORT: os.Getenv("DB_PORT"),
		}

		c.port = os.Getenv("PORT")
		c.dbMaster = InitDBSQL(dbMasterCfg)

		c.log = InitLog("api.log")
		c.ctxTimeout = 5 * time.Minute
	})

}

func (c *config) DB() *sql.DB {
	return c.dbMaster
}

func (c *config) Port() string {

	return c.port
}

func (c *config) ContextTimeout() time.Duration {
	return c.ctxTimeout
}

func (c *config) Log() *logrus.Entry {
	return c.log
}
