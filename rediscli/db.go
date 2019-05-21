package rediscli

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"../common"
)
// 创建数据库
func getDb()(*sql.DB,error){
	return sql.Open("sqlite3",common.CACHE_PATH+common.CACHE_DB)
}

// 插入数据
func Insert(){
	db,_ := getDb()
	defer db.Close()
}

// 更新数据
func Update(){
	db,_ := getDb()
	defer db.Close()
}

// 删除数据
func Delete(){
	db,_ := getDb()
	defer db.Close()

}

// 查询数据
func Query(){
	db,_ := getDb()
	defer db.Close()

}
