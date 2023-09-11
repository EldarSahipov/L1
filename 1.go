package main

import "fmt"

type Human struct {
	name    string
	surname string
}

func (h *Human) getName() string {
	return h.name
}

func (h *Human) getSurname() string {
	return h.surname
}

func (h *Human) setName(name string) {
	h.name = name
}

func (h *Human) setSurname(surname string) {
	h.surname = surname
}

type Action struct {
	name string
	Human
}

func (a *Action) getName() string {
	return a.name
}

func (a *Action) setName(name string) {
	a.name = name
}

func main() {
	// Создание объекта Action и заполнение полей
	action := Action{
		name: "Work",
		Human: Human{
			name:    "Karlo",
			surname: "Kapitanov",
		},
	}
	// Вывод имени действия и имени человека из структуры Action
	fmt.Println(action.name)       // При обращении к одинаковым полям структур, приоретит выше у наследника
	fmt.Println(action.Human.name) // Явное обращение к полю встроенной структуры

	fmt.Println(action.surname)       // Так как поля нет в структуре Action, а есть у встроенной структуры, то обращается к полю встроенной структуры
	fmt.Println(action.Human.surname) // Явное обращение к полю встроенной структуры

	fmt.Println(action.getName())       // Обращение к методу структуры Action
	fmt.Println(action.Human.getName()) // Обращение к методу структуры Human

	action.setSurname("Otto") // Так как метода нет в структуре Action, а есть у встроенной структуры, то обращается к методу встроенной структуры
	fmt.Println(action.surname)
	fmt.Println(action.Human.surname)
}
