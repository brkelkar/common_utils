package databases

import (
	"fmt"

	"github.com/brkelkar/common_utils/logger"
	"gorm.io/gorm"
)

//DB used as cursor for database connection
var DB *gorm.DB

// DBMsSQLConfig represents db configuration
type DBMsSQLConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBMsSQLConfig Create required config format
func BuildDBMsSQLConfig() *DBMsSQLConfig {
	dbConfig := DBMsSQLConfig{
		Host:     "127.0.0.1",
		Port:     1433,
		User:     "sqlserver",
		DBName:   "awacs_db",
		Password: "test",
	}
	logger.Debug(
		fmt.Sprintf("Connecting to Host_name: %s, at port %v, user_name %s, database name  %s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName))
	return &dbConfig
}

// DbMsSQLURL Create database connetion url
func DbMsSQLURL(dbConfig *DBMsSQLConfig) string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
