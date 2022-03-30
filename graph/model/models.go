package model

type Ownable interface {
	Owner() *User
	SetOwner(*User)
}

type Todo struct {
	ID    int
	Text  string
	Done  bool
	owner *User
}

func (t Todo) Owner() *User {
	return t.owner
}

func (t *Todo) SetOwner(user *User) {
	t.owner = user
}

type User struct {
	ID   int
	Name string
}
