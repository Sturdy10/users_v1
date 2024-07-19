package repositories

import (
	"database/sql"
	"errors"
	"log"
	"users_v1/modules/register/models"
)

type IRepositorie interface {
	RegisterMemberRepository(employee models.RegisterMember) error
	CheckExistingEmail(email string) (bool, error)
	GetallMembersRepository() ([]models.MemberResponse, error)
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repository{db: db}
}

func (r *repository) CheckExistingEmail(email string) (bool, error) {
	var existingEmail string
	err := r.db.QueryRow("SELECT orgmb_email FROM organize_member WHERE orgmb_email = $1", email).Scan(&existingEmail)
	if err == nil {
		return true, nil // หากพบว่าอีเมลซ้ำ
	} else if err != sql.ErrNoRows {
		return false, err // กรณีเกิด error อื่นที่ไม่ใช่ sql.ErrNoRows
	}
	return false, nil // หากไม่พบอีเมลซ้ำ
}

func (r *repository) RegisterMemberRepository(addMember models.RegisterMember) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println("failed to begin transaction:", err)
		return err
	}
	defer func() {
		if err != nil {
			log.Println("rolling back transaction due to error:", err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Println("failed to commit transaction:", err)
			tx.Rollback()
		}
	}()

	var orgdpID string
	err = tx.QueryRow("INSERT INTO organize_department(orgdp_name) VALUES ($1) RETURNING orgdp_id", addMember.OrgdpName).Scan(&orgdpID)
	if err != nil {
		log.Println("failed to insert orgdp_department:", err)
		return errors.New("failed to insert department")
	}

	var orgrlID string
	err = tx.QueryRow("INSERT INTO organize_role(orgrl_orgdp_id) VALUES ($1) RETURNING orgrl_id", orgdpID).Scan(&orgrlID)
	if err != nil {
		log.Println("failed to insert orgrl_id:", err)
		return err
	}

	var orgmbID string
	err = tx.QueryRow("INSERT INTO organize_member(orgmb_title, orgmb_name, orgmb_surname, orgmb_email, orgmb_mobile, orgmb_role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING orgmb_id ", addMember.OrgmbTitle, addMember.OrgmbName, addMember.OrgmbSurname, addMember.OrgmbEmail, addMember.OrgmbMobile, orgrlID).Scan(&orgmbID)
	if err != nil {
		log.Println("failed to insert organize_member:", err)
		return err
	}

	// เพิ่มข้อมูลรหัสผ่าน
	_, err = tx.Exec("INSERT INTO organize_member_credential(orgmbcr_orgmb_id, orgmbcr_password) VALUES ($1, $2)", orgmbID, addMember.GeneratedPassword)
	if err != nil {
		log.Println("failed to insert organize_member_credential:", err)
		return err
	}

	return nil
}

func (r *repository) GetallMembersRepository() ([]models.MemberResponse, error) {
	var members []models.MemberResponse

	rows, err := r.db.Query(`
        SELECT
		    m.orgmb_id,
            m.orgmb_title,
            m.orgmb_name,
            m.orgmb_surname,
            m.orgmb_email,
            m.orgmb_mobile,
            r.orgrl_orgdp_id AS role,
            d.orgdp_name AS department
        FROM
            organize_member m
        JOIN
            organize_role r ON m.orgmb_role = r.orgrl_id
        JOIN
            organize_department d ON r.orgrl_orgdp_id = d.orgdp_id
    `)
	if err != nil {
		log.Println("failed to query organize_member:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member models.MemberResponse
		err := rows.Scan(&member.OrgmbID, &member.OrgmbTitle, &member.OrgmbName, &member.OrgmbSurname, &member.OrgmbEmail, &member.OrgmbMobile, &member.OrgrlOrgdpID, &member.OrgdpName)
		if err != nil {
			log.Println("failed to scan organize_member:", err)
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
