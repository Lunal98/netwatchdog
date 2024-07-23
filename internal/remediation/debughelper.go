package remediationhelper

import (
	"fmt"
	"time"
)

func GetDebugHelper() Helper {
	var helper debughelper
	return helper
}

type debughelper struct {
}

// ResetInterface implements Helper.
func (d debughelper) ResetInterface() {
	fmt.Println("restarting interface")
	for range 4 {
		time.Sleep(2 * time.Second)
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("restarted interface")
}

// Restart implements Helper.
func (d debughelper) Restart() {
	fmt.Println("restarting computer")
	for range 6 {
		time.Sleep(2 * time.Second)
		fmt.Print(".")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("restarted computer")
}
