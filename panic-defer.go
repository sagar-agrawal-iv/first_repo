package main

import "fmt"

func panic_later() {
	panic("Problem arises")
}

func panic_defer() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in the main", r)
		}
	}()
	fmt.Println("This is before panic in panic-defer")
	panic_later()
	fmt.Println("This is afer panic in panic-defer")
}
