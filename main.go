package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlages := NewCmdFlag()
	cmdFlages.Execute(&todos)
	storage.Save(todos)
}
