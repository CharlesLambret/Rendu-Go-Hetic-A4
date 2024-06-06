package models

type Produit struct {
    ID          int     `db:"id" json:"id"`
    Titre       string  `db:"titre" json:"titre"`
    Description string  `db:"description" json:"description"`
    Prix        float64 `db:"prix" json:"prix"`
    Quantite    int     `db:"quantite" json:"quantite"`
    Actif       bool    `db:"actif" json:"actif"`
}
