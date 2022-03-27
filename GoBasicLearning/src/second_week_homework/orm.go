package second_week_homework

import "database/sql"

func QueryRows(s string) error {
	// 查找不到数据直接向上层调用方抛出Sentinel Error(sql.ErrNoRows)
	return sql.ErrNoRows
}
