package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"users_v1/modules/login/models"
)

type IRepositorie interface {
	LoginR(login models.LoginRequest) error
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repository{db: db}
}

func (r *repository) LoginR(login models.LoginRequest) error {
	var (
		mbID           string
		storedPassword string
		failedAttempts int
	)

	err := r.db.QueryRow("SELECT mb_id FROM member WHERE mb_username = $1", login.MbUsername).Scan(&mbID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("username or password does not match")
		}
		log.Printf("Error querying member: %v", err)
		return fmt.Errorf("db error")
	}

	err = r.db.QueryRow("SELECT mbc_password, failed_attempts FROM member_credential WHERE mbc_mb_id = $1", mbID).Scan(&storedPassword, &failedAttempts)
	if err != nil {
		log.Printf("Error querying credentials: %v", err)
		return fmt.Errorf("db error")
	}

	if failedAttempts >= 3 {
		log.Println("Account locked")
		return fmt.Errorf("account locked, contact admin")
	}

	if storedPassword != login.MbPassword {
		log.Println("Password mismatch")
		failedAttempts++
		_, err = r.db.Exec("UPDATE member_credential SET failed_attempts = $1 WHERE mbc_mb_id = $2", failedAttempts, mbID)
		if err != nil {
			log.Printf("Error updating attempts: %v", err)
		}

		if failedAttempts >= 3 {
			return fmt.Errorf("account locked, contact admin")
		}

		return fmt.Errorf("invalid credentials, %d attempts left", 3-failedAttempts)
	}

	_, err = r.db.Exec("UPDATE member_credential SET failed_attempts = 0 WHERE mbc_mb_id = $1", mbID)
	if err != nil {
		log.Printf("Error resetting attempts: %v", err)
	}

	return nil
}
