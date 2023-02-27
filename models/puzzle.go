package models

type Puzzle struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	FEN     string `gorm:"not null"`
	Moves   string `gorm:"not null"`
	Rating  int    `gorm:"not null"`
	NPieces uint8  `gorm:"not null"`
}
