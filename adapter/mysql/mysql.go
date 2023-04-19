package mysql

import (
	"database/sql"
	"log"
	"lykafe/news/config"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = connect()
	if err != nil {
		log.Fatal("[DEBUG] Cannot connect to mysql", err)
	}

	migrate(db)
}

func connect() (*sql.DB, error) {
	var sb strings.Builder
	sb.WriteString(config.EnvMysqlUser())
	sb.WriteString(":")
	sb.WriteString(config.EnvMysqlPassword())
	sb.WriteString("@tcp(")
	sb.WriteString(config.EnvMysqlAddr())
	sb.WriteString(")/")
	sb.WriteString(config.EnvMysqlDB())
	// sb.WriteString("?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC") //charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=UTC
	sb.WriteString("?charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=UTC")
	dsn := sb.String()
	log.Println("[DEBUG] dsn: ", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db, err
}

func migrate(db *sql.DB) {
	// check column `content_jp` exists, if not, add it
	if _, err := db.Exec(`ALTER TABLE news ADD COLUMN content_jp TEXT NULL AFTER content`); err != nil {
		// ignore error
	}

	// check column `description_js` exists, if not, add it
	if _, err := db.Exec(`ALTER TABLE news ADD COLUMN description_jp TEXT NULL AFTER description`); err != nil {
		// ignore error
	}
}
