package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	// UserService  model.UserService
	// TokenService model.TokenService
	MaxBodyBytes int64
}

type Config struct {
	R *gin.Engine
	// UserService     model.UserService
	// TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		// UserService:  c.UserService,
		// TokenService: c.TokenService,
		MaxBodyBytes: c.MaxBodyBytes,
	}

	// router for cors to be able to access from react
	c.R.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081", "*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "user_id", "user"},
	}))
	// Create an account group
	g := c.R.Group(c.BaseURL)

	// to test api
	// if gin.Mode() != gin.TestMode {

	// } else {

	// }

	g.GET("/", h.Home)
	g.POST("/receipt/process", h.ProcessReceipt)
	g.GET("receipt/:id/points", h.Points)

}

func (h *Handler) Home(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Its working": "kind of"})
}

func (h *Handler) Points(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Points": "processed"})
}
