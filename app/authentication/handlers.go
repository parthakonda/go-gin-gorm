package authentication

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListAnimalView(c *gin.Context) {
	DB.AutoMigrate(&Animal{})
	c.JSON(200, listAnimal())
}

func CreateAnimalView(c *gin.Context) {
	var animal Animal
	c.BindJSON(&animal)
	animal = createAnimal(animal)
	c.JSON(201, animal)
}

func GetAnimalView(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, getAnimal(id))
}

func UpdateAnimalView(c *gin.Context) {
	var animal Animal
	id := c.Params.ByName("id")
	if err := DB.Where("id = ?", id).First(&animal).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.BindJSON(&animal)
	DB.Save(&animal)
	c.JSON(201, animal)
}

func DeleteAnimalView(c *gin.Context) {
	id := c.Params.ByName("id")
	var animal Animal
	d := DB.Where("id = ?", id).Delete(&animal)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
