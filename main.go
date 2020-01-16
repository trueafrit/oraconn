package main
 
import (
    "fmt"
    "database/sql"
    _ "github.com/godror/godror"
)
 
func main(){
 
    db, err := sql.Open("godror", "ani/ani5@192.168.12.215:1521/aspwdm")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
     
     
    rows,err := db.Query("select count(*) from ani.ani_airports_t")
    if err != nil {
        fmt.Println("Error running query")
        fmt.Println(err)
        return
    }
    defer rows.Close()
 
    var thedate string
    for rows.Next() {
 
        rows.Scan(&thedate)
    }
    fmt.Printf("airports count is: %s\n", thedate)
}
