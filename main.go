package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

// *********************** Hotel data mock ***********************
var hoteis = []hotel{
	{
		ID:   1,
		Nome: "Hotel Teste A",
		EndComercial: endComercial{
			Logradouro:  "Rua X",
			Numero:      123,
			Bairro:      "Centro",
			Complemento: "Sala 2",
			Cidade:      "São Paulo",
			Estado:      "SP",
			CEP:         "01000-000",
		},
	},
	{
		ID:   2,
		Nome: "Hotel Teste B",
		EndComercial: endComercial{
			Logradouro:  "Av Y",
			Numero:      456,
			Bairro:      "Jardins",
			Complemento: "Apto 10",
			Cidade:      "São Paulo",
			Estado:      "SP",
			CEP:         "02000-000",
		},
	},
}

// *********************** Structs ***********************//
// 'hotel' represents data about a record hotel
type hotel struct {
	ID           int64        `json:"id" xml:"id" binding:"required"`
	Nome         string       `json:"nome" xml:"nome"`
	EndComercial endComercial `json:"end_comercial" xml:"end_comercial"`
}

// 'endComercial' represents data about the specific hotel commercial address
type endComercial struct {
	Logradouro  string `json:"logradouro" xml:"logradouro" binding:"required"`
	Bairro      string `json:"bairro" xml:"bairro" binding:"required"`
	Numero      int64  `json:"numero" xml:"numero"`
	Complemento string `json:"complemento" xml:"complemento"`
	Cidade      string `json:"cidade" xml:"cidade" binding:"required"`
	Estado      string `json:"estado" xml:"estado" binding:"required"`
	CEP         string `json:"cep" xml:"cep" binding:"required"`
}

type hotelList struct {
	XMLName xml.Name `xml:"hoteis"`
	Hoteis  []hotel  `xml:"hotel"`
}

// *********************** Endpoints ***********************
// getHoteisJSON respond with the list of all hotels as JSON
func getHoteisJSON(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hoteis)
}

// getHoteisXML respond with the list of all hotels as XML
func getHoteisXML(c *gin.Context) {
	response := hotelList{Hoteis: hoteis}
	c.XML(http.StatusOK, response)
}

// postHotelXML adds a hotel form JSON received in the request body
func postHotelJSON(c *gin.Context) {
	var newHotel hotel

	// Checking the JSON format
	if err := c.BindJSON(&newHotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hoteis = append(hoteis, newHotel)
	c.IndentedJSON(http.StatusCreated, newHotel)
}

// postHotelXML adds a hotel form XLM received in the request body
func postHotelXML(c *gin.Context) {
	var newHotel hotel

	// Checking the XML format
	if err := c.BindXML(&newHotel); err != nil {
		c.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hoteis = append(hoteis, newHotel)
	c.XML(http.StatusOK, newHotel)
}

// *********************** Main function ***********************
func main() {
	router := gin.Default()
	router.GET("/hoteis/json", getHoteisJSON)
	router.GET("/hoteis/xml", getHoteisXML)
	router.POST("/hoteis/json", postHotelJSON)
	router.POST("/hoteis/xml", postHotelXML)

	router.Run()
}
