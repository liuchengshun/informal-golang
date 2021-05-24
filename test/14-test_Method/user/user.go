package user

type User struct {
	Name string
}

func (u *User) YourName() string {
	return u.Name
}
