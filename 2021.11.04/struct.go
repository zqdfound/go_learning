package main

import "time"

type Employee struct {
	ID     int
	NAME   string
	AGE    int
	BIRTH  time.Time
	SALARY float32
}

func main() {
	var sam Employee
	sam.SALARY = 5000

	age := &sam.AGE
	*age = *age + 1

}
