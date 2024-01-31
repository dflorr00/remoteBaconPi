package tools

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Session struct{}

// Clave de encriptación
var jwtKey = "clave-secreta"

// Genera un jwt para la sesion del usuario indicado
func (s *Session) GeneraToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(jwtKey))
}

// Obtiene el id del usuario de la cookie sino encuentra devuelve 0
func (s *Session) DecodificaIDToken(tokenString string) int {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Método de firma no válido: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil

	})
	if err != nil {
		log.Printf("Error al decodificar el token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if userID, ok := claims["userID"].(float64); ok {
			return int(userID)
		}

	}

	return 0
}
