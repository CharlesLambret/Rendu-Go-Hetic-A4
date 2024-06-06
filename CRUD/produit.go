package crud

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	models "rendu-examen/modeles"
	"rendu-examen/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AjouterProduit(c *gin.Context) {
	var produit models.Produit
	if err := c.ShouldBindJSON(&produit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO produits (titre, description, prix, quantite) VALUES ($1, $2, $3, $4) RETURNING id`
	err := utils.BD.QueryRow(query, produit.Titre, produit.Description, produit.Prix, produit.Quantite).Scan(&produit.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produit)
}

func AfficherProduits(c *gin.Context) {
	var produits []models.Produit
	err := utils.BD.Select(&produits, "SELECT * FROM produits WHERE actif = TRUE")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produits)
}

func ModifierProduit(c *gin.Context) {
	var produit models.Produit
	if err := c.ShouldBindJSON(&produit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE produits SET titre = $1, description = $2, prix = $3, quantite = $4 WHERE id = $5`
	_, err := utils.BD.Exec(query, produit.Titre, produit.Description, produit.Prix, produit.Quantite, produit.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produit)
}

func SupprimerProduit(c *gin.Context) {
	id := c.Param("id")
	query := `UPDATE produits SET actif = FALSE WHERE id = $1`
	_, err := utils.BD.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produit désactivé"})
}

func ExporterProduits(c *gin.Context) {
	var produits []models.Produit
	err := utils.BD.Select(&produits, "SELECT * FROM produits")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create("exports/produits.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Titre", "Description", "Prix", "Quantite", "Actif"})

	for _, produit := range produits {
		writer.Write([]string{
			strconv.Itoa(produit.ID),
			produit.Titre,
			produit.Description,
			fmt.Sprintf("%.2f", produit.Prix),
			strconv.Itoa(produit.Quantite),
			strconv.FormatBool(produit.Actif),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produits exportés avec succès"})
}
