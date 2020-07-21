package mysql_tool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}

func OpenDB(username, password, addr, port, dbName string) *gorm.DB {
	config := strings.Join([]string{
		username, ":", password,
		"@tcp(", addr, ":", port, ")/",
		dbName,
		"?charset=utf8mb4"},
		"")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		logrus.Fatalf("OpenDB失败, 数据库名称: %s, 错误信息: %s", dbName, err)
	} else {
		logrus.Fatalf("OpenDB成功, 数据库名称: %s", dbName)
	}

	SetupDB(db)
	return db
}

func SetupDB(db *gorm.DB) {
	// 设置日志的模式
	db.LogMode(viper.GetBool("gormlog"))
	// 设置最大连接数
	db.DB().SetMaxOpenConns(viper.GetInt("db_test.MaxOpenConn"))
	// 设置最大空闲连接数
	db.DB().SetMaxIdleConns(viper.GetInt("db_test.MaxIdleConn"))
	// 设置每个链接的过期时间
	//db.DB().SetConnMaxLifetime(time.Second * 5)
}

func InitDB() *gorm.DB {

	return OpenDB(
		viper.GetString("db_test.username"),
		viper.GetString("db_test.password"),
		viper.GetString("db_test.addr"),
		viper.GetString("db_test.port"),
		viper.GetString("db_test.dbName"),
	)
}

// 初始化配置文件的读取
func InitConfig() {
	viper.SetConfigType("yaml") // 读取yaml配置文件
	//path, _ := os.Getwd()
	//viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigName("./db_tool/mysql_tool/mysql_config") // 读取文件名称为
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

}

func GetDB() *gorm.DB {
	// TODO 添加传入参数
	// 1. 使用配置文件中的哪个key参数
	// 2. 配置文件路径参数
	InitConfig()
	return InitDB()

	// TODO 还能优化, 水平不够, 时间来凑
}
