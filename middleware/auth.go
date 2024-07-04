package middleware

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler{
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == ""{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"No token provided"})
		}
		token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil,fiber.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")),nil
		})
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON((fiber.Map{"error":"Invalid token"}))
		}
		if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
			c.Locals("userId",claims["id"])
			return c.Next();
		}else{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid Token"})
		}
	}
}

func GenerateJWT(username string) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"exp":time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString,err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err!=nil{
		log.Fatal("Error in generating the key")
		return "",err;
	}
	return tokenString,nil
}