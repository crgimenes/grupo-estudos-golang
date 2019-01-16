package closer

import (
	"fmt"
	"io"
)

func Closer(a io.Closer) {
	err := a.Close()
	if err != nil {
		fmt.Println(err)
	}
}
