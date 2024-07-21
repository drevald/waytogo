package controllers

import (
	"github.com/ddreval/waytogo/internal/databases"
	"gitlab.com/go-box/pongo2gin/v5"
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2/v5"
	"github.com/samber/do"	
	"net/http"
	"io/fs"
	"embed"
)

//go:embed views/*
var viewsFS embed.FS

type TemplateController struct {
	templ *pongo2.TemplateSet
	db *databases.Database
}

func NewTemplate(di *do.Injector) (*TemplateController, error) {
	templFS, err := fs.Sub(viewsFS, "views")
	if err != nil {
	  return nil, err
	}
	db, err := do.Invoke[* databases.Database](di)
	if err != nil {
		return nil, err
	}
	templ := pongo2.NewSet("", pongo2.NewFSLoader(templFS))
	return &TemplateController{templ, db}, nil
}

func (ctl *TemplateController) Wire(router *gin.Engine) {
	router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateSet:ctl.templ,
	})
	router.GET("test", ctl.doTest)
}

func (ctl *TemplateController) doTest (c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", pongo2.Context{})
}