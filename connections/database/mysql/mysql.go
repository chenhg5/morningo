package database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"morningo/config"
)

var SqlDB *sql.DB

type sqlTx struct {
	Tx *sql.Tx
}

var SqlTx sqlTx

// 只会执行一次在执行程序启动的时候
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", config.GetEnv().DATABASE_USERNAME+
		":"+config.GetEnv().DATABASE_PASSWORD+"@tcp("+config.GetEnv().DATABASE_IP+
		":"+config.GetEnv().DATABASE_PORT+")/"+config.GetEnv().DATABASE_NAME + "?charset=utf8mb4")
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	// 设置数据库最大连接 减少timewait
	SqlDB.SetMaxIdleConns(2000)
	SqlDB.SetMaxOpenConns(2000)
}

func Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	rs, err := SqlDB.Query(query, args...)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		log.Fatalln(colErr)
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			switch typeVal[i].DatabaseTypeName() {
			case "INT":
				var s sql.NullInt64
				colVar[i] = &s
			case "TINYINT":
				var s sql.NullInt64
				colVar[i] = &s
			case "MEDIUMINT":
				var s sql.NullInt64
				colVar[i] = &s
			case "SMALLINT":
				var s sql.NullInt64
				colVar[i] = &s
			case "BIGINT":
				var s sql.NullInt64
				colVar[i] = &s
			case "FLOAT":
				var s sql.NullFloat64
				colVar[i] = &s
			case "DOUBLE":
				var s sql.NullFloat64
				colVar[i] = &s
			case "DECIMAL":
				var s sql.NullInt64
				colVar[i] = &s
			case "DATE":
				var s sql.NullString
				colVar[i] = &s
			case "TIME":
				var s sql.NullString
				colVar[i] = &s
			case "YEAR":
				var s sql.NullString
				colVar[i] = &s
			case "DATETIME":
				var s sql.NullString
				colVar[i] = &s
			case "TIMESTAMP":
				var s sql.NullString
				colVar[i] = &s
			case "VARCHAR":
				var s sql.NullString
				colVar[i] = &s
			case "MEDIUMTEXT":
				var s sql.NullString
				colVar[i] = &s
			case "LONGTEXT":
				var s sql.NullString
				colVar[i] = &s
			case "TINYTEXT":
				var s sql.NullString
				colVar[i] = &s
			case "TEXT":
				var s sql.NullString
				colVar[i] = &s
			default:
				var s interface{}
				colVar[i] = &s
			}
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			log.Fatalln(scanErr)
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			switch typeVal[j].DatabaseTypeName() {
			case "INT":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "TINYINT":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "MEDIUMINT":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "SMALLINT":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "BIGINT":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "FLOAT":
				temp := *(colVar[j].(*sql.NullFloat64))
				if temp.Valid {
					result[col[j]] = temp.Float64
				} else {
					result[col[j]] = nil
				}
			case "DOUBLE":
				temp := *(colVar[j].(*sql.NullFloat64))
				if temp.Valid {
					result[col[j]] = temp.Float64
				} else {
					result[col[j]] = nil
				}
			case "DECIMAL":
				temp := *(colVar[j].(*sql.NullInt64))
				if temp.Valid {
					result[col[j]] = temp.Int64
				} else {
					result[col[j]] = nil
				}
			case "DATE":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "TIME":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "YEAR":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "DATETIME":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "TIMESTAMP":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "VARCHAR":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "MEDIUMTEXT":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "LONGTEXT":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "TINYTEXT":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			case "TEXT":
				temp := *(colVar[j].(*sql.NullString))
				if temp.Valid {
					result[col[j]] = temp.String
				} else {
					result[col[j]] = nil
				}
			default:
				result[col[j]] = colVar[j]
			}
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		log.Fatalln(err)
		panic(err)
	}
	return results, rs
}

func Exec(query string, args ...interface{}) sql.Result {
	rs, err := SqlDB.Exec(query, args...)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	return rs
}

func BeginTransactionsByLevel() *sqlTx {

	//LevelDefault IsolationLevel = iota
	//LevelReadUncommitted
	//LevelReadCommitted
	//LevelWriteCommitted
	//LevelRepeatableRead
	//LevelSnapshot
	//LevelSerializable
	//LevelLinearizable

	tx, err := SqlDB.BeginTx(context.Background(),
		&sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		panic(err)
	}
	SqlTx.Tx = tx
	return &SqlTx
}

func BeginTransactions() *sqlTx {
	tx, err := SqlDB.BeginTx(context.Background(),
		&sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		panic(err)
	}
	SqlTx.Tx = tx
	return &SqlTx
}

func (SqlTx *sqlTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := SqlDB.Exec(query, args...)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return rs, nil
}

func (SqlTx *sqlTx) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rs, err := SqlTx.Tx.Query(query, args...)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		log.Fatalln(colErr)
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			// Tips: string类型如果为interface返回乱字符串
			if typeVal[i].ScanType().Name() == "RawBytes" {
				var s string
				colVar[i] = &s
			} else {
				var s interface{}
				colVar[i] = &s
			}
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			log.Fatalln(scanErr)
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			result[col[j]] = colVar[j]
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return results, nil
}

func setType(colVar []interface{}, i int, typeName string)  {
	switch typeName {
	case "INT":
		var s int
		colVar[i] = &s
	case "TINYINT":
		var s int
		colVar[i] = &s
	case "MEDIUMINT":
		var s int
		colVar[i] = &s
	case "SMALLINT":
		var s int
		colVar[i] = &s
	case "BIGINT":
		var s int
		colVar[i] = &s
	case "FLOAT":
		var s float32
		colVar[i] = &s
	case "DOUBLE":
		var s float32
		colVar[i] = &s
	case "DECIMAL":
		var s int
		colVar[i] = &s
	case "DATE":
		var s string
		colVar[i] = &s
	case "TIME":
		var s string
		colVar[i] = &s
	case "YEAR":
		var s string
		colVar[i] = &s
	case "DATETIME":
		var s string
		colVar[i] = &s
	case "TIMESTAMP":
		var s string
		colVar[i] = &s
	case "VARCHAR":
		var s string
		colVar[i] = &s
	case "MEDIUMTEXT":
		var s string
		colVar[i] = &s
	case "LONGTEXT":
		var s string
		colVar[i] = &s
	case "TINYTEXT":
		var s string
		colVar[i] = &s
	case "TEXT":
		var s string
		colVar[i] = &s
	default:
		var s interface{}
		colVar[i] = &s
	}
}