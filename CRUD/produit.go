package crud

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	models "rendu-examen/modeles"
	"rendu-examen/utils"
	"strconv"
	"strings"
)

func AjouterProduit() {
	reader := bufio.NewReader(os.Stdin)
	var produit models.Produit

	fmt.Print("Titre: ")
	produit.Titre, _ = reader.ReadString('\n')
	produit.Titre = strings.TrimSpace(produit.Titre)

	fmt.Print("Description: ")
	produit.Description, _ = reader.ReadString('\n')
	produit.Description = strings.TrimSpace(produit.Description)

	fmt.Print("Prix: ")
	fmt.Scanln(&produit.Prix)

	fmt.Print("Quantité: ")
	fmt.Scanln(&produit.Quantite)

	query := `INSERT INTO produits (titre, description, prix, quantite, actif) VALUES (?, ?, ?, ?, TRUE)`
    res, err := utils.BD.Exec(query, produit.Titre, produit.Description, produit.Prix, produit.Quantite)
    if err != nil {
        fmt.Println("Erreur lors de l'ajout du produit:", err)
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        fmt.Println("Erreur lors de la récupération de l'ID du produit:", err)
        return
    }
    produit.ID = int(id)

    fmt.Println("Produit ajouté avec succès:", produit)
}

func AfficherProduits() {
	var produits []models.Produit
	err := utils.BD.Select(&produits, "SELECT * FROM produits WHERE actif = TRUE")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des produits:", err)
		return
	}

	for _, produit := range produits {
		fmt.Printf("ID: %d, Titre: %s, Description: %s, Prix: %.2f, Quantité: %d\n", produit.ID, produit.Titre, produit.Description, produit.Prix, produit.Quantite)
	}
}

func ModifierProduit() {
	var produit models.Produit
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID du produit à modifier: ")
	fmt.Scanln(&produit.ID)

	fmt.Print("Nouveau Titre: ")
	produit.Titre, _ = reader.ReadString('\n')
	produit.Titre = strings.TrimSpace(produit.Titre)

	fmt.Print("Nouvelle Description: ")
	produit.Description, _ = reader.ReadString('\n')
	produit.Description = strings.TrimSpace(produit.Description)

	fmt.Print("Nouveau Prix: ")
	fmt.Scanln(&produit.Prix)

	fmt.Print("Nouvelle Quantité: ")
	fmt.Scanln(&produit.Quantite)

    query := `UPDATE produits SET titre = ?, description = ?, prix = ?, quantite = ? WHERE id = ?`
	_, err := utils.BD.Exec(query, produit.Titre, produit.Description, produit.Prix, produit.Quantite, produit.ID)
	if err != nil {
		fmt.Println("Erreur lors de la modification du produit:", err)
		return
	}

	fmt.Println("Produit modifié avec succès:", produit)
}

func DesactiverProduit() {
	var id int
	fmt.Print("ID du produit à désactiver: ")
	fmt.Scanln(&id)

    query := `UPDATE produits SET actif = FALSE WHERE id = ?`
	_, err := utils.BD.Exec(query, id)
	if err != nil {
		fmt.Println("Erreur lors de la désactivation du produit:", err)
		return
	}

	fmt.Println("Produit désactivé avec succès")
}

func ExporterProduits() {
	var produits []models.Produit
	err := utils.BD.Select(&produits, "SELECT * FROM produits")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des produits:", err)
		return
	}

	file, err := os.Create("exports/produits.csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
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

	fmt.Println("Produits exportés avec succès")
}
