package main

import (
	"database/sql"

	"net/http"

	"fmt"

	"log"

	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *sql.DB

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, 1)
}

func welcompage(c *gin.Context, products []Product) {
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
	rows, err := db.Query("SELECT descpro01, prevend01 from cadpro where prevend01 > 0 and codfil01 = 1 ORDER BY descpro01;")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var products []Product

	defer rows.Close()

	// Configurando o log fora do loop
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("log.txt", flags, 0666)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Redirecting logs to the file
	log.SetOutput(file)

	for rows.Next() {
		var pt Product
		if err := rows.Scan(&pt.Name, &pt.Price); err != nil {
			return nil, fmt.Errorf("todos os produtos: %v", err)
		}
		products = append(products, pt)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("todos produtos: %v", err)
	}

	return products, nil
}

func main() {
	connectMysql()
	allProducts, _ := getAllProducts()
	fmt.Println(allProducts)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static/")

	router.GET("/", func(c *gin.Context) {
		welcompage(c, allProducts)
	})

	router.GET("/:mesa", func(c *gin.Context) {
		welcompage(c, allProducts)
	})

	router.Run("localhost:8080")
}
