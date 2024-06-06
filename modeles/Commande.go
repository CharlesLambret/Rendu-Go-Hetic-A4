package models

import "time"

type Commande struct {
    ID        int       `db:"id" json:"id"`
    ClientID  int       `db:"client_id" json:"client_id"`
    ProduitID int       `db:"produit_id" json:"produit_id"`
    Quantite  int       `db:"quantite" json:"quantite"`
    Prix      float64   `db:"prix" json:"prix"`
    DateAchat time.Time `db:"date_achat" json:"date_achat"`
}
