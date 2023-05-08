package repositories

import (
	"database/sql"
	"fmt"
	"isjhar/template/echo-golang/utils"
	"log"
	"strings"

	"gopkg.in/guregu/null.v4"
)

var DB *sql.DB

const hostDefault = "localhost"
const portDefault = "1433"
const userDefault = "sa"
const passwordDefault = "password"
const databaseDefault = "database"

func init() {
	var err error
	dataSourceName := GetDataSourceName()

	DB, err = sql.Open("sqlserver", dataSourceName)
	if err != nil {
		log.Panicf("error %v \n", err)
	}

	if err = DB.Ping(); err != nil {
		log.Panicf("error %v \n", err)
	}
}

func GetDataSourceName() string {
	return fmt.Sprintf("server=%s; port=%s; user id=%s; password=%s; database=%s;",
		utils.GetEnvironmentVariable("DB_HOST", hostDefault),
		utils.GetEnvironmentVariable("DB_PORT", portDefault),
		utils.GetEnvironmentVariable("DB_USER", userDefault),
		utils.GetEnvironmentVariable("DB_PASSWORD", passwordDefault),
		utils.GetEnvironmentVariable("DB_NAME", databaseDefault),
	)
}

func GetOrderQuery(order null.String) string {
	if strings.ToLower(order.String) == "desc" {
		return "desc"
	}
	return "asc"
}
