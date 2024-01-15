package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetIP(c *fiber.Ctx) string {
	// Mendapatkan alamat IP asli dari header X-Forwarded-For atau X-Real-IP
	ip := c.Get("X-Forwarded-For")
	if ip == "" {
		ip = c.Get("X-Real-IP")
	}

	// Jika header tidak ditemukan, gunakan IP langsung dari koneksi
	if ip == "" {
		ip = c.IP()
	}

	return ip
}

func NewIP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("IP CLIENT IS : ", GetIP(c))

		return c.Next()
	}
}
