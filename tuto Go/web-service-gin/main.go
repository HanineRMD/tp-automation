package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album représente les données d'un album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums est notre "base de données" en mémoire
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Définition des routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// Démarrage du serveur
	router.Run("localhost:8080")
}

// getAlbums retourne tous les albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums ajoute un nouvel album
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Lie le JSON reçu à newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Ajoute le nouvel album à la slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID retourne un album spécifique par son ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Recherche l'album par ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
