package gocyclo

import "fmt"

func function1(i int) {
	if i < 0 {
		fmt.Println("negative")
	} else if i < 10 {
		fmt.Println("between 0 and 9 inclusive")
	} else if i < 20 {
		fmt.Println("between 10 and 19 inclusive")
	} else {
		fmt.Println("20 or more")
	}
}

func function2(i int) {
	switch {
	case i < 0:
		fmt.Println("negative")
	case i < 10:
		fmt.Println("between 0 and 9 inclusive")
	case i < 20:
		fmt.Println("between 10 and 19 inclusive")
	default:
		fmt.Println("20 or more")
	}
}

func function3(numbers []int) {
	for _, i := range numbers {
		switch {
		case i < 0:
			fmt.Println("negative")
		case i < 10:
			fmt.Println("between 0 and 9 inclusive")
		case i < 20:
			fmt.Println("between 10 and 19 inclusive")
		default:
			fmt.Println("20 or more")
		}
	}
}
