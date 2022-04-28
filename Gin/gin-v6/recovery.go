package gin_v6

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// trace 打印堆栈跟踪信息用于调试
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// Recovery 恐慌恢复中间件
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			// 恢复 panic 并打印堆栈跟踪信息
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
