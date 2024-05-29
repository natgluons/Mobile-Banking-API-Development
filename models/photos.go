package models

type Photos struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(300)" json:"title"`
	Caption  string `gorm:"type:varchar(300)" json:"caption"`
	Photourl string `gorm:"type:varchar(300)" json:"photourl"`
}
