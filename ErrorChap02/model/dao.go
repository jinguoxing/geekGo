package model

import (
	"database/sql"
	"fmt"
	"errors"
	pkgerrors "github.com/pkg/errors"
)


var NotRowsDataFound = errors.New("Not found mysql row data")

func GetName(id int) (string, error) {
	var name string
	sqlQuery := "select name from it_demo where id = ?"
	err := DB.QueryRow(sqlQuery, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// ErrNoRows is sentinel error 表示一个空的结果集
			// 这里如果视作一个错误 可以warp出去
			// 但是不应该把sql.ErrNoRows warp出去 因为会导致上层必须依赖于database包
			// 可以自定义一个error
			return "", pkgerrors.Wrapf(NotRowsDataFound, fmt.Sprintf("mysql:%s,err:%v", sqlQuery,err))
		}
	}
	return name, err
}