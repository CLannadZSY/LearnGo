package sql

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const FILEPATH = "mysql_config"
const FILETYPE = "yaml"

type DB struct {
	write  *conn
	read   *conn
	idx    int64
	master *DB
}

//conn database connection
type conn struct {
	*sql.DB
	conf *Config
}

// Rows rows.
type Rows struct {
	*Rows
	cancel func()
}

// Configuration file structure
type BaseConfig struct {
	UserName,
	PassWord,
	Addr,
	Port,
	DBName string
}

type Config struct {
	DSN          string        // write data source name.
	ReadDSN      string        // read data source name.
	Active       int           // pool
	Idle         int           // pool
	IdleTimeout  time.Duration // connect max life time.
	QueryTimeout time.Duration // query sql timeout
	ExecTimeout  time.Duration // execute sql timeout
	TranTimeout  time.Duration // transaction sql timeout
}

func InitReadConfig(KeyName, FilePath, FileType string) BaseConfig {
	if FilePath == "" {
		FilePath = FILEPATH
	}
	if FileType == "" {
		FileType = FILETYPE
	}
	viper.AddConfigPath(".")
	viper.SetConfigName(FilePath) // 读取文件名称为
	viper.SetConfigType(FileType)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		panic(err) // 读取配置文件失败致命错误
	}

	return BaseConfig{
		viper.GetString(KeyName + ".username"),
		viper.GetString(KeyName + ".password"),
		viper.GetString(KeyName + ".addr"),
		viper.GetString(KeyName + ".port"),
		viper.GetString(KeyName + ".dbName"),
	}
}

func ConnectMysql(WriteKey, ReadKey string, FilePath, Filetype string) (db *DB) {
	writeConfig := InitReadConfig(WriteKey, FilePath, Filetype)
	dsn := FormatConfig(&writeConfig)

	readConfig := InitReadConfig(ReadKey, FilePath, Filetype)
	readDSN := FormatConfig(&readConfig)

	c := Config{
		dsn,
		readDSN,
		50,
		30,
		time.Hour,
		5 * time.Minute,
		3 * time.Minute,
		10 * time.Minute,
	}
	db = OpenMySQL(&c)
	//defer db.Close()
	return db
}

func FormatConfig(c *BaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.PassWord,
		c.Addr,
		c.Port,
		c.DBName)
}

func OpenMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transaction timeout")
	}
	dbRet, err := Open(c)
	if err != nil {
		log.Fatalf("open mysql error(%v)", err)
	}
	return dbRet
}

func Open(c *Config) (*DB, error) {
	db := new(DB)
	d, err := Connect(c, c.DSN)
	if err != nil {
		return nil, err
	}
	w := &conn{DB: d, conf: c}
	rs := &conn{DB: d, conf: c}
	db.write = w
	db.read = rs
	db.master = &DB{write: db.write}
	return db, nil
}

func Connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	d.SetMaxOpenConns(c.Active)
	d.SetMaxIdleConns(c.Idle)
	d.SetConnMaxLifetime(c.IdleTimeout)
	return d, nil
}

func (db *DB) Close() (err error) {
	if e := db.write.Close(); e != nil {
		err = errors.WithStack(e)
	}

	if e := db.read.Close(); e != nil {
		err = errors.WithStack(e)
	}
	return
}
