package config

import (
	"html/template"
)

// app config holds the application config.

type AppConfig struct {
	TemplateCache map[string]*template.Template
}
