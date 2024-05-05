package mock

import (
	"fmt"
	"io"
	"time"
)

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, "Go!")
}
