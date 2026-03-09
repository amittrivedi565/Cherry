package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

func GetInstance() *Database {
	dsn := "root@tcp(localhost:3306)/cherry?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Conn: db,
	}
}

func (db *Database) ExecuteQuery(op string, query string) (interface{}, error) {

	switch op {

	case "select":

		rows, err := db.Conn.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		var results []map[string]interface{}

		for rows.Next() {

			values := make([]interface{}, len(columns))
			valuePtr := make([]interface{}, len(values))

			for i := range values {
				valuePtr[i] = &values[i]
			}

			err := rows.Scan(valuePtr...)
			if err != nil {
				return nil, err
			}

			row := make(map[string]interface{})

			for i, col := range columns {
				val := values[i]

				if b, ok := val.([]byte); ok {
					row[col] = string(b)
				} else {
					row[col] = val
				}
			}

			results = append(results, row)
		}

		return results, nil

	default:

		result, err := db.Conn.Exec(query)
		if err != nil {
			return nil, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"rows_affected": rowsAffected,
		}, nil
	}
}
