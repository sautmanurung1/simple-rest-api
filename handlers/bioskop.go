package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"simple-rest-api/models"
)

type BioskopHandler struct {
	DB *sql.DB
}

func NewBioskopHandler(db *sql.DB) *BioskopHandler {
	return &BioskopHandler{DB: db}
}

func (h *BioskopHandler) AddBioskop(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if b.Nama == "" && b.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3)`
	_, err := h.DB.Exec(query, b.Nama, b.Lokasi, b.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan"})
}
