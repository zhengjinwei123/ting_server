package userservice

import (
	"app/yut/dbservice/dbpagehelper"
	"app/yut/dbservice/dbuserresservice"
	"app/yut/dbservice/dbuserservice"
	"app/yut/manager/userManager"
	"app/yut/proto/dbproto"
	"app/yut/proto/netproto"
	"app/yut/utils/fileutils"
	"app/yut/utils/orm"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	request := &netproto.NetUserLoginRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetUserLoginResponse{}
	response.UserName = request.UserName

	if userManager.GetUsrSessionMgr().UserHasLogin(r) {
		response.Msg = "user has login";
		response.ResponseError(w);
		return;
	}

	var dbUser dbproto.DBUserInfo
	if err := dbuserservice.GetUser(request.UserName, request.Password, &dbUser); err != nil {
		response.Msg = "user [" +request.UserName + "] not found"
		response.ResponseError(w);
		return
	}

	err = userManager.GetUsrSessionMgr().SetUserLogin(&dbUser, dbUser.GroupId, w, r)
	if err != nil {
		response.Msg = "user login failed"
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	response := netproto.NetUserLoginResponse{
		UserName: "",
	}

	_ = userManager.GetUsrSessionMgr().SetUserLogout(w, r)
	response.ResponseSuccess(w)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserRegisterRequest{}


	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	err = dbuserservice.RegisterUser(request.UserName, request.Password, request.Email, request.GroupId)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func UserRegisterPup(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserRegisterRequest{}


	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	err = dbuserservice.RegisterUser(request.UserName, request.Password, request.Email, 2)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func MenuList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetMenuListResponse {}

	menuList := userManager.GetUsrSessionMgr().GetMenuList(r)

	for _, v := range menuList {
		t := &netproto.Menu {
			Id: v.Id,
			Icon: v.Icon,
			Desc: v.Desc,
			Sort: v.Sort,
			ParentId: v.ParentId,
			Link: v.Link,
		}
		response.MenuList = append(response.MenuList, t)
	}

	response.ResponseSuccess(w)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetUserListResponse {}

	var dbUserList []*dbproto.DBUserInfo

	err := dbuserservice.GetUserList(&dbUserList)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	for _, v := range dbUserList {
		t := &netproto.NetUserDetail {
			UserName: v.UserName,
			GroupId: v.GroupId,
			Email: v.Email,
		}
		response.UserList = append(response.UserList, t)
	}

	response.ResponseSuccess(w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserUpdateRequest{}


	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	if request.UserName == "admin" {
		response.Msg = "not auth"
		response.ResponseError(w)
		return
	}

	err = dbuserservice.UserUpdate(request.UserName, request.Email, request.GroupId)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserDeleteRequest{}


	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	if request.UserName == "admin" {
		response.Msg = "not auth"
		response.ResponseError(w)
		return
	}

	err = dbuserservice.UserDelete(request.UserName)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserUpdatePasswordRequest{}


	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	if request.UserName == "admin" {
		response.Msg = "not auth"
		response.ResponseError(w)
		return
	}

	err = dbuserservice.UserUpdatePassword(request.UserName, request.Password)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func DelImage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserImageDelRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	username := userManager.GetUsrSessionMgr().GetUserName(r)

	response := &netproto.NetResponse{}

	err = helperDelUserImage(username, request.ImageName)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func UploadRes(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserUploadResRequest{}
	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	err = dbuserresservice.UploadRes(username, request.ResType, request.ResName, request.ResDesc)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}
	response.ResponseSuccess(w)
}

func ResDelete(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	request := &netproto.NetUserDeleteResRequest{}
	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	response := &netproto.NetResponse{}

	var dbRes dbproto.DBUserResInfo
	err = dbuserresservice.GetUserResById(request.Id, &dbRes)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}
	username := userManager.GetUsrSessionMgr().GetUserName(r)
	err = helperDelUserImage(username, dbRes.ResName)
	if err != nil {
		l4g.Error(err.Error())
	}

	err = dbuserresservice.DeleteRes(username, request.Id)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}
	response.ResponseSuccess(w)
}

func ResListPageNateSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetGetUserResPagenateSearchRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	pageHelper := dbpagehelper.NewPageHelper(request.CurPage)

	response := &netproto.NetGetUserResPagenateSearchReponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	var dbResList []*dbproto.DBUserResInfo
	err = pageHelper.SearchUserRes(username,request.ResType, &dbResList)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseSuccess(w)
		return;
	}

	response.CurPage = pageHelper.CurPage;
	response.TotalPage = pageHelper.TotalPage()
	for _, v := range dbResList {
		detail := netproto.NetResAllDetail{
			Id: v.Id,
			ResType: v.ResType,
			ResDesc: v.ResDesc,
			ResName: v.ResName,
			CreateTm: v.CreateTm,
			UpdateTm: v.UpdateTm,
			Url: helperGetImagePath(username, v.ResName),
		}

		response.UserResList = append(response.UserResList, &detail)
	}
	response.ResponseSuccess(w)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetUserUpdateProfileRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	err = dbuserservice.UserUpdateProfile(username, request.NickName,
		request.Sex, request.UserDesc, request.WxImageName, request.ZfImageName)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	request := &netproto.NetUserProfileRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		_ = l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	response := &netproto.NetUserProfileResponse{}

	var dbProfile dbproto.DBUserProfile
    err = dbuserservice.GetUserProfile(request.UserName, &dbProfile)
    if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
    	return
	}

    response.NickName = dbProfile.NickName
    response.UserDesc = dbProfile.UserDesc
    response.Sex = dbProfile.Sex
    response.WXImage = dbProfile.WXImage
    response.ZFImage = dbProfile.ZFImage

    response.ResponseSuccess(w)
}

