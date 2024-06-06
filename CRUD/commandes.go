package crud

import (
	"encoding/csv"
	"fmt"
	"os"
	models "rendu-examen/modeles"
	"rendu-examen/utils"
	"strconv"
	"time"
)

func EffectuerCommande() {
	var commande models.Commande

	fmt.Print("ID du client: ")
	fmt.Scanln(&commande.ClientID)

	fmt.Print("ID du produit: ")
	fmt.Scanln(&commande.ProduitID)

	fmt.Print("Quantité: ")
	fmt.Scanln(&commande.Quantite)

	var produit models.Produit
    produitquery := "SELECT * FROM produits WHERE id = ?"
    err := utils.BD.Get(&produit, produitquery, commande.ProduitID)
    if err != nil {
        fmt.Println("Erreur lors de la récupération du produit:", err)
        return
    }

    commande.Prix = produit.Prix * float64(commande.Quantite)

    query := `INSERT INTO commandes (client_id, produit_id, quantite, prix, date_achat) VALUES (?, ?, ?, ?, ?)`
    res, err := utils.BD.Exec(query, commande.ClientID, commande.ProduitID, commande.Quantite, commande.Prix, time.Now())
    if err != nil {
        fmt.Println("Erreur lors de la création de la commande:", err)
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        fmt.Println("Erreur lors de la récupération de l'ID de la commande:", err)
        return
    }
    commande.ID = int(id)

	nouvelleQuantite := produit.Quantite - commande.Quantite
    updateProduitQuery := `UPDATE produits SET quantite = ? WHERE id = ?`
    _, err = utils.BD.Exec(updateProduitQuery, nouvelleQuantite, produit.ID)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour de la quantité du produit:", err)
        return
    }

    var client models.Client
    clientquery := "SELECT * FROM clients WHERE id = ?"
    err = utils.BD.Get(&client, clientquery, commande.ClientID)
    if err != nil {
        fmt.Println("Erreur lors de la récupération du client:", err)
        return
    }
    utils.EnvoyerEmail(client.Email, "Confirmation de commande", fmt.Sprintf("Vous avez commandé %d x %s pour %.2f EUR le %x", commande.Quantite, produit.Titre, commande.Prix, commande.DateAchat))

    fmt.Println("Commande effectuée avec succès:", commande)
}

func ExporterCommandes() {
	var commandes []models.Commande
	err := utils.BD.Select(&commandes, "SELECT * FROM commandes")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des commandes:", err)
		return
	}

	file, err := os.Create("exports/commandes.csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
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

	fmt.Println("Commandes exportées avec succès")
}
