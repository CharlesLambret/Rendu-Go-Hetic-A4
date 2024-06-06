package utils

import (
    "fmt"
    "github.com/jung-kurt/gofpdf"
	"rendu-examen/modeles"
)

func GenererPDFCommande(commande models.Commande, client models.Client, produit models.Produit) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    pdf.Cell(40, 10, "Détails de la commande")
    pdf.Ln(12)
    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, fmt.Sprintf("ID Commande: %d", commande.ID))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("ID Client: %d", commande.ClientID))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("ID Produit: %d", commande.ProduitID))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Quantité: %d", commande.Quantite))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Prix: %.2f", commande.Prix))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Date d'achat: %s", commande.DateAchat.Format("2006-01-02 15:04:05")))	

    err := pdf.OutputFileAndClose(fmt.Sprintf("pdf/commande_%d.pdf", commande.ID))
    if err != nil {
        fmt.Println("Erreur lors de la génération du PDF:", err)
        return err
    }

    fmt.Println("PDF généré avec succès pour la commande:", commande.ID)
    return nil
}