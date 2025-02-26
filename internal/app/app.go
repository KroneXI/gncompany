package app

import (
	"fmt"
	"net"
	"net/http"

	"github.com/KroneXI/gncompany/internal/logger"

	"github.com/KroneXI/gncompany/internal/config"
	"github.com/KroneXI/gncompany/internal/handlers"
	"github.com/KroneXI/gncompany/internal/storage"
)

// App представляет основное приложение
type App struct {
	server *http.Server
	lg     *logger.Logger
}

// NewApp создает новый экземпляр приложения
func NewApp(lg *logger.Logger) (*App, error) {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("cfg/config.yaml")
	if err != nil {
		lg.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация хранилища метрик
	store := storage.NewFileStorage("metrics.json")
	_, err = store.LoadMetrics()
	if err != nil {
		return nil, err
	}

	// Инициализация обработчиков
	handler := handlers.NewHandler(cfg, lg)

	// Настройка маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/dashboard", handler.Dashboard)

	// Создание HTTP-сервера
	server := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", cfg.Port),
		Handler: mux,
	}

	return &App{
		server: server,
		lg:     lg,
	}, nil
}

// Run запускает HTTP-сервер
func (a *App) Run() error {
	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("Ошибка запуска сервера: %w", err)
	}

	return nil
}
