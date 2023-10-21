package models

type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(100)" json:"name"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}
