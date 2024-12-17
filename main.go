package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getEnvOrDefault(key string, defaultValue string) string {
	var (
		v     string
		exist bool
	)
	if v, exist = os.LookupEnv(key); !exist {
		return defaultValue
	}
	return v
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(`
		<h1> Salut la classe! </h1>	
  `))
}
