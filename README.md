# Rendu-Go-Hetic-A4
Ce projet permet de gérer des produits, des clients et des commandes par ligne de commandes. Chaque client peut commander un seul produit avec une quantité spécifiée. Après chaque commande, un email de confirmation avec la commande en pdf est envoyé au client.

### Prérequis
-   Docker
-   Docker Compose
-   Go 
-   phpMyadmin

### Installation

git clone https://github.com/CharlesLambret/Rendu-Go-Hetic-A4

## Lancer le projet 

#### 1. Démarrer Docker 
docker-compose up -d

#### 2. Ajouter les utilisateurs
users.sh

#### 3. Installer les dépendances
go mod tidy

#### 4. Importer la base de données via le fichier tp-go.sql