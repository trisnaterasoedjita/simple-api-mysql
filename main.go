package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type students struct {
	id    int
	name  string
	age   int
	class string
}

type reqStudents struct {
	name  string
	age   int
	class string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:s!mp3lAp1@tcp(materi-simple-golang-db:3309)/golang_simple_api_mysql")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func queryAmbilData() ([]students, error) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from students")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []students

	for rows.Next() {
		var each = students{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.class)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil

}

func queryAddData(param reqStudents) error {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO students(name,age,class) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // close after used
	_, err = stmt.Exec(param.name, param.age, param.class)
	if err != nil {
		return err
	}
	return nil
}

func studentList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		data, error := queryAmbilData()

		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		mapData := make([]map[string]interface{}, 0, 0)
		for _, value := range data {
			temp := make(map[string]interface{})
			temp["id"] = value.id
			temp["name"] = value.name
			temp["age"] = value.age
			temp["class"] = value.class
			mapData = append(mapData, temp)
		}
		var result, err = json.Marshal(mapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func studentAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var param reqStudents
	param.name = r.FormValue("name")
	param.age, _ = strconv.Atoi(r.FormValue("age")) //convert string to integer
	param.class = r.FormValue("class")
	if r.Method == "POST" {
		err := queryAddData(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp := make(map[string]interface{})
		temp["msesage"] = "Data berhasil disimpan"

		var result, error = json.Marshal(temp)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/students", studentList)
	http.HandleFunc("/student", studentAdd)
	fmt.Println("starting web server at http://localhost:4023/")
	http.ListenAndServe(":4023", nil)
}
