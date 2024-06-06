package crud

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"rendu-examen/modeles"
	"rendu-examen/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func EffectuerCommande(c *gin.Context) {
	var commande models.Commande
	if err := c.ShouldBindJSON(&commande); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO commandes (client_id, produit_id, quantite, prix, date_achat) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := utils.BD.QueryRow(query, commande.ClientID, commande.ProduitID, commande.Quantite, commande.Prix, time.Now()).Scan(&commande.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, commande)

}

func ExporterCommandes(c *gin.Context) {
	var commandes []models.Commande
	err := utils.BD.Select(&commandes, "SELECT * FROM commandes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create("exports/commandes.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "ClientID", "ProduitID", "Quantite", "Prix", "DateAchat"})

	for _, commande := range commandes {
		writer.Write([]string{
			strconv.Itoa(commande.ID),
			strconv.Itoa(commande.ClientID),
			strconv.Itoa(commande.ProduitID),
			strconv.Itoa(commande.Quantite),
			fmt.Sprintf("%.2f", commande.Prix),
			commande.DateAchat.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Commandes exportées avec succès"})
}
