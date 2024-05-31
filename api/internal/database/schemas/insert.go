package main

import (
	"api/common/initializer"
	"api/internal/database"
	"api/internal/models"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main() {
	InserBroths()
	InserProteins()
	insertRecipes()
}

func InserBroths() {
	database.DB.Create(&models.Broth{
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "Salt",
		Description:   "Simple like the seawater, nothing more",
		Price:         10,
	})
	database.DB.Create(&models.Broth{
		ImageInactive: "https://tech.redventures.com.br/icons/miso/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/miso/active.svg",
		Name:          "Miso",
		Description:   "Paste made of fermented soybeans",
		Price:         12,
	})
	database.DB.Create(&models.Broth{
		ImageInactive: "https://tech.redventures.com.br/icons/shoyu/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/shoyu/active.svg",
		Name:          "Shoyu",
		Description:   "The good old and traditional soy sauce",
		Price:         10,
	})
}

func InserProteins() {
	database.DB.Create(&models.Protein{
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "Chasu",
		Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
		Price:         10,
	})
	database.DB.Create(&models.Protein{
		ImageInactive: "https://tech.redventures.com.br/icons/yasai/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/yasai/active.svg",
		Name:          "Yasai Vegetarian",
		Description:   "A delicious vegetarian lamen with a selection of season vegetables.",
		Price:         10,
	})
	database.DB.Create(&models.Protein{
		ImageInactive: "https://tech.redventures.com.br/icons/chicken/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/chicken/active.svg",
		Name:          "Karaague",
		Description:   "Three units of fried chicken, moyashi, ajitama egg and other vegetables.",
	})
}

func insertRecipes() {

	database.DB.Create(&models.Recipe{
		Name: "Salt and Chasu Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Salt and Yasai Vegetarian Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Salt and Karaague Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Shoyu and Chasu Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Shoyu and Yasai Vegetarian Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Shoyu and Karaague Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Miso and Chasu Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Miso and Yasai Vegetarian Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png",
	})

	database.DB.Create(&models.Recipe{	
		Name: "Miso and Karaague Ramen",
		Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png",
	})
}