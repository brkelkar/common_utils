package databases

import (
	"fmt"

	"github.com/brkelkar/common_utils/logger"
)

//DB used as cursor for database connection
//var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig Create required config format
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "AWACS",
		Password: "test",
	}
	logger.Debug(
		fmt.Sprintf("Connecting to Host_name: %s, at port %v, user_name %s, database name  %s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName))
	return &dbConfig
}

// DbURL Create database connetion url
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
