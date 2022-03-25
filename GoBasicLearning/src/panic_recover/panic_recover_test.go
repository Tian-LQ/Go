package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovered form panic: [\"%v\"]\n", err)
		}
		fmt.Println("Finally")
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong"))
	//os.Exit(1)
}
