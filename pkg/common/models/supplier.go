package models

type Supplier struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Address string `json:"address"`
}
