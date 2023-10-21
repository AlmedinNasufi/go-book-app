package models

type Review struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Text   string `gorm:"type:text" json:"text"`
	BookID uint   `json:"book_id"`
}
