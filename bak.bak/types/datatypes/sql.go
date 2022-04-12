package datatypes

// types/sql.go contains interfaces that implement
// methods contained in the standard library sql
// package.

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/skeptycal/goutil/types"
)

type (
	Row        = sql.Row
	Rows       = sql.Rows
	ColumnType = sql.ColumnType
	TxOptions  = sql.TxOptions

	Closer = types.Closer
	Errer  = types.Errer

	DB interface{}

	Conner interface {
		PingContext(ctx context.Context) error
		ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row
		PrepareContext(ctx context.Context, query string) (Stmter, error)
		Raw(f func(driverConn interface{}) error) (err error)
		BeginTx(ctx context.Context, opts *TxOptions) (Txer, error)

		Closer
	}

	Txer interface {
		Commit() error
		Rollback() error
		PrepareContext(ctx context.Context, query string) (Stmter, error)
		Prepare(query string) (Stmter, error)
		StmtContext(ctx context.Context, stmt Stmter)
		Stmt(stmt Stmter) Stmter
		ExecContext(ctx context.Context, query string, args ...Any) (Result, error)
		Exec(query string, args ...Any) (Result, error)
		QueryContext(ctx context.Context, query string, args ...Any) (*Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...Any) *Row
		QueryRow(query string, args ...Any) *Row
	}

	// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the OpenDB function should be called just once. It is rarely necessary to close a DB.
	Connecter interface {
		Open() (*sql.DB, error)
		OpenDB() (*sql.DB, error)
		PingContext(ctx context.Context) error
		Ping() error
	}

	// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the OpenDB function should be called just once. It is rarely necessary to close a DB.
	connectCloser interface {
		Connecter

		// Close closes the database and prevents new queries from starting.
		// Close then waits for all queries that have started processing
		// on the server to finish.
		//
		// It is rare to Close a DB, as the DB handle is meant to be
		// long-lived and shared between many goroutines.
		Closer
	}

	// Scanner implements the sql.Scan method.
	//
	// Reference: from Go standard library
	Scanner interface {
		// Scan assigns a value from a database driver.
		//
		// The src value will be of one of the following types:
		//
		//    int64
		//    float64
		//    bool
		//    []byte
		//    string
		//    time.Time
		//    nil - for NULL values
		//
		// An error should be returned if the value cannot be stored
		// without loss of information.
		//
		// Reference types such as []byte are only valid until the next call to Scan
		// and should not be retained. Their underlying memory is owned by the driver.
		// If retention is necessary, copy their values before the next call to Scan.
		//
		// Scan method from standard library sql.go:
		// Scan copies the columns in the current row into the values pointed
		// at by dest. The number of values in dest must be the same as the
		// number of columns in Rows.
		//
		// Scan converts columns read from the database into the following
		// common Go types and special types provided by the sql package:
		//
		//    *string
		//    *[]byte
		//    *int, *int8, *int16, *int32, *int64
		//    *uint, *uint8, *uint16, *uint32, *uint64
		//    *bool
		//    *float32, *float64
		//    *Any
		//    *RawBytes
		//    *Rows (cursor value)
		//    any type implementing Scanner (see Scanner docs)
		//
		// In the most simple case, if the type of the value from the source
		// column is an integer, bool or string type T and dest is of type *T,
		// Scan simply assigns the value through the pointer.
		//
		// Scan also converts between string and numeric types, as long as no
		// information would be lost. While Scan stringifies all numbers
		// scanned from numeric database columns into *string, scans into
		// numeric types are checked for overflow. For example, a float64 with
		// value 300 or a string with value "300" can scan into a uint16, but
		// not into a uint8, though float64(255) or "255" can scan into a
		// uint8. One exception is that scans of some float64 numbers to
		// strings may lose information when stringifying. In general, scan
		// floating point columns into *float64.
		//
		// If a dest argument has type *[]byte, Scan saves in that argument a
		// copy of the corresponding data. The copy is owned by the caller and
		// can be modified and held indefinitely. The copy can be avoided by
		// using an argument of type *RawBytes instead; see the documentation
		// for RawBytes for restrictions on its use.
		//
		// If an argument has type *Any, Scan copies the value
		// provided by the underlying driver without conversion. When scanning
		// from a source value of type []byte to *Any, a copy of the
		// slice is made and the caller owns the result.
		//
		// Source values of type time.Time may be scanned into values of type
		// *time.Time, *Any, *string, or *[]byte. When converting to
		// the latter two, time.RFC3339Nano is used.
		//
		// Source values of type bool may be scanned into types *bool,
		// *Any, *string, *[]byte, or *RawBytes.
		//
		// For scanning into *bool, the source may be true, false, 1, 0, or
		// string inputs parseable by strconv.ParseBool.
		//
		// Scan can also convert a cursor returned from a query, such as
		// "select cursor(select * from my_table) from dual", into a
		// *Rows value that can itself be scanned from. The parent
		// select query will close any cursor *Rows if the parent *Rows is closed.
		//
		// If any of the first arguments implementing Scanner returns an error,
		// that error will be wrapped in the returned error
		//
		//	 func (rs *Rows) Scan(dest ...Any) error {
		//	 	rs.closemu.RLock()
		//
		//	 	if rs.lasterr != nil && rs.lasterr != io.EOF {
		//	 		rs.closemu.RUnlock()
		//	 		return rs.lasterr
		//	 	}
		//	 	if rs.closed {
		//	 		err := rs.lasterrOrErrLocked(errRowsClosed)
		//	 		rs.closemu.RUnlock()
		//	 		return err
		//	 	}
		//	 	rs.closemu.RUnlock()
		//
		//	 	if rs.lastcols == nil {
		// 		return errors.New("sql: Scan called without calling Ne	xt")
		//	 	}
		//	 	if len(dest) != len(rs.lastcols) {
		// 		return fmt.Errorf("sql: expected %d destination ar	guments in Scan, not %d", len(rs.lastcols), len(dest))
		//	 	}
		//	 	for i, sv := range rs.lastcols {
		//	 		err := convertAssignRows(dest[i], sv, rs)
		//	 		if err != nil {
		// 			return fmt.Errorf(`sql: Scan error on column in	dex %d, name %q: %w`, i, rs.rowsi.Columns()[i], err)
		//	 		}
		//	 	}
		//	 	return nil
		//	 }
		Scan(src Any) error
	}

	// Rower implements sql.Stmt methods.
	//
	// Stmt is a prepared statement.
	// A Stmt is safe for concurrent use by multiple goroutines.
	//
	// If a Stmt is prepared on a Tx or Conn, it will be bound to a single
	// underlying connection forever. If the Tx or Conn closes, the Stmt will
	// become unusable and all operations will return an error.
	// If a Stmt is prepared on a DB, it will remain usable for the lifetime of the
	// DB. When the Stmt needs to execute on a new underlying connection, it will
	// prepare itself on the new connection automatically.
	Stmter interface {
		ExecContext(ctx context.Context, args ...Any) (Result, error)
		Exec(args ...Any) (Result, error)
		QueryContext(ctx context.Context, args ...Any) (Rowser, error)
		Query(args ...Any) (Rowser, error)
		QueryRowContext(ctx context.Context, args ...Any) *Row
		QueryRow(args ...Any) *Row
		Closer
	}

	// Rowser implements sql.Rows methods.
	//
	// Rows is the result of a query. Its cursor starts
	// before the first row of the result set. Use Next
	// to advance from row to row.
	Rowser interface {
		Next() bool
		NextResultSet() bool
		Columns() ([]string, error)
		ColumnTypes() ([]*ColumnType, error)

		Errer
		Scanner
		Closer
	}

	// ColumnTyper implements sql.ColumnType methods.
	// ColumnType contains the name and type of a column.
	ColumnTyper interface {
		Name() string
		Length() (length int64, ok bool)
		DecimalSize() (precision, scale int64, ok bool)
		ScanType() reflect.Type
		Nullable() (nullable, ok bool)
		DatabaseTypeName() string
	}

	Rower interface {
		Scanner
		Errer
	}

	// A Result summarizes an executed SQL command.
	//
	// Reference: from Go standard library
	Result interface {
		// LastInsertId returns the integer generated by the database
		// in response to a command. Typically this will be from an
		// "auto increment" column when inserting a new row. Not all
		// databases support this feature, and the syntax of such
		// statements varies.
		LastInsertId() (int64, error)

		// RowsAffected returns the number of rows affected by an
		// update, insert, or delete. Not every database or database
		// driver may support this.
		RowsAffected() (int64, error)
	}
)
