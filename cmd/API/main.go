package main

import (
	"log"
	"net/http"

	"main.go/Features/Add"
	"main.go/Features/del"
	"main.go/Features/done"
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

	//создаем все необходимое для list хендлера
	listRepository := list.NewRepository(db)
	listService := list.NewService(listRepository)
	listHandler := list.NewHandler(listService)

	//создаем все необходимое для del хендлера
	delRepository := del.NewRepository(db)
	delService := del.NewService(delRepository)
	DelHandler := del.NewHandler(delService)

	//создаем все необходимое для Done хендлера
	donerepository := done.NewRepository(db)
	doneService := done.NewService(donerepository)
	doneHandler := done.NewHandler(doneService)

	//подключаем http хендлеры к маршрутам
	httpDoneHandler := http.HandlerFunc(doneHandler.DoneHandler)
	httpDelHandler := http.HandlerFunc(DelHandler.DeleteHandler)
	httpListHandler := http.HandlerFunc(listHandler.ListHandler)
	httpAddHandler := http.HandlerFunc(addhandler.AddHandler)
	httpHelpHandler := http.HandlerFunc((help.HelpHandler))

	//создаем сервер
	srv := server.New(
		cfg.Port,
		httpAddHandler,
		httpHelpHandler,
		httpListHandler,
		httpDelHandler,
		httpDoneHandler,
	)

	//запускаем сервер
	log.Printf("сервер запускается на порту: %s", cfg.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка при запуске сервера: %s", err)
	}
}
