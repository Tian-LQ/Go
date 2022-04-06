package second_week_homework

import (
	"database/sql"
	"fmt"
	"strings"
)

var notFoundCode = 40001
var systemErr = 50001

func Biz2() error {
	err := Dao2("select * from user")

	if IsNoRow(err) {
		// 不管怎么说，出现了数据库查询的问题，可以转为业务领域错误，也可以继续向上传递
		fmt.Println(err.Error())
	} else if err != nil {

	}
	return nil
}

func Dao2(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		// 在这一步封装好查询参数，这样DEBUG就能知道请求什么数据没找到
		// 同时带上了堆栈信息方便定位
		// 我们没有仔细区别err是什么，反正就是告诉上游，出错了
		return fmt.Errorf("%d, not found", notFoundCode)
	} else if err != nil {
		return fmt.Errorf("%d, not found", systemErr)
	}
	// do something
	return nil
}

func mockError() error {
	return sql.ErrNoRows
}

func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}
