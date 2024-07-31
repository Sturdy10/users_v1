package repositories

import (
	"database/sql"
	"errors"
	"log"
	"users_v1/modules/password/models"

	"golang.org/x/crypto/bcrypt"
)

type IRepositorie interface {
	InitPasswordS(initPassword models.InitPassword) error
	ChangePasswordRepository(changePassword models.ChangePassword) error
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repository{db: db}
}

func (r *repository) InitPasswordS(initPassword models.InitPassword) error {
	// เข้ารหัสรหัสผ่านใหม่
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(initPassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v \n", err)
		return err
	}

	// ค้นหา mbID จาก mb_email
	var mbID string
	err = r.db.QueryRow("SELECT mb_id FROM member WHERE mb_email = $1", initPassword.MdUsername).Scan(&mbID)
	if err != nil {
		log.Println("failed to find member:", err)
		return err
	}

	// อัปเดตรหัสผ่านใน member_credential
	_, err = r.db.Exec("UPDATE member_credential SET mbcr_password = $1 WHERE mbcr_mb_id = $2", hashedNewPassword, mbID)
	if err != nil {
		log.Println("failed to update member_credential:", err)
		return err
	}

	return nil
}

func (r *repository) ChangePasswordRepository(changePassword models.ChangePassword) error {
	var mbID string
	err := r.db.QueryRow("SELECT mb_id FROM member WHERE mb_email = $1", changePassword.MdUsername).Scan(&mbID)
	if err != nil {
		log.Println("failed to find member:", err)
		return errors.New("failed to find email address")
	}

	var mbcrPassword string
	err = r.db.QueryRow("SELECT mbcr_password FROM member_credential WHERE mbcr_mb_id = $1", mbID).Scan(&mbcrPassword)
	if err != nil {
		log.Println("failed to update member_credential:", err)
		return err
	}

	// เปรียบเทียบรหัสผ่านเก่า
	err = bcrypt.CompareHashAndPassword([]byte(mbcrPassword), []byte(changePassword.Oldpassword))
	if err != nil {
		log.Println("the password is incorrect:", err)
		return errors.New("the password is incorrect")
	}

	// เข้ารหัสรหัสผ่านใหม่
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.Newpassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to hash new password:", err)
		return err
	}

	// อัปเดตรหัสผ่านใน member_credential
	_, err = r.db.Exec("UPDATE member_credential SET mbcr_password = $1 WHERE mbcr_mb_id = $2", hashedNewPassword, mbID)
	if err != nil {
		log.Println("failed to update member_credential:", err)
		return err
	}

	return nil
}
