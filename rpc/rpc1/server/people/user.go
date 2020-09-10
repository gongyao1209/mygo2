package people

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (u User)getName() string {
	return u.Name
}

func FindUserById() *User {

	u := User{
		Id:1,
		Name:"aaa",
	}

	return &u
}