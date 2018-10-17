package defs

type UserCredential struct { // 自动转 json
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comments struct {
	Id string
	VideoId string
	Author string
	Content string
}