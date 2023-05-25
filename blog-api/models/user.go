package models

type User struct {
	ID       uint64 `gorm:"primaryKey:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Posts    []Post `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Posts,omitempty"`
}
