package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/isjhar/iet/internal/config"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

var DB *sql.DB
var ORM *gorm.DB

func Connect() {
	var err error
	dataSourceName := GetDataSourceName()

	DB, err = sql.Open("sqlserver", dataSourceName)
	if err != nil {
		log.Panicf("error %v \n", err)
	}

	if err = DB.Ping(); err != nil {
		log.Panicf("error %v \n", err)
	}

	// OPEN CONNECTION  HERE
}

func GetDataSourceName() string {
	return fmt.Sprintf("server=%s; port=%s; user id=%s; password=%s; database=%s;",
		config.Database.Host.String,
		config.Database.Port.String,
		config.Database.User.String,
		config.Database.Password.String,
		config.Database.Database.String,
	)
}

func GetOrderQuery(order null.String) string {
	if strings.ToLower(order.String) == "desc" {
		return "desc"
	}
	return "asc"
}

func ToTsVectorSearchQuery(search null.String) string {
	if !search.Valid {
		return ""
	}
	return strings.Replace(strings.Trim(search.String, " "), " ", "&", -1) + ":*"
}

func TimeStampToUTC(timestamp null.Time) null.Time {
	var result null.Time
	if timestamp.Valid {
		timestampTime := timestamp.Time
		result = null.TimeFrom(time.Date(timestampTime.Year(), timestampTime.Month(), timestampTime.Day(), timestampTime.Hour(), timestampTime.Minute(), timestampTime.Second(), timestampTime.Nanosecond(), time.Local).UTC())
	}
	return result
}
