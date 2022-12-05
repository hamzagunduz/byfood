package main

import (
	"fmt"

	"database/sql"

	_ "strconv"

	"net/http"

	_ "errors"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"log"
)

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "byfood_user"
	dbpassword = "password"
	dbname     = "byfood_db"
)

type user struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Last_name string `json:"last_name"`
	Age       int    `json:"age"`
}

func dbConnection() (*sql.DB, error) {
	// Connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db, err
}

func createUser(c *gin.Context) {

	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("invalid JSON body: %v", err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With, Accept")

	db, err := dbConnection()

	if err != nil {
		panic(err)
	}

	sqlStatement := `INSERT INTO users(name, last_name, age) VALUES ($1, $2, $3)`

	result, err := db.Exec(sqlStatement, newUser.Name, newUser.Last_name, newUser.Age)

	if err != nil {
		panic(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		return
	}

	if n == 0 {
		e := "could not insert the record, please try again after sometime"
		log.Println(e)
		return
	}

	defer db.Close()

	if err != nil {
		panic(err)
	}

	concatenated := fmt.Sprintf("%s %s, Age: %d has been created with ID:%d", newUser.Name, newUser.Last_name, newUser.Age, newUser.Age)
	c.IndentedJSON(http.StatusCreated, concatenated)
}

// func updateUser(c *gin.Context) {

// }

// func getUserById(c *gin.Context) {
// 	// err = db.QueryRow(sqlStatement, name, last_name, age).Scan(&id)
// }

func getUsers(c *gin.Context) {

	var users = []user{}

	db, err := dbConnection()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name, last_name, age FROM users LIMIT $1", 100)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string
	var last_name string
	var age int

	for rows.Next() {
		err = rows.Scan(&id, &name, &last_name, &age)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, last_name, age)
		user := user{ID: id, Name: name, Last_name: last_name, Age: age}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	// router.GET("/users/:id", getUserById)
	router.Run("localhost:8080")
}
