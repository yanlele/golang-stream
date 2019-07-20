package defs

// request
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}


// Comments model
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}
