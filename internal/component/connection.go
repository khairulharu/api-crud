package component

import (
	"database/sql"
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql"

	"github.com/khairulharu/miniapps/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable ",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error when connect to database:%s", err.Error())
	}
	log.Println("connectinon to database succes")
	return db
}
