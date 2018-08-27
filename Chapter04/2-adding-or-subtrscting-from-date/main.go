package main

import (
	"time"
	"fmt"
)

func main(){
	current := time.Now()
	septDate := current.AddDate(1,1,0)

	fmt.Println(current.String())
	fmt.Println(septDate.String())

	// septDate.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))

	oneLessYears := septDate.AddDate(-1,0,0)
	fmt.Println(oneLessYears.String())

	tenMoreMinutes := septDate.Add(10 * time.Hour)
	fmt.Println(tenMoreMinutes)
}
