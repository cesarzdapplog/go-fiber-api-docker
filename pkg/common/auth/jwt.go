package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "go-fiber-api-docker/pkg/common/models"
)

// Clave secreta para firmar el token JWT
var jwtSecret = []byte("tu_cl4v3_s3cr3t4")

// Función para generar un token JWT
func generateJWTToken(user models.User) (string, error) {
    // Crear claims personalizados
    claims := jwt.MapClaims{
        "id":    user.Id,
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
    }

    // Crear token con claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Firmar el token con la clave secreta
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
