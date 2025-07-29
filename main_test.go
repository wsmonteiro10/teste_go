package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"teste_go/handlers"
	"teste_go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Client-Secret")
	r.Use(cors.New(config))

	r.GET("/ping", handlers.Ping)

	r.POST("/rapido", middleware.ClientSecretMiddleware(), handlers.UploadFile)
	r.POST("/ultrarapido", middleware.ClientSecretMiddleware(), handlers.UltraUploadFile)
	r.GET("/arquivos", middleware.ClientSecretMiddleware(), handlers.ListUploadFiles)

	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, `{"message":"pong"}`, resp.Body.String())
}

func TestRapidoWithoutSecret(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/rapido", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}

func TestRapidoWithSecret(t *testing.T) {
	router := setupRouter()

	body := strings.NewReader("dummy content")
	req, _ := http.NewRequest("POST", "/rapido", body)
	req.Header.Set("Client-Secret", "minha-senha-secreta")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.NotEqual(t, 401, resp.Code)
}

func TestUltraRapidoWithoutSecret(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/ultrarapido", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}

func TestUltraRapidoWithSecret(t *testing.T) {
	router := setupRouter()

	body := strings.NewReader("dummy content")
	req, _ := http.NewRequest("POST", "/ultrarapido", body)
	req.Header.Set("Client-Secret", "minha-senha-secreta")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.NotEqual(t, 401, resp.Code)
}
