package proto

type Menu struct {
	Id int
	Desc string
	Icon string
	ParentId int
	Link string
	Sort int
}

type ReqUrl struct {
	Id int
	Desc string
	Url string
}

type UserInfo struct {
	UserName string
	NickName string
	Email string
}