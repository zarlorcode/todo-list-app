package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Estructura de la tarea
type ToDo struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *gorm.DB

// Inicializar la base de datos
func initDB() {
	var err error

	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	// Obtener las variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construir la cadena de conexión
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	// Migrar la estructura de la tabla
	db.AutoMigrate(&ToDo{})
}

// Obtener todas las tareas
func getToDos(c *gin.Context) {
	var toDoList []ToDo
	db.Find(&toDoList)
	log.Println("GET /todos - Obtener todas las tareas")
	c.JSON(http.StatusOK, toDoList)
}

// Crear una nueva tarea
func createToDo(c *gin.Context) {
	var newToDo ToDo
	if err := c.BindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("POST /todos - Error al crear la tarea: JSON inválido")
		return
	}
	db.Create(&newToDo)
	log.Printf("POST /todos - Nueva tarea creada: %+v\n", newToDo)
	c.JSON(http.StatusCreated, newToDo)
}

// Actualizar una tarea
func updateToDo(c *gin.Context) {
	id := c.Param("id")
	var updatedToDo ToDo
	if err := c.BindJSON(&updatedToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Printf("PUT /todos/%s - Error al actualizar la tarea: JSON inválido\n", id)
		return
	}

	var existingToDo ToDo
	if result := db.First(&existingToDo, "id = ?", id); result.Error != nil {
		log.Printf("PUT /todos/%s - Tarea no encontrada\n", id)
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	existingToDo.Title = updatedToDo.Title
	existingToDo.Status = updatedToDo.Status
	db.Save(&existingToDo)
	log.Printf("PUT /todos/%s - Tarea actualizada: %+v\n", id, existingToDo)
	c.JSON(http.StatusOK, existingToDo)
}

// Eliminar una tarea
func deleteToDo(c *gin.Context) {
	id := c.Param("id")
	if result := db.Delete(&ToDo{}, id); result.RowsAffected == 0 {
		log.Printf("DELETE /todos/%s - Tarea no encontrada\n", id)
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	log.Printf("DELETE /todos/%s - Tarea eliminada\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func main() {
	// Inicializar la base de datos
	initDB()

	// Configurar el archivo de log
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file) // Redirigir logs al archivo

	router := gin.Default()

	// Rutas
	router.GET("/todos", getToDos)
	router.POST("/todos", createToDo)
	router.PUT("/todos/:id", updateToDo)
	router.DELETE("/todos/:id", deleteToDo)

	// Servir la aplicación
	router.Run(":8080")
}


