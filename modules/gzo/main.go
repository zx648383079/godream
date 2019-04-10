package gzo

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	fmt.Println(path)
}
