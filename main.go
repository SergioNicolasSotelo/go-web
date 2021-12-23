package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

func main() {

	router := gin.Default()

	router.GET("/hola/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.JSON(200, gin.H{
			"message": "Hola, " + nombre,
		})
	})
	/*-------------------------------------------------*/

	prods, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	productos := Productos{}
	err = json.Unmarshal(prods, &productos)
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/productos", func(c *gin.Context) {
		c.JSON(200, productos)
	})

	router.Run()
}
