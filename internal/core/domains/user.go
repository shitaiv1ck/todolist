package domains

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Username          string
	Password          string
	EncryptedPassword string
}

func NewUser(id int, username string, password string, encryptedPassword string) *User {
	return &User{
		ID:                id,
		Username:          username,
		Password:          password,
		EncryptedPassword: encryptedPassword,
	}
}

func NewUninitializedUser(username string, password string) *User {
	return &User{
		ID:                -1,
		Username:          username,
		Password:          password,
		EncryptedPassword: "",
	}
}

func (u *User) EncryptePassword() error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(encryptedPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}
