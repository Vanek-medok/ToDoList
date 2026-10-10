package main

import (
	"log"
	"net/http"

	"main.go/Features/Add"
	"main.go/Features/help"
	"main.go/Features/list"
	"main.go/core/config"
	"main.go/core/database"
	"main.go/core/server"
)

func main() {

	//создаем конфиг
	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//создаем репозиторий
	repository := Add.NewRepository(db)

	//создаем сервис
	service := Add.NewService(repository)

	//создаем хендлер
	addhandler := Add.NewHandler(service)

	//создаем репо
	listRepository := list.NewRepository(db)

	//создаем сервис
	listService := list.NewService(listRepository)

	//создаем хендлер
	listHandler := list.NewHandler(listService)

	//подключаем http хендлеры к маршрутам
	httpListHandler := http.HandlerFunc(listHandler.ListHandler)
	httpAddHandler := http.HandlerFunc(addhandler.AddHandler)
	httpHelpHandler := http.HandlerFunc((help.HelpHandler))

	//создаем сервер
	srv := server.New(
		cfg.Port,
		httpAddHandler,
		httpHelpHandler,
		httpListHandler,
	)

	//запускаем сервер
	log.Printf("сервер запускается на порту: %s", cfg.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка при запуске сервера: %s", err)
	}
}
