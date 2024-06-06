package crud

import (
	"encoding/csv"
	"net/http"
	"os"
	"rendu-examen/modeles"
	"rendu-examen/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AjouterClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO clients (nom, prenom, telephone, adresse, email) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := utils.BD.QueryRow(query, client.Nom, client.Prenom, client.Telephone, client.Adresse, client.Email).Scan(&client.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func AfficherClients(c *gin.Context) {
	var clients []models.Client
	err := utils.BD.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func ModifierClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE clients SET nom = $1, prenom = $2, telephone = $3, adresse = $4, email = $5 WHERE id = $6`
	_, err := utils.BD.Exec(query, client.Nom, client.Prenom, client.Telephone, client.Adresse, client.Email, client.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func ExporterClients(c *gin.Context) {
	var clients []models.Client
	err := utils.BD.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create("exports/clients.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Nom", "Prenom", "Telephone", "Adresse", "Email"})

	for _, client := range clients {
		writer.Write([]string{
			strconv.Itoa(client.ID),
			client.Nom,
			client.Prenom,
			client.Telephone,
			client.Adresse,
			client.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clients exportés avec succès"})
}
