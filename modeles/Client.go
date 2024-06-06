package models

type Client struct {
    ID        int    `db:"id" json:"id"`
    Nom       string `db:"nom" json:"nom"`
    Prenom    string `db:"prenom" json:"prenom"`
    Telephone string `db:"telephone" json:"telephone"`
    Adresse   string `db:"adresse" json:"adresse"`
    Email     string `db:"email" json:"email"`
}
