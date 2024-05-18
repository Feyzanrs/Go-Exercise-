package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Monday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on weekend")
	default:
		fmt.Println("Born some other day")
	}

}
