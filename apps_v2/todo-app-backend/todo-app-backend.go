package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

var todos = map[string]*Todo{
	"88422cbc-90b0-11eb-a8b3-0242ac130003": {
		ID:   "88422cbc-90b0-11eb-a8b3-0242ac130003",
		Name: "Get a haircut",
	},
	"8d284fd6-90b0-11eb-a8b3-0242ac130003": {
		ID:   "8d284fd6-90b0-11eb-a8b3-0242ac130003",
		Name: "Get a real job",
	},
}

func main() {
	port := getEnvOrDefault("PORT", "5678")

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.Default()) // ignore cors for now

	r.GET("/todos", getAllTodos())
	r.POST("/todos", createTodo())

	log.Printf("Server starting in port %s", port)
	r.Run(":" + port)
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type TodosResponse struct {
	Todos []Todo `json:"todos"`
}

func getAllTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		values := []Todo{}
		for _, value := range todos {
			values = append(values, *value)
		}
		response := TodosResponse{
			Todos: values,
		}
		c.JSON(http.StatusOK, response)
	}
}

type NewTodo struct {
	Name string `json:"name" binding:"required"`
}

type NewTodoResponse struct {
	Todo Todo `json:"todo"`
}

func createTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo NewTodo

		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo := Todo{
			ID:   uuid.NewString(),
			Name: newTodo.Name,
		}

		todos[todo.ID] = &todo

		response := NewTodoResponse{
			Todo: todo,
		}

		c.JSON(http.StatusOK, response)
	}
}
