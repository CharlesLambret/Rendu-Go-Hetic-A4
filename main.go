package main

import (
    "rendu-examen/CRUD"
    "rendu-examen/utils"
    "github.com/gin-gonic/gin"
)

func main() {
    utils.InitBD()

    r := gin.Default()

    r.POST("/produits", crud.AjouterProduit)
    r.GET("/produits", crud.AfficherProduits)
    r.PUT("/produits/:id", crud.ModifierProduit)
    r.DELETE("/produits/:id", crud.SupprimerProduit)
    r.GET("/exporter/produits", crud.ExporterProduits)

    r.POST("/clients", crud.AjouterClient)
    r.GET("/clients", crud.AfficherClients)
    r.PUT("/clients/:id", crud.ModifierClient)
    r.GET("/exporter/clients", crud.ExporterClients)

    r.POST("/commandes", crud.EffectuerCommande)
    r.GET("/exporter/commandes", crud.ExporterCommandes)

    r.Run(":8080")
}
