package main
import (
 "fmt"
 "net/http"
 "os"
 "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)
func main() {
 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "3001"
 }
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
 })
 http.ListenAndServe(":" + PORT, nil)
}
