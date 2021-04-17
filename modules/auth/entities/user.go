package entities

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Sex       int
	Avatar    string
	Money     int
	Token     string
	Status    int
	CreatedAt int
	UpdatedAt int
}

func (User) TableName() string {
	return "user"
}
