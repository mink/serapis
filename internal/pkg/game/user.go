package game

type User struct {
	Name string
	SSO  string
}

var Users []*User

func init() {
	Users = append(Users, &User{
		Name: "Zac",
		SSO:  "lol123",
	})
	Users = append(Users, &User{
		Name: "Test",
		SSO:  "test",
	})
}
