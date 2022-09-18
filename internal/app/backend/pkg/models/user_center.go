package models

type UserCenter struct {
	BasicModel
	Account  string `gorm:"type:varchar(300);index" json:"account"`
	Password string `gorm:"type:varchar(600)" json:"password"`
	Sale     string `gorm:"type:varchar(600)" json:"sale"`
}
