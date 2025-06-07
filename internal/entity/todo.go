package entity

type Todo struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Completed bool
}
