package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "net/http/pprof"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"_"`
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("users", listUserHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	// mux.HandleFunc("/cpu", CPUIntensiveEndpoint)
	go http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":6060", nil)
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decore(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := db.Exec(
		"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
		u.ID, u.Name, u.Email,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// func CPUIntensiveEndpoint(w http.ResponseWriter, r *http.Request) {
// 	result := fibonacci(60)
// 	w.Write([]byte(strconv.Itoa(result)))
// }

// func generateLargeString(n int) string{
// 	var buffer bytes.Buffer
// 	for i := 0; i<n; i++{
// 		for j :=0; j<n;j++{
// 			buffer.WriteString(strconv.Itoa(i + j*j))
// 		}
// 	}
// 	return buffer.String()
// }
