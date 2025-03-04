package main

import (
	"fmt"
	"time"
)

const u string = "New"

func vals() (int, int) {
	return 3, 7
}
func main() {

	s := "Sagar"
	var t string = "Agrawal"
	fmt.Println("hello world", s, t, u)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	// for i := range 5 {
	for i := 0; i < 5; i++ {

		fmt.Printf("%d ", i)
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday", time.Now().Weekday())
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	c, f := vals()
	fmt.Println(c, f)

	// panic_defer()

	// panic("a problem")

	// string_help()

	time_func()
	// _, err := os.Create("/tmp/file")
	// if err != nil {
	// 	panic(err)
	// }

}
