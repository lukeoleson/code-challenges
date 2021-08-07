package main

import (
	"fmt"
	"github.com/lukeoleson/code_signal/intro"
)

func main() {
	ans := intro.Permute([]string{"A", "B", "C"})
	fmt.Println(ans)
}