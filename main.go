package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

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

var productList []Productos
var idIncrement int

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
	/*-------------------------------------------------*/

	pr := router.Group("/product")

	pr.POST("/", Guardar())

	router.Run()
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Productos
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		idIncrement++
		req.Id = idIncrement
		if validar(req) == true {
			productList = append(productList, req)
		} else {
			c.JSON(404, gin.H{
				"error": "Uno o más campos inválidos.",
			})
		}
		c.JSON(200, req)
	}
}
func getField(v interface{}, name string) (interface{}, error) {
	// v debe ser un puntero a una estructura
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return nil, errors.New("v debe ser un puntero a una estructura")
	}

	// Obtengo el valor subyacente al puntero
	// Textual de la documentacion de reflect:
	// Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	rv = rv.Elem()

	// Obtengo el campo de la estruct a partir de su nombre
	fv := rv.FieldByName(name)
	// Verifico que el campo exista dentro de mi estructura
	if !fv.IsValid() {
		return nil, fmt.Errorf("%s no existe en la estructura", name)
	}

	// Si el campo no está exportado, no deberiamos poder acceder a el
	if !fv.CanSet() {
		return nil, fmt.Errorf("no es posible acceder al campo %s", name)
	}

	// Si el valor es el zero value de su tipo, devolvemos un error
	if fv.IsZero() {
		return nil, fmt.Errorf("el campo %s esta vacio", name)
	}

	// retornamos el valor del campo, y un error nulo
	return fv, nil
}

func validar(p Productos) bool {
	listaDeCamposRequeridos := []string{"Id", "Nombre", "Color", "Precio", "Stock", "Codigo", "Publicado", "FechaCreacion"}

	for _, nombreDelCampo := range listaDeCamposRequeridos {

		_, err := getField(&p, nombreDelCampo)
		if err != nil {
			fmt.Printf("%s", err)
			fmt.Printf("el valor del campo %s es requerido", nombreDelCampo)
			return false

		}
	}
	return true
}
