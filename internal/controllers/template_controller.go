package controllers

import (
	"embed"
	"github.com/ddreval/waytogo/internal/databases"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gitlab.com/go-box/pongo2gin/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
	"fmt"
)

//go:embed views/*
var viewsFS embed.FS

type TemplateController struct {
	logger *logrus.Logger
	templ *pongo2.TemplateSet
	db    *gorm.DB
}

func NewTemplate(di *do.Injector) (*TemplateController, error) {
	templFS, err := fs.Sub(viewsFS, "views")
	if err != nil {
		return nil, err
	}
	db, err := do.Invoke[*gorm.DB](di)
	if err != nil {
		return nil, err
	}
	templ := pongo2.NewSet("", pongo2.NewFSLoader(templFS))
	logger, err := do.Invoke[*logrus.Logger](di)
	if err != nil {
		return nil, err
	}
	return &TemplateController{logger, templ, db}, nil
}

func (ctl *TemplateController) Wire(router *gin.Engine) {	
	fmt.Println("Fmt Wire")
	ctl.logger.Debug("Wire")
	router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateSet: ctl.templ,
	})
	router.GET("test", ctl.doTest)
	router.GET("users", ctl.doUsers)
}

func (ctl *TemplateController) doTest(c *gin.Context) {
	ctl.logger.Error("do test")
	ctl.logger.Warn("do test")
	ctl.logger.Info("do test")
	ctl.logger.Debug("do test")
	fmt.Println("fmt do test")
	c.HTML(http.StatusOK, "test.html", pongo2.Context{})
}

func (ctl *TemplateController) doUsers(c *gin.Context) {
	ctl.logger.Debug("do users")
	var users []databases.User
    result := ctl.db.Find(&users)
    if result.Error != nil {
        fmt.Printf("failed to retrieve users: %v", result.Error)
    }
	c.HTML(http.StatusOK, "users.html", pongo2.Context{"users":users})
}