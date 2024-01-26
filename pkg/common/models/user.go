package models
import "golang.org/x/crypto/bcrypt"
type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
    Email    string `json:"email" gorm:"unique;not null"`
    Password string `json:"-" gorm:"not null"`
}

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