	
package main
 
import (
    "database/sql"
	"encoding/json"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "127.0.0.1"
    port     = 5432
    user     = "postadmin"
    password = "postadmin"
    dbname   = "postadmin"
)

type schema struct {
	user_id	string `json:"userid"`
	name	string `json:"name"`
	age		int `json:"age"`
	phone	string `json:"phone"`
}

func main() {
	// connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
         
	// open database
	db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
     
	// close database
    defer db.Close()

	readStmt := "SELECT user_id, COALESCE(name, ''), COALESCE(age, 0), COALESCE(phone, '') FROM user_table"
	rows, e := db.Query(readStmt)
	CheckError(e)
	
	var jsonData []*schema
	for rows.Next() {
		r := new(schema)
		err := rows.Scan(&r.user_id, &r.name, &r.age, &r.phone)
		fmt.Println(r.user_id, r.name, r.age, r.phone)
		CheckError(err)
		jsonData = append(jsonData, r)
	}
	CheckError(err)

	m, merr := json.Marshal(jsonData)
	CheckError(merr)
	fmt.Println(jsonData)
	fmt.Println(string(m))
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}