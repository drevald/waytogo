package controllers

import (
	"embed"
	"github.com/samber/do"	
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)


//go:embed static/*
var staticFS embed.FS
type StaticController struct {

}
func NewStatic (di *do.Injector) (*StaticController, error) {
	return &StaticController{}, nil
}

func (ctl *StaticController) Wire(router *gin.Engine) {
	router.Use(static.Serve("/", static.EmbedFolder(staticFS, "static")))
}