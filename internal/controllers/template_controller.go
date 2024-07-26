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
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a plain-text password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}




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
	router.GET("login", ctl.doLogin)
	router.POST("login", ctl.postLogin)
	router.GET("register", ctl.doRegister)
	router.POST("register", ctl.postRegister)
}

/////////////////////////////////////////////////////////////////////


func (ctl *TemplateController) doRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", pongo2.Context{})
}

func (ctl *TemplateController) postRegister(c *gin.Context) {
	ctl.logger.Info("Register form SUBMITTED")
	username := c.PostForm("username")
	password := c.PostForm("password")
	securePassword, _ := HashPassword(password)
	newUser := databases.User{
		Username: username,
		Password: securePassword,
	}
	ctl.db.Create(&newUser)
	if ctl.db.Error != nil {
		ctl.logger.Error(ctl.db.Error)
	}
	c.HTML(http.StatusOK, "login.html", pongo2.Context{})
}

func (ctl *TemplateController) postLogin(c *gin.Context) {
	ctl.logger.Info("FORM SUBMITTED")
	username := c.PostForm("username")
	password := c.PostForm("password")
	ctl.logger.Info(fmt.Sprintf("username = %v password = %v", username, password))
	var user databases.User
	result := ctl.db.Where("username = ?", username).First(&user)
	if result.Error != nil || !CheckPasswordHash(password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	c.
	c.HTML(http.StatusOK, "test.html", pongo2.Context{})
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

func (ctl *TemplateController) doLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", pongo2.Context{})
}