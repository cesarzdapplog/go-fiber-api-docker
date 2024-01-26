package models
import (
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
)
type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
    Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
    Password string `json:"-" gorm:"not null" validate:"required"`
}
var validate = validator.New()
func (u *User) SetPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword) // Convertir el hash a string si es necesario
    return nil
}

// CheckPassword verifica si la contraseña proporcionada coincide con la contraseña encriptada en la base de datos
func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Validate() error {
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}