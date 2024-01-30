package main

import (
	"database/sql"

	"net/http"

	"fmt"

	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *sql.DB

/*var products = []Product{
	{Name: "Pizza Margherita", Price: 15.99},
	{Name: "Hamburguer Clássico", Price: 8.99},
	{Name: "Salada Caesar", Price: 7.49},
	{Name: "Sushi Misto", Price: 18.99},
	{Name: "Spaghetti Bolognese", Price: 12.99},
}*/

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, 1)
}

func welcompage(products, c *gin.Context) {
	mesa := c.Param("mesa")
	//c.HTML(http.StatusOK, "index.html", gin.H{"content": "Index page..."})
	//c.IndentedJSON(http.StatusOK, products)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Se for o JSON
		c.JSON(http.StatusOK, products)
	default:
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "INDEX PAGE",
			"mesa":  mesa,
		})
	}
}

func connectMysql() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 ("root"),
		Passwd:               ("senha1"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3307",
		DBName:               "pira",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT descpro01, prevend01 from cadpro ORDER BY descpro01;")
	if err != nil {
		log.Fatal(err)
	}

	var products []Product

	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("Todos os produtos: %v", err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Todos produtos: %v", err)
	}
	return products, nil
}

var products []Product = getAllProducts()

func main() {
	connectMysql()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static/")

	router.GET("/", welcompage)

	router.GET("/:mesa", welcompage)

	router.Run("localhost:8080")
}
