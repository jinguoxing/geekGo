package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"fmt"
)

type dbObject struct {

	db *sql.DB
}


func(d *dbObject)Open() *sql.DB {

	var err error
	d.db ,err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		fmt.Printf("数据库连接失败,err:%v", err)
	}
	d.db.SetMaxOpenConns(10)
	d.db.SetMaxIdleConns(5)
	return d.db
}

func (d *dbObject) Close() {

	d.Close()
}

func Insert(){

	dbc := dbObject{}
	db := dbc.Open()
	defer db.Close()

	res ,err := db.Exec("insert into articles (title,content,created_at,updated_at)values(?,?,?,?)",
		"Hello Redis","你好 Redis！",GetNowTime(),GetNowTime())

	if err!=nil{
		fmt.Printf("数据库插入失败,err:%v",err)
	}

	rowsaffected,err := res.RowsAffected()

	if err != nil {
		fmt.Printf("获取受影响行数失败,err:%v",err)
		return
	}
	fmt.Println("受影响行数:" ,rowsaffected )

}

type artInfo struct {

	id int
	title string
	content string

}

func GetOne(){

	dbc := dbObject{}
	db := dbc.Open()
	defer db.Close()

	stmt ,_ := db.Prepare("select id,title from articles where id = ?")

	row := stmt.QueryRow(1)

	art := &artInfo{}

	err := row.Scan(&art.id,&art.title)

	if err != nil {
		fmt.Printf("获取数据失败,err:%v",err)
		return
	}

	fmt.Println(art.id,art.title)
}


func GetAll(){

	dbc := dbObject{}
	db := dbc.Open()
	defer db.Close()

	stmt ,_ := db.Prepare("select id,title from articles where id >= ?")

	rows,_ := stmt.Query(1)

	defer rows.Close()

	art := &artInfo{}

	for rows.Next(){
		err := rows.Scan(&art.id,&art.title)

		if err != nil {
			fmt.Printf("获取数据失败,err:%v",err)
			return
		}
		fmt.Println(art.id,art.title)
	}
}



func main(){

	//Insert()

	//GetOne()
	GetAll()
}



func GetNowTime() time.Time  {

	location,_:= time.LoadLocation("Asia/Shanghai")

	return time.Now().In(location)

}