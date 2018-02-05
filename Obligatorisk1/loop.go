package algorithms

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		fmt.Println(sig)
	}
}
