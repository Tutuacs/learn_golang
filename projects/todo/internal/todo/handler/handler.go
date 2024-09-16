package todo

import "todo-app/internal/todo"

func TesteHandlerTodo() {

	db, err := todo.NewStore()
	if err != nil {
		return
	}

	db.UserCreate()

	return
}
