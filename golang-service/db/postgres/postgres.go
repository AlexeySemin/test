package postgres

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	defaultDBName = "testagg"
	dafaultHost   = "localhost"
)

type postgresConfig struct {
	Name     string
	Host     string
	User     string
	Password string
}

func (cfg *postgresConfig) getConnString() string {
	userStr := cfg.User
	if cfg.Password != "" {
		userStr = userStr + ":" + url.PathEscape(cfg.Password)
	}

	params := ""
	if cfg.Host == "localhost" {
		params = "?sslmode=disable"
	}

	return fmt.Sprintf("postgres://%s@%s/%s%s", userStr, cfg.Host, cfg.Name, params)
}

func newPostgresConfig() (*postgresConfig, error) {
	dbName := defaultDBName
	if dbn := os.Getenv("POSTGRES_DB_NAME"); dbn != "" {
		dbName = dbn
	}

	dbHost := dafaultHost
	if dbh := os.Getenv("POSTGRES_DB_HOST"); dbh != "" {
		dbHost = dbh
	}

	dbUser := os.Getenv("POSTGRES_DB_USER")
	if dbUser == "" {
		u, err := user.Current()
		if err == nil {
			dbUser = u.Username
		} else {
			return nil, err
		}
	}

	dbPassword := os.Getenv("POSTGRES_DB_PASSWORD")

	return &postgresConfig{dbName, dbHost, dbUser, dbPassword}, nil
}

// NewDB init and return new connection
func NewDB(syncedModels []interface{}) (*gorm.DB, error) {
	config, err := newPostgresConfig()
	if err != nil {
		return nil, err
	}

	connString := config.getConnString()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	for _, model := range syncedModels {
		log.Print("Model ", reflect.TypeOf(model), " is syncing...")

		migration := db.AutoMigrate(model)
		if migration.Error != nil {
			return nil, migration.Error
		}
	}

	log.Print("Models synced")

	return db, nil
}
