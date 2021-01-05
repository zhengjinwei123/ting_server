package netproto

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SResponse struct {
	Data interface{} `json:"data"`
	Status int `json:"status"`
}

func outputJson(data interface{}, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	resp := &SResponse{
		Data: data,
		Status: status,
	}
	ret_json, _ := json.Marshal(resp)

	//log.Printf("response:data [%v] [%s] \n", resp, ret_json)

	_, err := io.WriteString(w, string(ret_json))
	if err != nil {
		log.Printf("outputJson error: %s \n", ret_json)
	}
}

/*--------------------------------网络response 基类------------------------------------------------*/
type NetResponse struct {
	Msg string `json:"msg"`
}

func (this *NetResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

/*--------------------------------网络response 子类------------------------------------------------*/

type NetInvalidParamResponse struct {
	NetResponse
}

func (this *NetInvalidParamResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_INVALID_PARAM, writer)
}

func (this *NetInvalidParamResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// NotLogin
type NetNotLoginResponse struct {
	NetResponse
}

func (this *NetNotLoginResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_NOT_LOGIN, writer)
}

func (this *NetNotLoginResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// NoAuth
type NetNoAuthResponse struct {
	NetResponse
}
func (this *NetNoAuthResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_NO_AUTH, writer)

}

