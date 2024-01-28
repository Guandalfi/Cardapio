package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"Price"`
}

// albums slice to seed record album data.
var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99},
}

type product struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

var products = []product {
	{Name: "Pizza Margherita", Price: 15.99},
	{Name: "Hamburguer Clássico", Price: 8.99},
	{Name: "Salada Caesar", Price: 7.49},
	{Name: "Sushi Misto", Price: 18.99},
	{Name: "Spaghetti Bolognese", Price: 12.99},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func welcompage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"content": "Index page..."})
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID locates the album whose ID value matches th id parameter sent by the client, then returns tha algum as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for an album whose ID values matches the parameter.
	for _, a := range albums {		
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static/")

	router.GET("/", welcompage)

	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
