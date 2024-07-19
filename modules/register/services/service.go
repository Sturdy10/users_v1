package services

import (
	"errors"
	"log"

	"users_v1/modules/register/models"
	"users_v1/modules/register/repositories"
	"users_v1/pkg/utility/generatepw"
	"users_v1/pkg/utility/mail"

	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	RegisterMemberService(addMember models.RegisterMember) error
	GetallMemberService() ([]models.MemberResponse, error)
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) RegisterMemberService(addMember models.RegisterMember) error {
	// ตรวจสอบว่าอีเมลซ้ำหรือไม่
	existingEmail, err := s.r.CheckExistingEmail(addMember.OrgmbEmail)
	if err != nil {
		log.Println("failed to check existing email:", err)
		return err
	}
	if existingEmail {
		return errors.New("email already exists")
	}

	// สร้างรหัสผ่านสุ่ม
	generatedPassword, err := generatepw.GenerateRandomPassword(12)
	if err != nil {
		log.Println("failed to generate random password:", err)
		return err
	}

	// เข้ารหัสรหัสผ่านที่สร้างขึ้น
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password. Error: %v\n", err)
		return err
	}

	// แปลงรหัสผ่านที่เข้ารหัสแล้วเป็น string
	hashedPasswordString := string(hashedPassword)

	// กำหนดรหัสผ่านที่เข้ารหัสแล้วให้กับ request register
	addMember.GeneratedPassword = hashedPasswordString

	// ส่งอีเมล
	err = mail.MailPassword(addMember.OrgmbEmail, generatedPassword)
	if err != nil {
		log.Println("failed to send email:", err)
		return err
	}

	// เรียกใช้ repository ในการเพิ่มข้อมูลสมาชิก
	err = s.r.RegisterMemberRepository(addMember)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) GetallMemberService() ([]models.MemberResponse, error) {
	members, err := s.r.GetallMembersRepository()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return members, nil

}
