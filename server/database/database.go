package database

import (
	"database/sql"
	"emp-parser/server/globals"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init ..
func Init() {
	var err error
	db, err = sql.Open("mysql", "root:gogoedt@tcp(192.168.99.100:3306)/edt")
	if err != nil {
		globals.ErrLogger.Println("Erreur à la connexion SQL")
		panic(err.Error())
	}
}

// Record ..
func Record(list []globals.Creneau) {
	fmt.Println("Enregistement et base de la liste des créneaux...")
	for _, c := range list {
		recordCreneau(c)
	}
}

func recordCreneau(c globals.Creneau) {
	stmtIns, err := db.Prepare("INSERT INTO creneaux VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(c.Summary, c.Location, c.Description)
	if err != nil {
		panic(err.Error())
	}

}
