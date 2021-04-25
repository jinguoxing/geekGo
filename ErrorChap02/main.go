package main

import (
	"ErrorHandle/model"
	"fmt"
	"errors"
	"log"
)

/***

问题：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

解答：

ErrNoRows is sentinel error 表示一个空的结果集
这里如果视作一个错误 可以warp出去
但是不应该把sql.ErrNoRows warp出去 因为会导致上层必须依赖于database包
可以自定义一个error
Dao 层 通过 return errors.Wrapf(code.NotFound, fmt.Sprintf("mysql: %s error: %v", sql, err))

业务层操作:
if errors.Is(err, code.NotFound} {

}
 */


func main()  {

	name, err := model.GetName(21)

	if errors.Is(err, model.NotRowsDataFound) {
		// 判断Not Found 记录日志 返回信息
		log.Printf("%+v", err)
		return
	}

	// 获取成功
	fmt.Printf("name is %s\n", name)
}
