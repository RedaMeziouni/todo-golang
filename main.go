package main

import ( 
	"fmt"
	"github.com/gofiber/fiber/v2"
	"devred.io/todolist/models"
	"devred.io/todolist/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdatedTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
} 

func main() {
  app := fiber.New()
  app.Use(cors.New())
  initDatabase()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  setupRoutes(app)
  app.Listen(":3000")
}