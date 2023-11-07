package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		log.Println("commitrollback", err)
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
