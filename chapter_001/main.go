package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

// MODELS
type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Pool *pgxpool.Pool

func InitDB() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	var err error
	Pool, err = pgxpool.New(context.Background(), "postgres://admin:secret@localhost:5432/studentdb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to DB")
}

// CONNECT TO DB
// func ConnectToDB() *pgxpool.Pool {
// 	dbUrl := "postgres://admin:secret@localhost:5432/studentdb"
// 	config, err := pgxpool.ParseConfig(dbUrl)
// 	config.MaxConns = 5

// 	if err != nil {
// 		log.Fatalf("Unable to parse config: %v", err)
// 	}
// 	pool, err := pgxpool.ConnectConfig(context.Background(), config)
// 	if err != nil {
// 		log.Fatalf("Unable to connect to database: %v\n", err)
// 	}
// 	fmt.Println("Connected to PostgreSQL with pool!")
// 	return pool
// }

// HANDLERS
func GetPing(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.String(http.StatusOK, fmt.Sprintf("pong from %s", hostname))
}

func GetStudents(c *gin.Context) {
	rows, err := Pool.Query(context.Background(), "SELECT id, name, email, created_at, updated_at FROM students LIMIT 10")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch data"})
		return
	}

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedAt, &student.UpdatedAt); err != nil {
			log.Printf("Error reading row: %v", err)
			continue
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating rows"})
		return
	}

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var studentRequest StudentRequest
	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inserting"})

	go func(student StudentRequest) {
		res, err := Pool.Exec(context.Background(),
			"INSERT INTO students (name, email) VALUES ($1, $2)",
			student.Name, student.Email,
		)
		if err != nil {
			log.Println("Insert failed", err)
		} else {
			log.Println("Inserted", res.RowsAffected())
		}
	}(studentRequest)

}

func CountStudents(c *gin.Context) {
	var count string
	err := Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM students").Scan(&count)
	if err != nil {
		log.Println("Count query failed:", err)
	}
	c.String(http.StatusOK, count)
}

func main() {
	InitDB()
	defer Pool.Close()
	router := gin.Default()

	router.GET("/ping", GetPing)
	router.GET("/students", GetStudents)
	router.GET("/students/count", CountStudents)
	router.POST("/students", CreateStudent)

	router.Run()
}
