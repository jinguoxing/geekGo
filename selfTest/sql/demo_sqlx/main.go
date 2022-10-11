package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"

	"fmt"
)

type dbObject struct {

	db *sqlx.DB
}


func(d *dbObject)Open() *sqlx.DB {

	var err error
	d.db ,err = sqlx.Open("mysql","root:123456@tcp(127.0.0.1:3306)/testdb")

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

	Id int  `db:"id"`
	Title string `db:"title"`
	content string `db:"content"`
}

func GetOne(){

	dbc := dbObject{}
	db := dbc.Open()
	defer db.Close()

	var art artInfo

	err := db.Get(&art,"select id,title from articles where id = ?",1)
	if err != nil {
		fmt.Printf("获取数据失败,err:%v",err)
		return
	}

	fmt.Println(art.Id,art.Title)
}


//func GetAll(){
//
//	dbc := dbObject{}
//	db := dbc.Open()
//	defer db.Close()
//
//	stmt ,_ := db.Prepare("select id,title from articles where id >= ?")
//
//	rows,_ := stmt.Query(1)
//
//	defer rows.Close()
//
//	art := &artInfo{}
//
//	for rows.Next(){
//		err := rows.Scan(&art.id,&art.title)
//
//		if err != nil {
//			fmt.Printf("获取数据失败,err:%v",err)
//			return
//		}
//		fmt.Println(art.id,art.title)
//	}
//}



func main(){

	//Insert()

	GetOne()
	//GetAll()
}



func GetNowTime() time.Time  {

	location,_:= time.LoadLocation("Asia/Shanghai")

	return time.Now().In(location)

}