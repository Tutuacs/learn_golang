package models

import "service/db"

func List() (todos []Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM todos`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			// ! Ao ter problema em escanear um todo específico ele continua buscando os outros...
			continue
		}

		todos = append(todos, todo)
	}

	return

}
