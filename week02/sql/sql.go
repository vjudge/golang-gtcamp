package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strconv"
)

type User struct {
	id int
	name string
	nickname string
}

func InitClient() (db *sql.DB, err error)  {
	db, err = sql.Open("mysql", "rw_test:Txy_test168@tcp(rm-j6crf85vc749b3v6v4o.mysql.rds.aliyuncs.com:3306)/testgo?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "sql connect is fail")
		return
	}
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "sql link fail")
		return
	}

	return
}

func QueryUser(db *sql.DB, id int) (err error) {
	rows, err := db.Query("select * from user where id = " + strconv.Itoa(id))
	fmt.Println("err:", err)
	fmt.Println("rows:", rows)
	if err != nil {
		return errors.Wrap(err, "db query fail")
	}
	return nil
}