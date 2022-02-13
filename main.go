package main

import ( 
	"fmt"
	"github.com/gofiber/fiber/v2"
	"devred.io/todolist/models"
	"devred.io/todolist/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=devred dbname=gotodo port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DATABASE !")
	}
	fmt.Println("Database Connected Succefully")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App){
	app.Get("/todos", models.GetTodos)
} 

func main() {
  app := fiber.New()
  initDatabase()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  setupRoutes(app)
  app.Listen(":3000")
}