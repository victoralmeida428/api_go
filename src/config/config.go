package config

import (
	"api/src/logJson"
	"database/sql"
	"fmt"
	"github.com/golobby/dotenv"
	_ "github.com/lib/pq"
	"os"
)

type Config struct {
	Debug bool  `env:"DEBUG"`
	Port  int16 `env:"PORT"`

	Database struct {
		Host string `env:"DB_HOST"`
		Name string `env:"DB_NAME"`
		Port int16  `env:"DB_PORT"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASS"`
	}

	Logger *logJson.Logger
	Error  *logJson.Application
	DB     *sql.DB
}

func (c *Config) Init(port int16, env string) error {
	file, err := os.Open(env)
	if err != nil {
		return err
	}

	err = dotenv.NewDecoder(file).Decode(c)
	if c.Port == 0 {
		c.Port = port
	}

	logger := logJson.New(os.Stdout, logJson.LevelInfo)
	Error := logJson.GetError(logger)
	c.Logger = logger
	c.Error = Error
	dbdns := fmt.Sprintf("postgres://%s:%s@%s/%s",
		c.Database.User, c.Database.Pass, c.Database.Host, c.Database.Name,
	)
	db, err := sql.Open("postgres", dbdns)
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}

func New() *Config {
	return &Config{}
}
