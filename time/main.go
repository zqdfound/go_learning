package main

import (
	"fmt"
)

//func main() {
//	s := gocron.NewScheduler(time.UTC)
//	s.Every(1).Seconds().Do(printTime)
//	s.StartAsync()
//	select {}
//}
//func printTime() {
//	fmt.Println(time.Now())
//}

func fullName(firstName *string, lastName *string) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic has been recovered")
		}
	}()
	if firstName == nil {
		panic("Firsr Name can't be null")
	}
	if lastName == nil {
		panic("Last Name can't be null")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func main() {
	//firName := "paul"
	//fullName(&firName, nil)
	//fmt.Println("returned normally from main")
	//var sb strings.Builder
	//sb.WriteString("a")
	var a int
	var b *int
	fmt.Printf("%T", a)
	fmt.Printf("%T", b)

}
