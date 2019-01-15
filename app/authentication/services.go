package authentication

func createAnimal(animal Animal) Animal {
	DB.Create(&animal)
	return animal
}

func listAnimal() []Animal {
	var animals []Animal
	DB.Find(&animals)
	return animals
}

func getAnimal(id string) Animal {
	var animal Animal
	if err := DB.Where("id = ?", id).First(&animal).Error; err != nil {
		return animal
	}
	return animal
}
