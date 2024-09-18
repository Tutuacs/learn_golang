package utils

import (
	"fmt"

	"todo-app/pkg/logs"
)

func ErrorPanic(err error) {
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("ocorreu um erro inesperado: %v", err))
		panic(err)
	}
}
