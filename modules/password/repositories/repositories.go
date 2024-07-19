package repositories

import (
	"database/sql"
	"errors"
	"log"
	"users_v1/modules/password/models"

	"golang.org/x/crypto/bcrypt"
)

type IRepositorie interface {
	InitPasswordRepository(InitPassword models.InitPassword) error
	ChangePasswordRepository(changePassword models.ChangePassword) error
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repository{db: db}
}

func (r *repository) InitPasswordRepository(initPassword models.InitPassword) error {
	// เข้ารหัสรหัสผ่านใหม่
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(initPassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v \n", err)
		return err
	}

	// ค้นหา orgmbID จาก orgmb_email
	var orgmbID string
	err = r.db.QueryRow("SELECT orgmb_id FROM organize_member WHERE orgmb_email = $1", initPassword.OrgmbEmail).Scan(&orgmbID)
	if err != nil {
		log.Println("failed to find organize_member:", err)
		return err
	}

	// อัปเดตรหัสผ่านใน organize_member_credential
	_, err = r.db.Exec("UPDATE organize_member_credential SET orgmbcr_password = $1 WHERE orgmbcr_orgmb_id = $2", hashedNewPassword, orgmbID)
	if err != nil {
		log.Println("failed to update organize_member_credential:", err)
		return err
	}

	return nil
}

func (r *repository) ChangePasswordRepository(changePassword models.ChangePassword) error {
	var orgmbID string
	err := r.db.QueryRow("SELECT orgmb_id FROM organize_member WHERE orgmb_email = $1", changePassword.OrgmbEmail).Scan(&orgmbID)
	if err != nil {
		log.Println("failed to find organize_member:", err)
		return errors.New("failed to find email address")
	}

	var orgmbcrPassword string
	err = r.db.QueryRow("SELECT orgmbcr_password FROM organize_member_credential WHERE orgmbcr_orgmb_id = $1", orgmbID).Scan(&orgmbcrPassword)
	if err != nil {
		log.Println("failed to update organize_member_credential:", err)
		return err
	}

	// เปรียบเทียบรหัสผ่านเก่า
	err = bcrypt.CompareHashAndPassword([]byte(orgmbcrPassword), []byte(changePassword.Oldpassword))
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

	// อัปเดตรหัสผ่านใน organize_member_credential
	_, err = r.db.Exec("UPDATE organize_member_credential SET orgmbcr_password = $1 WHERE orgmbcr_orgmb_id = $2", hashedNewPassword, orgmbID)
	if err != nil {
		log.Println("failed to update organize_member_credential:", err)
		return err
	}

	return nil
}
