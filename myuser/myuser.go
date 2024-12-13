package myuser

import "fmt"

type MyUser struct {
	Id   int
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func (u MyUser) String() string {
	return fmt.Sprintf("Информация о '%s':\nВозраст %d\n",
		u.Name, u.Age)
}

func (u *MyUser) SetNewName(newName string) {
	u.Name = newName
}
