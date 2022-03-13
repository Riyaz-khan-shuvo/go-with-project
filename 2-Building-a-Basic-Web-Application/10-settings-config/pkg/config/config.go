package config

import (
	"html/template"
	"log"
)

// app config holds the application config.

type AppConfig struct {
	useCache      bool
	TemplateCache map[string]*template.Template
	infoLog       *log.Logger
}
