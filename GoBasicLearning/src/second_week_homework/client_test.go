package second_week_homework

import (
	"database/sql"
	"fmt"
	"testing"
)

// TestSqlErrorNoRows 这里相当于是DAO层和应用层一个简单的约定
func TestSqlErrorNoRows(t *testing.T) {
	var name string
	sqlStr := "select name from users where id = ?"
	err := QueryRows(sqlStr)
	if err != nil {
		// 应用层再针对sql.ErrNoRows做逻辑处理
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			// 应用层在此处直接将该错误处理掉即可(错误降级)
			// 打印错误相关日志
			fmt.Printf(`execute SQL: '%s', meet error: %v\n`, sqlStr, err)
		} else {
			panic(err)
		}
	}
	fmt.Println(name)
}
