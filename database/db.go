package database

import (
	"io"

	logger "github.com/apsdehal/go-logger"
	"github.com/asdine/storm"
)

type DB struct {
	*storm.DB
	*logger.Logger
}

// NewDB returns a new database object,
// it configures the database for you.
func NewDB(dbPath, format string, logLevel int, logWriter io.WriteCloser) (*DB, error) {
	//Note! Do not use logger as you have no idea if logWriter has been configured for output yet
	var db DB
	log, err := logger.New("db logger", 1, logWriter)
	if err != nil {
		return &db, err
	}
	log.SetLogLevel(logger.LogLevel(logLevel))
	log.SetFormat(format)
	db.Logger = log
	if err := db.ConfigureDB(dbPath); err != nil {
		log.ErrorF("Error configuring the database ", err)
		return &db, err
	}
	return &db, nil
}
