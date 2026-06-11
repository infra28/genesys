package handlers

import (
	"backend/pkg/database"
	"backend/pkg/models"
	"backend/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func Login(c fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var data Request
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	result := database.DB.Preload("Role").Where("email = ?", data.Email).First(&user)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	match, _ := utils.ComparePassword(data.Password, user.Password)
	if !match {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	// --- GENERATE JWT TOKEN ---
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})
	t, _ := token.SignedString([]byte(viper.GetString("JWT_SECRET")))

	// 2. Buat Refresh Token (Lama, misal 7 Hari)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	rt, _ := refreshToken.SignedString([]byte(viper.GetString("JWT_SECRET")))

	// 3. Simpan Refresh Token di Cookie (HTTP-Only)
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    rt,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HTTPOnly: true,  // Tidak bisa dibaca oleh JavaScript di Browser
		Secure:   false, // Hanya lewat HTTPS
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"access_token": t,
		"name":         user.Name,
		"role":         user.Role.Name,
	})
}

func Refresh(c fiber.Ctx) error {
	// Ambil token dari cookie
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthenticated"})
	}

	// Validasi token tersebut
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Refresh token expired"})
	}

	// Jika valid, buat Access Token baru
	claims := token.Claims.(jwt.MapClaims)

	// Opsional: Query DB lagi untuk ambil Role terbaru
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})
	t, _ := newAccessToken.SignedString([]byte(viper.GetString("JWT_SECRET")))

	return c.JSON(fiber.Map{"access_token": t})
}

func Register(c fiber.Ctx) error {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var data Request
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input format"})
	}

	hashedPassword, _ := utils.HashPassword(data.Password)

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		RoleID:   2, // Default sebagai 'user'
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

func Logout(c fiber.Ctx) error {
	// Menghapus cookie dengan mengatur masa berlakunya ke waktu lampau
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 jam yang lalu
		HTTPOnly: true,
		Secure:   false, // Pastikan ini sama dengan saat login (false untuk localhost HTTP)
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil logout",
	})
}
