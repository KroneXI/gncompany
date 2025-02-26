package main

import (
	"log"
	"os"

	"github.com/KroneXI/gncompany/internal/logger"

	"github.com/KroneXI/gncompany/internal/app"
)

func main() {
	// Инициализация приложения
	zapLogger := logger.New()
	appInstance, err := app.NewApp(zapLogger)
	if err != nil {
		log.Fatalf("Ошибка инициализации приложения: %v", err)
	}

	// Запуск сервера
	if err = appInstance.Run(); err != nil {
		zapLogger.Fatal(err.Error())
		os.Exit(1)
	}
}
