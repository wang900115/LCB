package gorm

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option struct {
	User     string
	Password string
	Host     string
	Name     string
	Debug    bool
}

func NewOption(conf *viper.Viper) Option {
	return Option{
		User:     conf.GetString("mysql.user"),
		Password: conf.GetString("mysql.password"),
		Host:     conf.GetString("mysql.host"),
		Name:     conf.GetString("mysql.name"),
		Debug:    conf.GetBool("mysql.debug"),
	}
}

func NewMysql(option Option) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		option.User,
		option.Password,
		option.Host,
		option.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Mysql: %v", err)
	}

	return db
}
