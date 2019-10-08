package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库实体
type User struct {
	ID      int
	Name    string
	Age     int
	Sex     int
	AddDate time.Time
}

func main() {
	//1、获取数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")
	checkErr(err)
	defer db.Close()
	fmt.Println("数据库连接成功")
	//2、判断连接是否有效
	err = db.Ping()
	checkErr(err)
	fmt.Println("数据库连接有效")
	//3、创建表
	sql := `
        CREATE TABLE IF NOT EXISTS users(
            id INT NOT NULL AUTO_INCREMENT,
            name VARCHAR(100) NOT NULL,
            age INT NOT NULL,
            sex TINYINT,
            add_date DATETIME,
            PRIMARY KEY(id)
        )
    `
	_, err = db.Exec(sql)
	checkErr(err)

	//4、添加数据
	sql = "INSERT INTO users (name,age,sex,add_date) VALUES (?,?,?,?)"
	res, err := db.Exec(sql, "张三", 18, 1, time.Now())
	checkErr(err)
	fmt.Println(res.LastInsertId())

	//5、查询数据
	sql = "SELECT id,name,age,sex,add_date FROM users"
	rows, err := db.Query(sql)
	checkErr(err)
	defer rows.Close()
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
		checkErr(err)
		fmt.Println(user, user.AddDate.Format("2006/01/02 15:04:05"))
	}
	err = rows.Err()
	checkErr(err)

	//6、查询一行
	sql = "SELECT id,name,age,sex,add_date FROM users"
	row := db.QueryRow(sql)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	checkErr(err)
	fmt.Println(user)

	//7、命令 PrepareStatement 预编译SQL
	sql = "UPDATE users SET name=? WHERE id=?;"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	defer stmt.Close()
	result, err := stmt.Exec("李四", 1)
	checkErr(err)
	fmt.Println(result.RowsAffected())

	//8、查询
	sql = "SELECT id,name,age,sex,add_date FROM users WHERE id=?"
	stmt2, err := db.Prepare(sql)
	checkErr(err)
	defer stmt2.Close()
	row = stmt2.QueryRow(1)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	checkErr(err)
	fmt.Println(user)

	//9、事务
	tx, err := db.Begin()
	checkErr(err)
	_, err = tx.Exec("UPDATE users SET age=? WHERE id=?", 20, 1)
	checkTxErr(err, tx)
	_, err = tx.Exec("UPDATE users SET sex=? WHERE id=?", 1, 1)
	checkTxErr(err, tx)
	err = tx.Commit()
	checkTxErr(err, tx)

	//10、查询一行
	sql = "SELECT id,name,age,sex,add_date FROM users"
	row = db.QueryRow(sql)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	checkErr(err)
	fmt.Println(user)

	//11、删除表
	sql = "DROP TABLE users"
	_, err = db.Exec(sql)
	checkErr(err)
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 事务错误
func checkTxErr(err error, tx *sql.Tx) {
	if err != nil {
		log.Println(err)
		err = tx.Rollback()
		checkErr(err)
	}
}
