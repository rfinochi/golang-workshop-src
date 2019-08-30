package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var repositoryType string

func createItemEndpoint(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)
	
	repo := createRepository()
	repo.CreateItem(newItem)
	
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func getItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("num"))

	repo := createRepository()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, repo.GetItem(id))
}

func getItemsEndpoint(c *gin.Context) {
	 repo := createRepository()
	 c.JSON(http.StatusOK, repo.GetItems())
}

func updateItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("num"))

	var updatedItem Item
	c.BindJSON(&updatedItem)
	
	repo := createRepository()
	updatedItem.ID = id
	repo.UpdateItem(updatedItem)
	
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func deleteItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("num"))

	repo := createRepository()
	repo.DeleteItem(id)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func createRepository() TodoRepository {
	if (repositoryType == "Mongo"){
		return &MongoRepository{}
	} else {
		return &InMemory{}
	}
}

func main() {
	repositoryType = "Mongo"

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/todo")

	api.GET("/", getItemsEndpoint)
	api.GET("/:num", getItemEndpoint)
	api.POST("/", createItemEndpoint)
	api.PATCH("/:num", updateItemEndpoint)
	api.DELETE("/:num", deleteItemEndpoint)

	router.Run(":8080")
}
