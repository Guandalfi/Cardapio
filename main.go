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
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	SubGroupID   int     `json:"groupid"`
	SubGroupName string  `json:"groupname"`
	ClassID      int     `json:"classid"`
	ClassName    string  `json:"classname"`
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
	rows, err := db.Query("select a.descpro01, a.prevend01, b.codgss00, b.descgss00, COALESCE(c.codigo, 0), coalesce(c.descricao, 'sem classe') from cadpro a left join cadgss b on a.codgss01 = b.codgss00 left join tabcla c on a.classe01 = c.codigo where a.codfil01 = 1 and a.prevend01 > 0 and a.sitpro01 is null order by a.descpro01;")
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
		if err := rows.Scan(&pt.Name, &pt.Price, &pt.SubGroupID, &pt.SubGroupName, &pt.ClassID, &pt.ClassName); err != nil {
			return nil, fmt.Errorf("todos os produtos: %v", err)
		}
		products = append(products, pt)
	}
	log.Println(products)
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
