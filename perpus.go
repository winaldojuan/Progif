package main
import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

type myBuku struct {
    id_buku int
    judul string
    penulis string
    tahun_terbit int
    penerbit string
}

func main(){
    port:=8181
    http.HandleFunc("/buku/", func(w http.ResponseWriter, r*http.Request){
        switch r.Method{
        case "GET":
             GetAllBuku(w,r)
        case "POST":
        
        case "DELETE":
        default:
            http.Error(w, "Invalid request method.", 405)
        }
    })
    log.Printf("Server starting on port %v\n",port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func GetAllBuku(w http.ResponseWriter, r *http.Request){
    db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
    if err!=nil{
        log.Fatal(err)
    }
    defer db.Close()

    buku:=myBuku{}

    rows,err :=db.Query("select id_buku, judul, penulis, tahun_terbit, penerbit from buku")
    if err!=nil{
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next(){
        err := rows.Scan(&buku.id_buku, &buku.judul, &buku.penulis, &buku.tahun_terbit, &buku.penerbit)
        if err!=nil{
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&buku)
    }
	err = rows.Err()
	
}
