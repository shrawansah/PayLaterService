package databases

import (
	"database/sql"
	"fmt"
	"simpl.com/configs"
	. "simpl.com/loggers"

	_ "github.com/go-sql-driver/mysql"
)

var simplDBH *sql.DB

func GetConnection() *sql.DB {
	if simplDBH != nil {
		return simplDBH
	}

	host := configs.Configs.DBHost
	user := configs.Configs.DBUser
	password := configs.Configs.DBPassword
	port := configs.Configs.DBPort
	schema := configs.Configs.DBSchema
	driver := configs.Configs.DBDriver

	connectionString := fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s", user, password, host, port, schema)
	connectionString += "?parseTime=true&charset=utf8"

	db, err := sql.Open(driver, connectionString)
	if err != nil {
		Logger.Error("Failed connecting to database." + err.Error())
	}

	simplDBH = db
	return simplDBH
				
}