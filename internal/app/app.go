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
	cfg    *config.Config
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
		cfg:    cfg,
		server: server,
		lg:     lg,
	}, nil
}

func (a *App) Run() error {
	// TODO: Запуск HTTPS-сервера, пока проблемы в связи отсутствием доменного имени
	return a.runHTTP()
}

// runHTTPS запускает HTTPS-сервер
//
//nolint:unused
func (a *App) runHTTPS() error {
	certFile := "server.crt"
	keyFile := "server.key"

	if err := a.server.ListenAndServeTLS(certFile, keyFile); err != nil {
		return fmt.Errorf("Ошибка запуска HTTPS-сервера: %w", err)
	}

	return nil
}

// runHTTP запускает HTTP-сервер
func (a *App) runHTTP() error {
	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("Ошибка запуска HTTP-сервера: %w", err)
	}

	return nil
}
