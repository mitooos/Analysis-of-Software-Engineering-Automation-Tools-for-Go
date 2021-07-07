package check

const unused = 1
const used = 2

type inefficient struct {
	bool1   bool
	string1 string
	int1    int32
	bool2   bool
	string2 string
	int2    int32
}

type efficient struct {
	bool1   bool
	bool2   bool
	int1    int32
	int2    int32
	string1 string
	string2 string
}

func newEfficient() *efficient {
	return &efficient{
		bool1:   true,
		bool2:   false,
		int1:    int32(used),
		int2:    int32(used),
		string1: "usedd",
		string2: "usedd",
	}
}
