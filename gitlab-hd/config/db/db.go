package db

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //gorm库
	_ "github.com/jinzhu/gorm/dialects/mysql" //gorm对应的mysql驱动
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
	"time"
)

var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

//db的初始化函数，与数据库建立连接
func InitDB() {
	//判断是否已经初始化了
	if isInit {
		return
	}
	//组装连接配置
	//parseTime是查询结果是否自动解析为时间
	//loc是Mysql的时区设置
	DbType := viper.GetString("db.type")
	user := viper.GetString("db.user")
	pwd := viper.GetString("db.pwd")
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	dbname := viper.GetString("db.dbname")
	maxIdleConns := viper.GetInt("db.maxIdleConns")
	maxOpenConns := viper.GetInt("db.maxOpenConns")
	connMaxLifetime := viper.GetInt("db.connMaxLifetime")
	logMode := viper.GetBool("db.logMode")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		port,
		dbname)
	//与数据库建立连接，生成一个*gorm.DB类型的对象
	GORM, err = gorm.Open(DbType, dsn)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	//打印sql语句
	GORM.LogMode(logMode)

	//开启连接池
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	GORM.DB().SetMaxIdleConns(maxIdleConns)
	// 设置了连接可复用的最大时间
	GORM.DB().SetMaxOpenConns(maxOpenConns)
	// 设置了连接可复用的最大时间
	GORM.DB().SetConnMaxLifetime(time.Duration(connMaxLifetime))

	isInit = true

	//自动生成表
	//GORM.AutoMigrate(model.Chart{}, model.Event{})
	/*err := GORM.AutoMigrate(model.Project{})
	if err != nil {
		_ = fmt.Errorf("自动生成user表失败")
		panic(err)
	}*/

	logger.Info("连接数据库成功!")
}

//db的关闭函数
func Close() error {
	logger.Info("关闭数据库连接")
	return GORM.Close()
}