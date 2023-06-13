package entities

type User struct {
	ID       uint   `gorm:"primaryKey;"`
	Username string `gorm:"size:255;unique;not null;"`
	Name     string `gorm:"size:255;not null;"`
}
