package myuser

import "fmt"

type MyUser struct {
	Name                  string
	Age                   uint16
	Money                 uint16
	Avg_grades, Happiness float64
}

func (u MyUser) String() string {
	return fmt.Sprintf("Информация о '%s'\nВозраст %d \nБаланс: %d \nУровень счастья: %f",
		u.Name, u.Age, u.Money, u.Happiness)
}

func (u *MyUser) SetNewName(newName string) {
	u.Name = newName
}
