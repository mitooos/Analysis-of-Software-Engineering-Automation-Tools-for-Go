package dupl

import "fmt"

func function1(i int) {
	switch {
	case i <= 10:
		fmt.Println("less than or equal 10")
	case i <= 20:
		fmt.Println("less than or equal 20")
	case i <= 30:
		fmt.Println("less than or equal 30")
	case i <= 40:
		fmt.Println("less than or equal 40")
	case i <= 50:
		fmt.Println("less than or equal 50")
	case i <= 60:
		fmt.Println("less than or equal 60")
	case i <= 70:
		fmt.Println("less than or equal 70")
	default:
		fmt.Println("default")
	}
}

func function2(i int) {
	switch {
	case i <= 80:
		fmt.Println("less than or equal 80")
	case i <= 90:
		fmt.Println("less than or equal 90")
	case i <= 100:
		fmt.Println("less than or equal 100")
	case i <= 110:
		fmt.Println("less than or equal 110")
	case i <= 120:
		fmt.Println("less than or equal 120")
	case i <= 130:
		fmt.Println("less than or equal 130")
	case i <= 140:
		fmt.Println("less than or equal 140")
	default:
		fmt.Println("default")
	}
}
