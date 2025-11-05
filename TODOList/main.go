package main

import (
	todo "TODOList/TODO"
	"TODOList/http"
	"log"
)

func main() {
	/*
		Стартуем сервак, даем ему кастомные хендлеры
	*/

	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		log.Fatalln("Server ne rabotaet:(", err)
	}
}
