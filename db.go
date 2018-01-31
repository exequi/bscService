package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func dbmain() {
	db, err := sql.Open("sqlite3", "./bs.db")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM user_account")
	checkErr(err)

	for rows.Next() {
		var id string
		var account string
		var password string
		var createtime string
		err = rows.Scan(&id, &account, &password, &createtime)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(account)
		fmt.Println(password)
		fmt.Println(createtime)
	}

	fmt.Println("#####")
	delete(db)

	fmt.Println("#####")
	update(db)

	fmt.Println("#####")
	insert(db)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func delete(db *sql.DB) {

	//删除数据
	stmt, err := db.Prepare("delete from user_account where id=?")
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func update(db *sql.DB) {
	stmt, err := db.Prepare("update  user_account set account=? where id=?")
	checkErr(err)
	res, err := stmt.Exec("gosqlite", "a4142d9efe6411e785deade1db9ff619")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func insert(db *sql.DB) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO user_account(id, account, password) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("goid", "goaccount", "gopassword")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
