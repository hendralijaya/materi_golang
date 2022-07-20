package domain

type Book struct {
	Id     uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Title  string `json:"title" gorm:"type:varchar(255);not null"`
	Author string `json:"author" gorm:"type:varchar(255);not null"`
}