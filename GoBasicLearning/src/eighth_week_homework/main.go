package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var client redis.UniversalClient
var ctx context.Context

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	ctx = context.Background()
}

func main() {
	batchLoad(10000, "valueSize10_count10k", getSpecifiedString(10))
	batchLoad(50000, "valueSize10_count50k", getSpecifiedString(10))
	batchLoad(500000, "valueSize10_count500k", getSpecifiedString(10))

	batchLoad(10000, "valueSize1000_count10k", getSpecifiedString(1000))
	batchLoad(50000, "valueSize1000_count50k", getSpecifiedString(1000))
	batchLoad(500000, "valueSize1000_count500k", getSpecifiedString(1000))

	batchLoad(10000, "valueSize5000_count10k", getSpecifiedString(5000))
	batchLoad(50000, "valueSize5000_count50k", getSpecifiedString(5000))
	batchLoad(500000, "valueSize5000_count00k", getSpecifiedString(5000))

	analysis()
}

// getSpecifiedString 获取指定大小string
func getSpecifiedString(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}

// write 批量插入[key, val]
func batchLoad(num int, key, value string) {
	for i := 0; i < num; i++ {
		// 生成不同key值
		k := fmt.Sprintf("%s:%v", key, i)
		// 插入[key, val]对
		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

// 分析 info memory
func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		panic(err)
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
