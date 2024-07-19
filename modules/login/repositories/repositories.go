package repositories

import (
	"database/sql"
	"errors"
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

	var MbID, storedPassword string
	var failedAttempts int

	err := r.db.QueryRow("SELECT mb_id FROM member WHERE mb_username = $1", login.MbUsername).Scan(&MbID)
	if err != nil {
		log.Println("failed to find member:", err)
		return err
	}

	
	err = r.db.QueryRow("SELECT mbc_password, failed_attempts FROM member_credential WHERE mbc_mb_id = $1", MbID).Scan(&storedPassword, &failedAttempts)
	if err != nil {
		log.Println("failed to find member_credential:", err)
		return err
	}


	if failedAttempts >= 3 {
		log.Println("account locked due to too many failed attempts")
		return errors.New("account locked due to too many failed attempts. Please contact admin")
	}

	
	if storedPassword != login.MbPassword {
		log.Println("password mismatch")

		_, err := r.db.Exec("UPDATE member_credential SET failed_attempts = failed_attempts + 1 WHERE mbc_mb_id = $1", MbID)
		if err != nil {
			log.Println("failed to update failed attempts:", err)
		}

		return errors.New("invalid credentials")
	}

	_, err = r.db.Exec("UPDATE member_credential SET failed_attempts = 0 WHERE mbc_mb_id = $1", MbID)
	if err != nil {
		log.Println("failed to reset failed attempts:", err)
	}

	return nil
}
