package benchmark

type slice []int

func (s slice) traverseReference(action func(int)) error {
	for i := range s {
		action(s[i])
	}
	return nil
}

func (s slice) traverseValue(action func(int)) error {
	for _, element := range s {
		action(element)
	}
	return nil
}