func UploadImage(w http.ResponseWriter, r *http.Request) {


	response := &netproto.NetUserImageUploadResponse{}

	p, _ := filepath.Abs(filepath.Dir("./public/"))

	err := r.ParseMultipartForm(32 << 20);
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	username := r.FormValue("username")
	if username == "" || !userManager.GetUsrSessionMgr().ValidUser(username) {
		response.Msg = "invalid request,please login first"
		response.ResponseError(w)
		return
	}

	file, header, err := r.FormFile("upload_file");
	defer file.Close();
	if err != nil {
		l4g.Error("%s", err.Error())
	}

	if !fileutils.IsDir(p + "/upload") {
		err = os.MkdirAll(p + "/upload", os.ModePerm);
		if err != nil {
			response.Msg = err.Error()
			response.ResponseError(w)
			return
		}
	}

	uploadDir := p + "/upload/" + username + "/"
	//创建上传目录
	if !fileutils.IsDir(uploadDir) {
		err = os.MkdirAll(uploadDir, os.ModePerm);
		if err != nil {
			response.Msg = err.Error()
			response.ResponseError(w)
			return
		}
	}
	//创建上传文件

	ss := strings.Split(header.Filename, ".")
	if len(ss) < 2 {
		response.Msg = "invalid file format,file no suffix"
		response.ResponseError(w)
		return
	}
	datetime := time.Now().Format("2006-01-02-15-04-05")

	newFileName := fmt.Sprintf("%s_%s.%s", ss[0], datetime, ss[len(ss)-1])

	filePath := uploadDir + newFileName
	if fileutils.Exists(filePath) {
		response.Msg = "file " + newFileName + " exists"
		response.ResponseError(w)
		return
	}

	cur, err := os.Create(filePath);
	defer cur.Close();
	if err != nil {
		l4g.Error("err: %s", err.Error())
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}
	//把上传文件数据拷贝到我们新建的文件
	io.Copy(cur, file);

	response.ImagePath = helperGetImagePath(username, newFileName) //"/upload/"+username+"/"+newFileName

	response.ResponseSuccess(w)
}