package models

import "service/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM todos WHERE id=$1`

	// * Use row when get complex responses
	row := conn.QueryRow(sql, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return

}
