package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//CREDENTIALS
const db_user = "BD"
const db_passwd = "12341234"
const db_addr = "localhost"
const db_db = "dai"

//STRUCT FACT
type fact struct {
	Id_fact     int
	Name        string
	Description string
	Photo       string
}

//STRUCT BIRD
type Bird struct {
	Name   string `json: "name" `
	Numero int    `json: "numero"`
}

//FUNCTION FOR TESTS
func Something(w http.ResponseWriter, r *http.Request) {
	//REST METHODS
	switch r.Method {
	//CASO SEJA UM GET
	case "GET":
		//FACTOS
		var factos []fact
		//SETAR O TIPO DE DADOS
		w.Header().Add("content-type", "application/json")
		//STATUS OK
		w.WriteHeader(http.StatusOK)
		//ESTABLECER A CONEXAO
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db))
		if err != nil {
			log.Fatal(err)
		}
		//FAZER SELECT
		resp, err := db.Query("SELECT * FROM `fact`")
		if err != nil {
			log.Fatal(err)
		}
		//GUARDAR DATA
		for resp.Next() {
			var rfact fact
			err = resp.Scan(&rfact.Id_fact, &rfact.Name, &rfact.Description, &rfact.Photo)
			if err != nil {
				log.Fatal(err)
			}
			//por na slice
			factos = append(factos, rfact)
		}
		//QUANDO A FUNCAO ACABAR ISTO FECHA
		defer db.Close()
		//TRANSFORMAR EM JSON
		resposta, _ := json.Marshal(factos)
		//MSG
		w.Write(resposta)
		return
	//CASO SEJA UM POST
	case "POST":
		//facts
		var factos []fact
		//LER O BODY
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			//CASO DE ERRO DAR COMO ERRO INTERNO
			w.WriteHeader(http.StatusInternalServerError)
			//ESCREVER ISTO
			w.Write([]byte(err.Error()))
		}

		//COLOCAR OS ITEMS DO BODY NO STRUCT
		err = json.Unmarshal(bodyBytes, &factos)
		//CASO DE ERRO
		if err != nil {
			//BAD REQUEST
			w.WriteHeader(http.StatusBadRequest)
			//ESCREVER
			w.Write([]byte(err.Error()))
		}
		//ESTABLECER A CONEXAO
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db))
		if err != nil {
			log.Fatal(err)
		}
		for _, factudos := range factos {
			q := "INSERT INTO `fact` (name,description,photo) VALUES(?,?,?)"
			insert, err := db.Prepare(q)
			if err != nil {
				log.Fatal(err)
			}
			defer insert.Close()
			_, err = insert.Exec(factudos.Name, factudos.Description, factudos.Photo)
			if err != nil {
				log.Fatal(err)
			} else {
				w.WriteHeader(http.StatusOK)
			}

		}
		return
	//CASO SEJA UM PUT
	case "PUT":
		var factos []fact
		//LER O BODY
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			//CASO DE ERRO DAR COMO ERRO INTERNO
			w.WriteHeader(http.StatusInternalServerError)
			//ESCREVER ISTO
			w.Write([]byte(err.Error()))
		}

		//COLOCAR OS ITEMS DO BODY NO STRUCT
		err = json.Unmarshal(bodyBytes, &factos)
		//CASO DE ERRO
		if err != nil {
			//BAD REQUEST
			w.WriteHeader(http.StatusBadRequest)
			//ESCREVER
			w.Write([]byte(err.Error()))
		}
		//ESTABLECER A CONEXAO
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db))
		if err != nil {
			log.Fatal(err)
		}

		for _, factudos := range factos {
			q := "UPDATE `fact` SET `photo`=? WHERE `id_fact`=?"
			update, err := db.Prepare(q)
			if err != nil {
				log.Fatal(err)
			}
			defer update.Close()
			_, err = update.Exec(factudos.Photo, factudos.Id_fact)
			if err != nil {
				log.Fatal(err)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Sucess"))
			}
		}
		return
	//CASO SEJA UM DELETE
	case "DELETE":
		return
	//CASO A ROTA NAO TENHA UM DOS METODOS A CIMA
	default:
		//STATUS NAO ALLOWED
		w.WriteHeader(http.StatusMethodNotAllowed)
		//ESCREVER ALGO
		w.Write([]byte("ERRO"))
		return

	}

}
