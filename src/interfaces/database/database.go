package database

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type DBRepository interface {
	GetMaster() *sqlx.DB
	GetSlave() *sqlx.DB
	GetTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(dbTx *sqlx.Tx) error
	RollbackTransaction(dbTx *sqlx.Tx) error
	Close()
}

const (
	DriverMySQL = "mysql"
)

type DSNConfig struct {
	DSN string
}

//DBConfig for databases configuration
type DBConfig struct {
	SlaveDSN  string
	MasterDSN string
}

//DB configuration
type DB struct {
	DBConnection *sqlx.DB
	DBString     string
	MaxIdleConn  int
	MaxConn      int
}

//Db object
var (
	Master *DB
	Slave  *DB
)

type Store struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

type Options struct {
	dbTx *sqlx.Tx
}

func (s *Store) GetMaster() *sqlx.DB {
	return s.Master
}

func (s *Store) GetSlave() *sqlx.DB {
	return s.Slave
}

func New(cfg DBConfig, dbDriver string) (*Store, error) {
	masterDSN := cfg.MasterDSN
	slaveDSN := cfg.SlaveDSN

	Master = &DB{
		DBString:    masterDSN,
		MaxIdleConn: 10,
		MaxConn:     200,
	}

	err := Master.ConnectAndMonitor(dbDriver, "master")
	if err != nil {
		return &Store{}, err
	}
	Slave = &DB{
		DBString:    slaveDSN,
		MaxIdleConn: 10,
		MaxConn:     200,
	}
	err = Slave.ConnectAndMonitor(dbDriver, "slave")
	if err != nil {
		return &Store{}, err
	}

	return &Store{Master: Master.DBConnection, Slave: Slave.DBConnection}, nil
}

// Connect to database
func (d *DB) Connect(driver string) error {
	var db *sqlx.DB
	var err error
	db, err = sqlx.Open(driver, d.DBString)

	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
	} else {
		d.DBConnection = db
		err = db.Ping()
		if err != nil {
			log.Println("[Error]: DB connection error", err.Error())
		}
		return err
	}

	db.SetMaxOpenConns(d.MaxConn)
	db.SetMaxIdleConns(d.MaxIdleConn)

	return err
}

// ConnectAndMonitor to database
func (d *DB) ConnectAndMonitor(driver string, name string) error {
	err := d.Connect(driver)

	if err != nil {
		log.Printf("Not connected to database %s, name %s", d.DBString, name)
		return err
	} else {
		log.Printf("Success connecting to database %s name %s", d.DBString, name)
	}

	return nil
}

func (dbStore *Store) GetTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return dbStore.GetMaster().BeginTxx(ctx, nil)
}

func (dbStore *Store) CommitTransaction(dbTx *sqlx.Tx) error {
	return dbTx.Commit()
}

func (dbStore *Store) RollbackTransaction(dbTx *sqlx.Tx) error {
	return dbTx.Rollback()
}

func (dbStore *Store) Close() {
	dbStore.Master.Close()
	dbStore.Slave.Close()
}
