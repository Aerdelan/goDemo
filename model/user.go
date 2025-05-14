package model

// 定义返回数据结构体

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Area struct {
	ID       int    `json:"id"`
	Areaname string `json:"Areaname"`
	Prientid int    `json:"Prientid"`
	Children []Area `json:"Children"`
}
