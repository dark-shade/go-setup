package utils

import (
	"fmt"
	"os"
)

// CheckDebug prints the debug msg with the prefix 'Debug:'.
func CheckDebug(msg interface{}) {
	fmt.Fprintln(os.Stdout, "Debug:", msg)
}
