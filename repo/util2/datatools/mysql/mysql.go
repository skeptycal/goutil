package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// New returns a new database connection pool (DB) given a
// configuration object and a database name.
// DB is a database handle representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple
// goroutines.
//
// The database name is optional and is used to create a connection string (DSN).
// Important settings included in the configuration object are applied to
// the open pool before it is returned. These are the defaults:
//
//      DB.SetConnMaxLifetime(time.Minute * 3)
//      DB.SetMaxOpenConns(10)
//      DB.SetMaxIdleConns(10)
//
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines, however DB.Close is
// available to close the database and prevent new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// The sql package creates and frees connections automatically; it
// also maintains a free pool of idle connections. If the database has
// a concept of per-connection state, such state can be reliably observed
// within a transaction (Tx) or connection (Conn). Once DB.Begin is called, the
// returned Tx is bound to a single connection. Once Commit or
// Rollback is called on the transaction, that transaction's
// connection is returned to DB's idle connection pool. The pool size
// can be controlled with SetMaxIdleConns.
func New(dbconfig DBConfig, database string) (*sql.DB, error) {

	// Open database connection.
	db, err := sql.Open("mysql", dbconfig.dsn(database))
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
