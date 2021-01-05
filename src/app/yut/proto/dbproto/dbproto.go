package dbproto

type DBUserInfo struct {
	UserName string `db:"username"`
	Password string `db:"password"`
	Email string `db:"email"`
	GroupId int `db:"group_id"`
	NickName string `db:"nickname"`
	UserDesc string `db:"userdesc"`
	WXImage string `db:"wx_rcode_img"`
	ZFImage string `db:"zf_rcode_img"`
	Sex int `db:"sex"`
}

type DBUserProfile struct {
	NickName string `db:"nickname"`
	UserDesc string `db:"userdesc"`
	WXImage string `db:"wx_rcode_img"`
	ZFImage string `db:"zf_rcode_img"`
	Sex int `db:"sex"`
}

type DBGroupInfo struct {
	GroupId int `db:"id"`
	MenuStr string `db:"menus"`
	Desc string `db:"desc"`
	AuthStr string `db:"auths"`
}

type DBGroupListInfo struct {
	GroupId int `db:"id"`
	Desc string `db:"desc"`
}

type DBBlogCategoryInfo struct {
	CategoryId int `db:"category_id"`
	Category string `db:"desc"`
}

type DBBlogAllInfo struct {
	Id int `db:"id"`
	Status int `db:"status"`
	BlogType int `db:"type"`
	UserName string `db:"username"`
	NickName string `db:"nickname"`
	BlogName string `db:"name"`
	BlogUrl string `db:"url"`
	CategoryId int `db:"category_id"`
	Content string `db:"content"`
	CreateTm int `db:"create_tm"`
	PublishTm int `db:"publish_tm"`
	UpdateTm string `db:"update_tm"`
	ArticleCls string `db:"article_cls"`
}

type DBUserResInfo struct {
	Id int `db:"id"`
	UserName string `db:"username"`
	ResType int `db:"res_type"`
	ResDesc string `db:"res_desc"`
	ResName string `db:"res_name"`
	CreateTm int `db:"create_tm"`
	UpdateTm string `db:"update_tm"`
}