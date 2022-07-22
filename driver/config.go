package driver

import (
	"github.com/golang-api-user/user/models"
	"fmt"
	"log"
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Config() *gorm.DB {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.db_name")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	getConnection, err := gorm.Open("postgres", connStr)

	err = getConnection.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// common.InitTable(getConnection)

	getConnection.SingularTable(true)

	getConnection.Debug().AutoMigrate(
		&models.User{},
	)

	return getConnection

}
