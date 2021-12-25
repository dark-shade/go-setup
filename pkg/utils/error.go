package utils

import (
	"fmt"
	"os"
)

// CheckErrFatal prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErrFatal(err interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// CheckErrNonFatal prints the msg with the prefix 'Error:' but does not exit the code. If the msg is nil, it does nothing.
func CheckErrNonFatal(err interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
