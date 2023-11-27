package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/assignments", assignments)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}

type Assignments struct{
	Max_score int `json:"max_score"`
	Title string `json:"title"`
	Description string `json:"description"`
	Created_at string `json:"created_at"`
}

func assignments(res http.ResponseWriter, req *http.Request) {
	lesson_id := req.URL.Query()["lesson_id"][0]

	//ссылка пока нерабочая
	connStr := "postgresql://postgres:Pa$$w0rd@localhost:5432/postgres?sslmode=disable"

	db, _ := sql.Open("postgres", connStr)
	defer db.Close()

	var assignmentsList []Assignments

	rows, _ := db.Query(`SELECT max_score, title, description, created_at FROM ASSIGNMENTS WHERE lesson_id = $1`, lesson_id)
	for rows.Next() {
		var rowSlice Assignments
		err := rows.Scan(&rowSlice.Max_score, &rowSlice.Title, &rowSlice.Description, &rowSlice.Created_at)
		if err != nil{
			fmt.Println(err)
		}
		assignmentsList = append(assignmentsList, rowSlice)
	}

	result, _ := json.Marshal(assignmentsList)
	res.Write(result)
}
