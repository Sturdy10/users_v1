package generatepw

import (
	"math/rand"
	"time"
)

// นี่คือฟังก์ชันสำหรับสร้างรหัสผ่านแบบสุ่ม
func GenerateRandomPassword(lenNumber int) (string, error) {
	// ชุดอักขระที่ใช้สำหรับสร้างรหัสผ่าน
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@$&."
	// กำหนด random number generator ด้วย seed จากเวลาปัจจุบัน
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// สร้าง slice เพื่อเก็บรหัสผ่าน
	password := make([]byte, lenNumber)
	// วนลูปเพื่อสร้างแต่ละตัวอักษรของรหัสผ่าน
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	// คืนค่ารหัสผ่านและ nil เป็นผลลัพธ์
	return string(password), nil
}
