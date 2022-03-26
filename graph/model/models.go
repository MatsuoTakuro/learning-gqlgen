package model

type Ownable interface {
	Owner() *User
}

type Todo struct {
	ID    int
	Text  string
	Done  bool
	Owner *User
}
type User struct {
	ID   int
	Name string
}
