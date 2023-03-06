package main

import (
	"log"

	"github.com/uptrace/bun"
)

func updateSearchElementsTable(db *bun.DB) func() {
	return func() {
		_, err := db.Exec("CALL update_search_elements();")
		if err != nil {
			log.Println("cannot updateSearchTable", err)
		} else {
			log.Println("search table updated")
		}
		_, err = db.Exec("CALL update_lexemes();")
		if err != nil {
			log.Println("cannot updateLexemes", err)
		} else {
			log.Println("search table updated")
		}
	}
}
