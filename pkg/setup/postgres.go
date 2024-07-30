package setup

import (
	"fmt"
	"time"

	"deliver/config"
	"deliver/internal/constants"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var fieldDatabase = logrus.Fields{
	constants.LoggerCategory: constants.LoggerCategoryDatabase,
}

// SQLXConfig holds the configuration for the database instance
type SQLXConfig struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

// InitializeSQLXDatabase returns a new DBInstance
func (cfg *SQLXConfig) InitializeSQLXDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	logrus.WithFields(fieldDatabase).Infof("Setting maximum number of open connections to %d", cfg.MaxOpenConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	logrus.WithFields(fieldDatabase).Infof("Setting maximum number of idle connections to %d", cfg.MaxIdleConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	logrus.WithFields(fieldDatabase).Infof("Setting maximum lifetime for a connection to %s", cfg.MaxLifetime)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

// SetupPostgresConnection sets up a connection to the PostgreSQL database using sqlx
func SetupPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
	var dsn string

	switch cfg.Environment {
	case constants.EnvironmentDevelopment:
		dsn = cfg.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = cfg.DBPostgreURL
	default:
		return nil, fmt.Errorf("unknown environment: %v", cfg.Environment)
	}

	// Setup sqlx config for PostgreSQL
	sqlxConfig := SQLXConfig{
		DriverName:     cfg.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	// Initialize PostgreSQL connection with sqlx
	conn, err := sqlxConfig.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
