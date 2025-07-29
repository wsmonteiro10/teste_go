package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile permite upload de arquivo .txt ou .csv grande e salva no servidor
// @Summary ETL rápido: Upload de TXT ou CSV e grava no Postgres
// @Description Recebe arquivos grandes via multipart/form-data, processa e salva no banco de forma eficiente
// @Tags ETL Rápido
// @Security ClientSecret
// @securityDefinitions.apikey ClientSecret
// @in header
// @name Client-Secret
// @Accept multipart/form-data
// @Param file formData file true "Arquivo para upload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /rapido [post]
func UploadFile(c *gin.Context) {

	start := time.Now()

	// Limita o tamanho máximo do upload em 5GB (ajuste conforme desejar)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 5<<30) // 5 GB

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err = os.Mkdir("uploads", 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pasta uploads: " + err.Error()})
			return
		}
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao receber arquivo: " + err.Error()})
		return
	}
	defer file.Close()

	filename := header.Filename
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != ".txt" && ext != ".csv" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de arquivo não suportado. Use .txt ou .csv"})
		return
	}

	savePath := filepath.Join("uploads", filename)
	out, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar arquivo: " + err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo: " + err.Error()})
		return
	}

	end := time.Now()

	c.JSON(http.StatusOK, gin.H{
		"message":     fmt.Sprintf("Arquivo %s enviado com sucesso!", filename),
		"execucao_ms": fmt.Sprintf("%.2f", end.Sub(start).Seconds()*1000),
	})
}

// ultraUploadFile permite upload de arquivo .txt ou .csv grande e salva no servidor
// @Summary ETL ultrarrápido: Upload de TXT ou CSV direto no Postgres sem índice
// @Description Insere dados em massa sem índice para máxima performance
// @Tags ETL Rápido
// @Security ClientSecret
// @securityDefinitions.apikey ClientSecret
// @in header
// @name Client-Secret
// @Accept multipart/form-data
// @Param file formData file true "Arquivo para upload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /ultrarapido [post]
func UltraUploadFile(c *gin.Context) {
	// Limita o tamanho máximo do upload em 5GB (ajuste conforme desejar)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 5<<30) // 5 GB

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao receber arquivo: " + err.Error()})
		return
	}
	defer file.Close()

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err = os.Mkdir("uploads", 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pasta uploads: " + err.Error()})
			return
		}
	}

	filename := header.Filename
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != ".txt" && ext != ".csv" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de arquivo não suportado. Use .txt ou .csv"})
		return
	}

	savePath := filepath.Join("uploads", filename)
	out, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar arquivo: " + err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Arquivo %s enviado com sucesso!", filename)})
}

// Retorna todos os arquivos da pasta uploads
// @Summary Retorna todos os arquivos da pasta uploads
// @Description Retorna todos os arquivos da pasta uploads
// @Tags Arquivos
// @Security ClientSecret
// @securityDefinitions.apikey ClientSecret
// @in header
// @name Client-Secret
// @Produce json
// @Success 200 {object} []map[string]string
// @Router /listas_arquivos [get]
func ListUploadFiles(c *gin.Context) {
	files, err := os.ReadDir("uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler pasta uploads: " + err.Error()})
		return
	}

	var list []map[string]string
	for _, file := range files {
		list = append(list, map[string]string{
			"name": file.Name(),
		})
	}

	c.JSON(http.StatusOK, list)
}
