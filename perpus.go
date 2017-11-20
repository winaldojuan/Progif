package main
import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    
    "net/http"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)

type myBuku struct {
    ID_buku int
    Judul string
    Penulis string
    Tahun_terbit int
    Penerbit string
}
type myBuku2 struct{
	ID_buku int `json:"ID_buku, omitempty"`
	Judul string `json:"Judul, omitempty"`
	Penulis string `json:"Penulis, omitempty"`
	Tahun_terbit int `json:"Tahun_terbit, omitempty"`
	Penerbit string `json:"Penerbit, omitempty"`
}



func main(){
    port:=8181
    http.HandleFunc("/test", func(w http.ResponseWriter, r*http.Request){
        http.ServeFile(w,r,"testing.html")
    })
    http.HandleFunc("/update/", func (w http.ResponseWriter, r *http.Request){
        http.ServeFile(w,r,"up.html")
    })

  http.HandleFunc("/put/", func(w http.ResponseWriter, r *http.Request){
     
                http.ServeFile(w,r,"update.html")
       
    })
  
    http.HandleFunc("/buku/", func(w http.ResponseWriter, r*http.Request){
        switch r.Method{
		case "GET":
			if r.URL.Query().Get("judul") != ""{
				judul := r.URL.Query().Get("judul")
				SearchByJudul(w,r,judul)
			}else if r.URL.Query().Get("penerbit") != ""{
				penerbit := r.URL.Query().Get("penerbit")
				SearchByPenerbit(w,r,penerbit)
			}else if r.URL.Query().Get("tahun") != ""{
				tahun := r.URL.Query().Get("tahun")
				SearchByTahun(w,r,tahun)
			}else{
				GetAllBuku(w,r)
			}
            
        case "POST":
            InsertBuku(w,r)
		case "DELETE":
			s := r.URL.Path[len("/buku/"):]
			DeleteBuku(w,r,s)
		case "PUT":
			s := r.URL.Path[len("/buku/"):]
			UpdateBuku(w,r,s)
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

    rows,err :=db.Query("select ID_buku, Judul, Penulis, Tahun_terbit, Penerbit from buku")

    if err!=nil{
        log.Fatal(err)
    }
    defer rows.Close()

    
    for rows.Next(){

        err := rows.Scan(&buku.ID_buku, &buku.Judul, &buku.Penulis, &buku.Tahun_terbit, &buku.Penerbit)
        //fmt.printf("%v", buku.id_buku)
        if err!=nil{
            log.Fatal(err)
        }
	
        json.NewEncoder(w).Encode(&buku)
    }
	err = rows.Err()
	
}

func SearchByJudul(w http.ResponseWriter, r *http.Request, s string){
	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
    if err!=nil{
        log.Fatal(err)
    }
    defer db.Close()

    buku:=myBuku{}

    rows,err :=db.Query("select ID_buku, Judul, Penulis, Tahun_terbit, Penerbit from buku where Judul like ?", "%"+s+"%")

    if err!=nil{
        log.Fatal(err)
    }
    defer rows.Close()

    
    for rows.Next(){

        err := rows.Scan(&buku.ID_buku, &buku.Judul, &buku.Penulis, &buku.Tahun_terbit, &buku.Penerbit)
        //fmt.printf("%v", buku.id_buku)
        if err!=nil{
            log.Fatal(err)
        }
	
        json.NewEncoder(w).Encode(&buku)
    }
	err = rows.Err()
}

func SearchByPenerbit(w http.ResponseWriter, r *http.Request, s string){
	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
    if err!=nil{
        log.Fatal(err)
    }
    defer db.Close()

    buku:=myBuku{}

    rows,err :=db.Query("select ID_buku, Judul, Penulis, Tahun_terbit, Penerbit from buku where Penerbit like ?", "%"+s+"%")

    if err!=nil{
        log.Fatal(err)
    }
    defer rows.Close()

    
    for rows.Next(){

        err := rows.Scan(&buku.ID_buku, &buku.Judul, &buku.Penulis, &buku.Tahun_terbit, &buku.Penerbit)
        //fmt.printf("%v", buku.id_buku)
        if err!=nil{
            log.Fatal(err)
        }
	
        json.NewEncoder(w).Encode(&buku)
    }
	err = rows.Err()
}

func SearchByTahun(w http.ResponseWriter, r *http.Request, s string){
	sint, _ :=strconv.Atoi(s)
	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
    if err!=nil{
        log.Fatal(err)
    }
    defer db.Close()

    buku:=myBuku{}

    rows,err :=db.Query("select ID_buku, Judul, Penulis, Tahun_terbit, Penerbit from buku where tahun_terbit =?", sint)

    if err!=nil{
        log.Fatal(err)
    }
    defer rows.Close()

    
    for rows.Next(){

        err := rows.Scan(&buku.ID_buku, &buku.Judul, &buku.Penulis, &buku.Tahun_terbit, &buku.Penerbit)
        //fmt.printf("%v", buku.id_buku)
        if err!=nil{
            log.Fatal(err)
        }
	
        json.NewEncoder(w).Encode(&buku)
    }
	err = rows.Err()
}

func InsertBuku (w http.ResponseWriter, r *http.Request) {
	var buku myBuku2
	dec := json.NewDecoder(r.Body)
	err:=dec.Decode(&buku)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()

	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
	if err != nil{
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO test (ID_buku,Judul, Penulis,Tahun_terbit, Penerbit) VALUES (?,?,?,?,?)")
	if err != nil{
		log.Fatal(err)
	}
	_, err = stmt.Exec(buku.ID_buku, buku.Judul, buku.Penulis, buku.Tahun_terbit, buku.Penerbit)
}

func DeleteBuku(w http.ResponseWriter, r *http.Request, id string){
	idbuku,_ := strconv.Atoi(id)

	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
	if err != nil{
		log.Fatal(err)
	}

	rows,err := db.Query("DELETE FROM test WHERE ID_buku=?",idbuku)

	defer rows.Close()

}

func UpdateBuku (w http.ResponseWriter, r *http.Request, id string) {
	var buku myBuku2
	idbuku,_ := strconv.Atoi(id)
	dec := json.NewDecoder(r.Body)
	err:=dec.Decode(&buku)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()

	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
	if err != nil{
		log.Fatal(err)
	}

	stmt, err := db.Prepare("UPDATE test SET ID_buku = ?,Judul = ?, Penulis = ?,Tahun_terbit =?, Penerbit=? where ID_buku=?")
	if err != nil{
		log.Fatal(err)
	}
	_, err = stmt.Exec(buku.ID_buku, buku.Judul, buku.Penulis, buku.Tahun_terbit, buku.Penerbit, idbuku)
}