package models

import (
	"html/template"

	"github.com/KroneXI/gncompany/internal/config"
)

type HomeResponse struct {
	Groups      []config.Group
	PhoneNumber string
	Styles      template.CSS
}
