package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/KroneXI/gncompany/internal/models"
)

// FileStorage реализует хранение метрик в файле
type FileStorage struct {
	filename string
	mu       sync.Mutex
}

// NewFileStorage создает новый экземпляр FileStorage
func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

// LoadMetrics загружает метрики из файла
func (fs *FileStorage) LoadMetrics() (*models.Metrics, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filename)
	if err != nil {
		// Если файл не существует, возвращаем пустые метрики
		return &models.Metrics{}, nil
	}
	defer closeFile(file)

	var metrics models.Metrics
	err = json.NewDecoder(file).Decode(&metrics)
	if err != nil {
		return nil, err
	}

	return &metrics, nil
}

// SaveMetrics сохраняет метрики в файл
func (fs *FileStorage) SaveMetrics(metrics *models.Metrics) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Create(fs.filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	return json.NewEncoder(file).Encode(metrics)
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}
