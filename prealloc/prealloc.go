package prealloc

func getUsers() []*User {
	return nil
}

type User struct {
	Email string
	Name  string
}

func (*User) doSomething() {}

func inefficient() []*User {
	users := getUsers()
	var ans []*User
	for _, user := range users {
		user.doSomething()
		ans = append(ans, user)
	}

	return ans
}

func efficient() []*User {
	users := getUsers()
	ans := make([]*User, len(users))
	for _, user := range users {
		user.doSomething()
		ans = append(ans, user)
	}

	return ans
}
