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

func AjouterClient() {
	reader := bufio.NewReader(os.Stdin)
	var client models.Client

	fmt.Print("Nom: ")
	client.Nom, _ = reader.ReadString('\n')
	client.Nom = strings.TrimSpace(client.Nom)

	fmt.Print("Prenom: ")
	client.Prenom, _ = reader.ReadString('\n')
	client.Prenom = strings.TrimSpace(client.Prenom)

	fmt.Print("Telephone: ")
	client.Telephone, _ = reader.ReadString('\n')
	client.Telephone = strings.TrimSpace(client.Telephone)

	fmt.Print("Adresse: ")
	client.Adresse, _ = reader.ReadString('\n')
	client.Adresse = strings.TrimSpace(client.Adresse)

	fmt.Print("Email: ")
	client.Email, _ = reader.ReadString('\n')
	client.Email = strings.TrimSpace(client.Email)

	query := `INSERT INTO clients (nom, prenom, telephone, adresse, email) VALUES (?, ?, ?, ?, ?)`
    res, err := utils.BD.Exec(query, client.Nom, client.Prenom, client.Telephone, client.Adresse, client.Email)
    if err != nil {
        fmt.Println("Erreur lors de l'ajout du client:", err)
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        fmt.Println("Erreur lors de la récupération de l'ID du client:", err)
        return
    }
    client.ID = int(id)

    fmt.Println("Client ajouté avec succès:", client)
}

func AfficherClients() {
	var clients []models.Client
	err := utils.BD.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des clients:", err)
		return
	}

	for _, client := range clients {
		fmt.Printf("ID: %d, Nom: %s, Prenom: %s, Telephone: %s, Adresse: %s, Email: %s\n", client.ID, client.Nom, client.Prenom, client.Telephone, client.Adresse, client.Email)
	}
}

func ModifierClient() {
	var client models.Client
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID du client à modifier: ")
	fmt.Scanln(&client.ID)

	fmt.Print("Nouveau Nom: ")
	client.Nom, _ = reader.ReadString('\n')
	client.Nom = strings.TrimSpace(client.Nom)

	fmt.Print("Nouveau Prenom: ")
	client.Prenom, _ = reader.ReadString('\n')
	client.Prenom = strings.TrimSpace(client.Prenom)

	fmt.Print("Nouveau Telephone: ")
	client.Telephone, _ = reader.ReadString('\n')
	client.Telephone = strings.TrimSpace(client.Telephone)

	fmt.Print("Nouvelle Adresse: ")
	client.Adresse, _ = reader.ReadString('\n')
	client.Adresse = strings.TrimSpace(client.Adresse)

	fmt.Print("Nouvel Email: ")
	client.Email, _ = reader.ReadString('\n')
	client.Email = strings.TrimSpace(client.Email)

	query := `UPDATE clients SET nom = ?, prenom = ?, telephone = ?, adresse = ?, email = ? WHERE id = ?`
    _, err := utils.BD.Exec(query, client.Nom, client.Prenom, client.Telephone, client.Adresse, client.Email, client.ID)
    if err != nil {
        fmt.Println("Erreur lors de la modification du client:", err)
        return
    }

    fmt.Println("Client modifié avec succès:", client)

}

func ExporterClients() {
	var clients []models.Client
	err := utils.BD.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des clients:", err)
		return
	}

	file, err := os.Create("exports/clients.csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
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

	fmt.Println("Clients exportés avec succès")
}
