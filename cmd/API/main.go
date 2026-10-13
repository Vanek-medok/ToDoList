package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"main.go/Features/Add"
	"main.go/Features/del"
	"main.go/Features/done"
	"main.go/Features/help"
	"main.go/Features/list"
	"main.go/core/config"
	"main.go/core/database"
	"main.go/core/logger"
	"main.go/core/server"
)

func main() {

	logFile, err := logger.Init()
	if err != nil {
		log.Fatalf("Не удалось создать лог файл: %v", err)
	}
	defer logFile.Close()

	slog.Info("Запуск приложения")

	//создаем конфиг
	cfg := config.Load()

	slog.Info("Покдлючение к базе данных")
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Не удалось подключится к базе данных:", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	slog.Info("Подключение к БД успешно")

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

	// //запускаем сервер
	// log.Printf("сервер запускается на порту: %s", cfg.Port)

	// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	log.Fatalf("Ошибка при запуске сервера: %s", err)
	// }

	slog.Info("Запуск HTTP сервера", "addr", srv.Addr)

	if err = runServer(srv); err != nil {
		slog.Error("Ошибка HTTP сервера", "error", err)
		os.Exit(1)
	}

}

func runServer(srv *http.Server) error {
	serverErr := make(chan error, 1)

	go func() {
		serverErr <- srv.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	select {
	case err := <-serverErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		log.Println("Получен сигнал завершения сервера")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}

		log.Println("Сервер корректно завершен")
		return nil
	}
}
