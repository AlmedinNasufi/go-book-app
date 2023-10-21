package models

type Book struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Title    string   `gorm:"type:varchar(100)" json:"title"`
	AuthorID uint     `json:"author_id"`
	Reviews  []Review `gorm:"foreignKey:BookID" json:"reviews"`
}
