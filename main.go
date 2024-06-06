package main

import (
	"fmt"
	"os"
	crud "rendu-examen/CRUD"
	"rendu-examen/utils"
)

func main() {
	utils.InitBD()

	for {
		afficherMenu()
		var choix int
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			crud.AjouterProduit()
		case 2:
			crud.AfficherProduits()
		case 3:
			crud.ModifierProduit()
		case 4:
			crud.DesactiverProduit()
		case 5:
			crud.ExporterProduits()
		case 6:
			crud.AjouterClient()
		case 7:
			crud.AfficherClients()
		case 8:
			crud.ModifierClient()
		case 9:
			crud.ExporterClients()
		case 10:
			crud.EffectuerCommande()
		case 11:
			crud.ExporterCommandes()
		case 0:
			fmt.Println("Au revoir!")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func afficherMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. Ajouter un produit")
	fmt.Println("2. Afficher l'ensemble des produits")
	fmt.Println("3. Modifier un produit")
	fmt.Println("4. Désactiver un produit")
	fmt.Println("5. Exporter l'ensemble des produits sous forme CSV")
	fmt.Println("6. Ajouter un client")
	fmt.Println("7. Afficher l'ensemble des clients")
	fmt.Println("8. Modifier un client")
	fmt.Println("9. Exporter l'ensemble des clients sous forme CSV")
	fmt.Println("10. Effectuer une commande")
	fmt.Println("11. Exporter l'ensemble des commandes")
	fmt.Println("0. Quitter")
	fmt.Print("Entrez votre choix: ")
}
