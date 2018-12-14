package config

import "github.com/go-sql-driver/mysql"

type Connection struct {
	mysql.Config
	MaxIdleConns int
	MaxOpenConns int
}

func GetCons() map[string]Connection {

	//return map[string]*Connections{
	//	"official_account": &Connections{
	//		DATABASE_USERNAME: "11",
	//		DATABASE_IP:       "11",
	//		DATABASE_PASSWORD: "11",
	//		DATABASE_NAME:     "11",
	//		DATABASE_PORT:     "3306",
	//	},
	//}

	return map[string]Connection{}
}
