package model

import (
	"fmt"
	"log"
	
	_ "github.com/joho/godotenv/autoload"
	"github.com/joho/godotenv"
)

func GetRoutine() {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("SELECT date FROM routine").Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func PostRoutine(routine, memo string) {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("INSERT INTO routine(date, todo, memo, is_achieved)) VALUES (now(), ?, ?,0)", routine, memo).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func PutRoutine(routine, memo string, is_achieved, id int) {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("UPDATE routine SET todo=?, memo=?, is_achieved=? WHERE id=?", routine, memo, is_achieved, id).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func DeleteRoutine(id int) {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("DELETE FROM routine WHERE id=?", id).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func GetUser() {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("SELECT * FROM user").Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func PostUser(name string) {
	db, err := ConnectDB()	

	var cal [6]string
	err = db.QueryRow("INSERT INTO user(name) VALUES (?)", name).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func PutUser(name string, id int) {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("UPDATE user SET name=? WHERE id=?", name, id).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func DeleteUser(id int) {
	db, err := ConnectDB()

	var cal [6]string
	err = db.QueryRow("DELETE FROM user WHERE id=?", id).Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}