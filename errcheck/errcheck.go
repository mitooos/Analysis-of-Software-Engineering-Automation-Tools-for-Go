package errcheck

import "fmt"

func returnsError() error {
	return fmt.Errorf("errorr")
}

func ignoresError() {
	returnsError()
}

func ignoresErrorButAssigns() {
	_ = returnsError()
}

func ignoresFirstError() {
	err := returnsError()
	err = returnsError()
	if err != nil {
		panic(err)
	}
}
