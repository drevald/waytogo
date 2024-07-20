package controllers

import (
	"embed"
	"github.com/flosh/pongo2/v5"
)

//go:embed views/*
var viewsFS embed.FS
type TemplateController struct {
	templ *pongo2.TemplateSet
}