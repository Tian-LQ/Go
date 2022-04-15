package rolling

import (
	"sync"
	"time"
)

// Number 统计10秒内的 请求数量|成功响应数量|失败响应数量|超时数量|拒绝请求数量
// Number 滑动窗口 保存着10秒内的请求数据信息(单bucket保存秒级数据信息)
type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

// numberBucket 保存每秒的请求数据信息
type numberBucket struct {
	Value float64
}

// NewNumber 初始化方法
func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

// getCurrentBucket 获取当前这一秒的bucket
func (r *Number) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	// 若窗口中没有这一秒的bucket则添加，若存在则返回
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

// removeOldBuckets 删除当前滑动窗口中十秒前的过期bucket
func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		// TODO: configurable rolling window
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

// Increment 为当前(秒)bucket增加统计数据信息i
func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	r.removeOldBuckets()
}

// UpdateMax 更新当前bucket当中的value值为max(value, n)
func (r *Number) UpdateMax(n float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	// 更新当前bucket统计的value
	// 若n大于value则更新为n，否则不变
	b := r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.removeOldBuckets()
}

// Sum 统计整个滑动窗口当中十秒内的bucket数据value总和
func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	// 筛选出近10秒内的bucket
	for timestamp, bucket := range r.Buckets {
		// TODO: configurable rolling window
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}

	return sum
}

// Max 返回整个滑动窗口当中十秒内bucket的最大value
func (r *Number) Max(now time.Time) float64 {
	var max float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	// 筛选出近10秒内的bucket
	for timestamp, bucket := range r.Buckets {
		// TODO: configurable rolling window
		if timestamp >= now.Unix()-10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

// Avg 统计滑动窗口十秒内bucket的value平均值
func (r *Number) Avg(now time.Time) float64 {
	return r.Sum(now) / 10
}
