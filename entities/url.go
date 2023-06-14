package entities

type Url struct {
	ID       uint   `gorm:"primaryKey;"`
	Url      string `gorm:"size:255; not null"`
	Redirect string `gorm:"size:255; not null"`
	UserID   uint   `gorm:"not null"`
	User     User
}
