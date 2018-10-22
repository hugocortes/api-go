package sql

// UserService handles communication with the User model
type UserService service

const (
	userTable = "user"
)

// User model
type User struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

// GetUserByID returns single user
func (conn *DB) GetUserByID(ID string) User {
	var user User
	conn.db.Table(userTable).Select("*").Where("id = ?", ID).Find(&user)

	return user
}
