package errgroup

import (
	"context"
	"fmt"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	g, _ := WithContext(context.Background())
	var a, b, c []int
	g.Go(func() error {
		a = append(a, 1)
		return nil
	})

	g.Go(func() error {
		b = append(b, 2)
		return nil
	})

	g.Go(func() error {
		c = append(c, 3)
		return nil
	})

	err := g.Wait()
	fmt.Println(err)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
