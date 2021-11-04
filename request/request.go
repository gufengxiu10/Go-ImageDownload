package request

import (
	"fmt"
)

type Request struct{}

func (b Request) get() {
	fmt.Println(b)
}
