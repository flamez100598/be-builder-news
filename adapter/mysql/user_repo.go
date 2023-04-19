package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type UserRepo struct {
	db *sql.DB
	upsertUser *sql.Stmt
}

func NewUserRepo() *UserRepo {
	upsertUser, err := db.Prepare(`INSERT INTO user(id, username, email, name, avatar) VALUES(UNHEX(?), ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE username = ?, email = ?, name = ?, avatar = ?`)
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare upsertUser ", err)
	}
	return &UserRepo {
		db: db,
		upsertUser: upsertUser,
	}
}

func (m *UserRepo) UpdateUser(id, username, email, name, avatar string) error {
	_, err := m.upsertUser.Exec(id, username, email, name, avatar, username, email, name, avatar)
	if err != nil {
		log.Println("[DEBUG] UpdateUser err: ", err)
		return err
	}
	// TODO: check res.LastInsertId()
	return err
}