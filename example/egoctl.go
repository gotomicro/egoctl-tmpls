// +build egoctl
package main

type User struct {
	Uid      int    `gorm:"AUTO_INCREMENT" json:"id" binding:"required" ego:"primary_key"` // id
	UserName string `gorm:"not null" json:"userName" binding:"required"`                   // 昵称
}
