// Package mysql provides support for persistent database connections.
//
// Copyright(C)2020 Micael Treanor
//
//
// Requirements:
// uses github.com/go-sql-driver/mysql which requires
// MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+)
//
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// mySqlUserVariable and mySqlPassword are the names of the environment variables
	// used to store connection information.
	mySQLUserNameEnvVar = "MYSQL_USERNAME"
	mySQLPasswordEnvVar = "MYSQL_PASSWORD"

	defaultProtocol  = "tcp"
	defaultMySQLHost = "localhost" // defaults for localhost are most secure
	defaultMySQLPort = "33060"     // depending on the MySQL version; this may need to be 3306

	// this is the 'driver name' used by helper functions that smooth out connections
	mySQLDriverName = "mysql"
)

// ErrMySQLNotImplemented returns an error if the method is not yet implemented
var ErrMySQLNotImplemented error = errors.New("not implemented")

// DBConfig defines the configuration for the database connection
type DBConfig interface {
	Config() string
	dsn(database string) string
	Load(file string) error
	Connect(dbname string) (*sql.DB, error)
	// Query(query string) (sql.Result, error)
	Save(file string) error
}

// dbConfig defines the configuration object for the database connection.
// Important settings included in the configuration object are applied to
// the open pool before it is returned by DbConnect. These are the defaults:
//
//      DB.SetConnMaxLifetime(time.Minute * 3)
//      DB.SetMaxOpenConns(10)
//      DB.SetMaxIdleConns(10)

type dbConfig struct {
	Username     string
	password     string
	Protocol     string `default:"tcp"`
	Host         string `default:"localhost"` // defaults for localhost are most secure
	Port         string `default:"33060"`     // depending on the MySQL version; this may need to be 3306
	Logging      bool   `default:"false"`
	Maxlifetime  int    `default:"180000000000"` // time.Minute * 3
	Maxopenconns int    `default:"10"`
	Maxidleconns int    `default:"10"`
}

// NewDBConfig returns a new MySQL database connection configuration object.
//
// The username and password are required and are automatically read in from
// environment variables.
//
// The logging option must be supplied and is either true or false.
// The host, port, and protocol are optional and are often set to the defaults.
// These options cannot be changed once the connection is established. A new
// connection must be created in place of the old one.
func NewDBConfig(host, port, protocol string, logging bool) (DBConfig, error) {

	username := os.Getenv(mySQLUserNameEnvVar)
	if username == "" {
		return nil, fmt.Errorf("environment variable %s for MySQL username not found", mySQLUserNameEnvVar)
	}

	password := os.Getenv(mySQLPasswordEnvVar)
	if password == "" {
		return nil, fmt.Errorf("environment variable %s for MySQL password not found", mySQLPasswordEnvVar)
	}

	if host == "" {
		host = defaultMySQLHost
	}
	if port == "" {
		port = defaultMySQLPort
	}
	if protocol == "" {
		protocol = defaultProtocol
	}

	return &dbConfig{
		Username: username,
		password: password,
		Protocol: protocol,
		Host:     host,
		Port:     port,
	}, nil
}

// func (db *dbConfig) Query(query string) (sql.Result, error) {
// 	// todo - stuff
// 	return nil, nil
// }

func (db *dbConfig) Config() string {
	return fmt.Sprintf(`%s://%s:%s`, db.Protocol, db.Host, db.Port)
}

// Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
//
// Most users will open a database via a driver-specific connection helper function that returns a *DB. No database drivers are included in the Go standard library. See https://golang.org/s/sqldrivers for a list of third-party drivers.
//
// Important settings
//
// db.SetConnMaxLifetime() is required to ensure connections are closed by the driver safely before connection is closed by MySQL server, OS, or other middlewares. Since some middlewares close idle connections by 5 minutes, we recommend timeout shorter than 5 minutes. This setting helps load balancing and changing system variables too.
//
// db.SetMaxOpenConns() is highly recommended to limit the number of connection used by the application. There is no recommended limit number because it depends on application and MySQL server.
//
// db.SetMaxIdleConns() is recommended to be set same to (or greater than) db.SetMaxOpenConns(). When it is smaller than SetMaxOpenConns(), connections can be opened and closed very frequently than you expect. Idle connections can be closed by the db.SetConnMaxLifetime(). If you want to close idle connections more rapidly, you can use db.SetConnMaxIdleTime() since Go 1.15.
//
// Open may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call Ping.
//
// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB.
func (db *dbConfig) Connect(dbname string) (*sql.DB, error) {
	return New(db, dbname)
}

// DSN returns the entire DSN authentication string including a database name.
// This allows the selection of the current database to act upon.
//
// DSN Format:
// Except for the dbname, all values are optional. So the minimal DSN is:
//
//      /dbname
// If you do not want to preselect a database, leave dbname empty.
//
// DSN format:
//      [username[:password]@][defaultProtocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//
// A DSN in its fullest form:
//      username:password@defaultProtocol(address)/dbname?param=value
func (db *dbConfig) dsn(dbname string) string {
	// "%s:%s@%s(%s:%s)/%s"
	return fmt.Sprintf("%s:%s@%s(%s:%s/%s)", db.Username, db.password, db.Protocol, db.Host, db.Port, dbname)
}

// Load loads the database configuration from a json file
//
// Not Implemented
func (db *dbConfig) Load(file string) error {
	// load json config file
	return ErrMySQLNotImplemented
}

// Load saves the database configuration to a json file
//
// Not Implemented
func (db *dbConfig) Save(file string) error {
	// save json config file
	return ErrMySQLNotImplemented
}
