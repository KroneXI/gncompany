package handlers

import (
	"html/template"

	"github.com/KroneXI/gncompany/internal/logger"

	"net/http"
	"os"

	"github.com/KroneXI/gncompany/internal/config"
	"github.com/KroneXI/gncompany/internal/models"
	"github.com/KroneXI/gncompany/internal/storage"
)

// Handler содержит зависимости для обработчиков
type Handler struct {
	lg       *logger.Logger
	cfg      *config.Config
	metrics  *models.Metrics
	storage  *storage.FileStorage
	username string
	password string
}

// NewHandler создает новый экземпляр Handler
func NewHandler(cfg *config.Config, lg *logger.Logger) *Handler {
	// Инициализация хранилища метрик
	store := storage.NewFileStorage("metrics.json")
	metrics, err := store.LoadMetrics()
	if err != nil {
		lg.Errorf("Ошибка загрузки метрик: %v", err)
	}

	return &Handler{
		lg:       lg,
		cfg:      cfg,
		metrics:  metrics,
		storage:  store,
		username: cfg.AdminAUth.User,
		password: cfg.AdminAUth.Password,
	}
}

// Home обрабатывает главную страницу
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Читаем файл со стилями
	css, err := os.ReadFile("static/styles.css")
	if err != nil {
		http.Error(w, "Ошибка загрузки стилей", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, models.HomeResponse{
		Groups:      h.cfg.Server.User.Groups,
		PhoneNumber: h.cfg.Server.User.PhoneNumber,
		Styles:      template.CSS(css),
	})
	if err != nil {
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
	}

	h.metrics.Increment()
	if err = h.storage.SaveMetrics(h.metrics); err != nil {
		h.lg.Warnf("Ошибка сохранения метрик: %v", err)
	}
}

// Dashboard обрабатывает дашборд
func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != h.username || password != h.password {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, "Ошибка рендеринга дашборда", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, models.DashboardResponse{
		TotalVisits: h.metrics.GetTotalVisits(),
	})
	if err != nil {
		http.Error(w, "Ошибка рендеринга дашборда", http.StatusInternalServerError)
	}
}