func (this *NetNoAuthResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// login
type NetUserLoginRequest struct {
	UserName string `http:"username"`
	Password string `http:"password"`
}

type NetUserLoginResponse struct {
	NetResponse

	UserName string `json:"username"`
}

func (this *NetUserLoginResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetUserLoginResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// register
type NetUserRegisterRequest struct {
	UserName string `http:"username"`
	Email string `http:"email"`
	Password string `http:"password"`
	GroupId int `http:"group_id"`
}



// get menu
type Menu struct {
	Id int `json:"id"`
	Desc string `json:"desc"`
	Icon string `json:"icon"`
	ParentId int `json:"parent_id"`
	Link string `json:"link"`
	Sort int `json:"sort"`
}

type NetMenuListResponse struct {
	NetResponse

	MenuList []*Menu `json:"menulist"`
}

func (this *NetMenuListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetMenuListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// get group list
type Group struct {
	Id int `json:"id"`
	Desc string `json:"desc"`
}

type NetGroupResponse struct {
	NetResponse

	GroupList []*Group `json:"grouplist"`
}

func (this *NetGroupResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGroupResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// get auth list
type Auth struct {
	Id int `json:"id"`
	Desc string `json:"desc"`
	Url string `json:"url"`
}
type NetAuthListResponse struct {
	NetResponse

	AuthList []*Auth `json:"authlist"`
	MenuList []*Auth `json:"menulist"`
}

func (this *NetAuthListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetAuthListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}
// get group detail list
type GroupDetail struct {
	Id int `json:"id"`
	Desc string `json:"desc"`
	Menus string `json:"menus"`
	Auths string `json:"auths"`
}

type NetGroupDetailListResponse struct {
	NetResponse

	GroupList []*GroupDetail `json:"grouplist"`
}

func (this *NetGroupDetailListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGroupDetailListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// add group
type NetGroupAddRequest struct {
	GroupName string `http:"group"`
	Menus string `http:"menus"`
	Auths string `http:"auths"`
}


// update group auth
type NetGroupAuthUpdateRequest struct {
	GroupId int `http:"id"`
	Menus string `http:"menus"`
	Auths string `http:"auths"`
}


// user list
type NetUserDetail struct {
	UserName string `json:"username"`
	Email string `json:"email"`
	GroupId int `json:"group_id"`
}

type NetUserListResponse struct {
	NetResponse

	UserList []*NetUserDetail `json:"userlist"`
}

func (this *NetUserListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetUserListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}


// group delete
type NetGroupDeleteRequest struct {
	GroupId int `http:"id"`
}

// user update
type NetUserUpdateRequest struct {
	GroupId int `http:"group_id"`
	UserName string `http:"username"`
	Email string `http:"email"`
}

// user delete
type NetUserDeleteRequest struct {
	UserName string `http:"username"`
}

// user update password
type NetUserUpdatePasswordRequest struct {
	UserName string `http:"username"`
	Password string `http:"password"`
}

// add blog category
type NetBlogCategoryAddRequest struct {
	Category string `http:"category"`
}

// add blog
type NetBlogAddRequest struct {
	BlogName string `http:"blog_name"`
	BlogType int `http:"blog_type"`
	CategoryId int `http:"category_id"`
	Content string `http:"content"`
}

// get user blogs
type NetBlogAllDetail struct {
	Id int `json:"id"`
	Status int `json:"status"`
	BlogType int `json:"type"`
	UserName string `json:"username"`
	BlogName string `json:"name"`
	BlogUrl string `json:"url"`
	CategoryId int `json:"category_id"`
	Content string `json:"content"`
	CreateTm int `json:"create_tm"`
	PublishTm int `json:"publish_tm"`
	UpdateTm string `json:"update_tm"`
}

type NetGetUserBlogListResponse struct {
	NetResponse

	UserBlogList []*NetBlogAllDetail `json:"bloglist"`
}

func (this *NetGetUserBlogListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGetUserBlogListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// get category
type NetBlogCategoryDetail struct {
	CategoryId int `json:"category_id"`
	Category string `json:"desc"`
}

type NetGetUserBlogCategoryListResponse struct {
	NetResponse

	CategoryList []*NetBlogCategoryDetail `json:"categories"`
}

func (this *NetGetUserBlogCategoryListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGetUserBlogCategoryListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// publish blog
type NetBlogPublishRequest struct {
	BlogId int `http:"id"`
	Status int `http:"status"`
}

// GET BLOG
type NetPubBlogGetResponse struct {
	NetResponse

	ArticleName string `json:"article_name"`
	ArticleCls string `json:"article_cls"`
	PublishTm int `json:"publish_tm"`
	UpdateTm string `json:"update_tm"`
	Content string `json:"content"`
	Type int `json:"type"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
}

func (this *NetPubBlogGetResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetPubBlogGetResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// update blog
type NetBlogUpdateRequest struct {
	BlogId int `http:"id"`
	BlogName string `http:"blog_name"`
	CategoryId int `http:"category_id"`
	Content string `http:"content"`
}

// delete blog
type NetBlogDeleteRequest struct {
	BlogId int `http:"id"`
}

// get blog pagenate

type NetGetBlogPagenateSearchRequest struct {
	CurPage int `http:"cur_page"`
}

type NetGetBlogPagenateListResponse struct {
	NetResponse

	UserBlogList []*NetBlogAllDetail `json:"datalist"`
	CurPage int `json:"cur_page"`
	TotalPage int `json:"total_page"`
}

func (this *NetGetBlogPagenateListResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGetBlogPagenateListResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// image upload
type NetUserImageUploadResponse struct {
	NetResponse

	ImagePath string `json:"image_path"`
}

func (this *NetUserImageUploadResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetUserImageUploadResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// image del
type NetUserImageDelRequest struct {
	ImageName string `http:"image_name"`
}

// update profile
type NetUserUpdateProfileRequest struct {
	NickName string `http:"nickname"`
	UserDesc string `http:"userdesc"`
	Sex int `http:"sex"`
	WxImageName string `http:"wx_image"`
	ZfImageName string `http:"zf_image"`
}

// get profile
type NetUserProfileRequest struct {
	UserName string `http:"username"`
}
type NetUserProfileResponse struct {
	NetResponse

	NickName string `json:"nickname"`
	Sex int `json:"sex"`
	UserDesc string `json:"userdesc"`
	WXImage string `json:"wx_image"`
	ZFImage string `json:"zf_image"`
}

func (this *NetUserProfileResponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetUserProfileResponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// user upload res
type NetUserUploadResRequest struct {
	ResType int `http:"res_type"`
	ResName string `http:"res_name"`
	ResDesc string `http:"res_desc"`
}

// get res pagenate
type NetResAllDetail struct {
	Id int `json:"id"`
	ResType int `json:"res_type"`
	ResName string `json:"res_name"`
	ResDesc string `json:"res_desc"`
	CreateTm int `json:"create_tm"`
	UpdateTm string `json:"update_tm"`
	Url string `json:"url"`
}

type NetGetUserResPagenateSearchRequest struct {
	ResType int `http:"res_type"`
	CurPage int `http:"cur_page"`
}

type NetGetUserResPagenateSearchReponse struct {
	NetResponse

	UserResList []*NetResAllDetail `json:"datalist"`
	CurPage int `json:"cur_page"`
	TotalPage int `json:"total_page"`
}

func (this *NetGetUserResPagenateSearchReponse) ResponseError(writer http.ResponseWriter) {
	outputJson(*this, NET_STATUS_UNKNOWN, writer)

}

func (this *NetGetUserResPagenateSearchReponse) ResponseSuccess(writer http.ResponseWriter) {
	outputJson(*this, 0, writer)
}

// delete res
type NetUserDeleteResRequest struct {
	Id int `http:"id"`
}