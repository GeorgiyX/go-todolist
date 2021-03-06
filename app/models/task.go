package models

type Task struct {
	Id          uint   `json:"id" gorm:"primaryKey;colum:id;<-:false;autoIncrement:true"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null;default:''"`
	Checked     *bool  `json:"checked" gorm:"not null;default:false"`
}

//easyjson:json
type Tasks []*Task
