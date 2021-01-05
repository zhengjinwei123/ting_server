package globalservice

import (
	"app/yut/dbservice/dbgroupservice"
	"app/yut/manager/authManager"
	"app/yut/manager/menuManager"
	"app/yut/manager/userManager"
	"app/yut/proto/dbproto"
	"app/yut/proto/netproto"
	"app/yut/utils/orm"
	l4g "github.com/alecthomas/log4go"
	"net/http"
)

func GroupList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetGroupResponse{}

	var dbGroupList []*dbproto.DBGroupInfo
	err := dbgroupservice.GetGroupList(&dbGroupList)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	for _, v := range dbGroupList {
		groupInfo := &netproto.Group{
			Id: v.GroupId,
			Desc: v.Desc,
		}
		response.GroupList = append(response.GroupList, groupInfo)
	}

	response.ResponseSuccess(w)
}

func GroupDetailList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	response := &netproto.NetGroupDetailListResponse{}

	var dbGroupList []*dbproto.DBGroupInfo
	err := dbgroupservice.GetGroupList(&dbGroupList)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}
	for _, v := range dbGroupList {
		groupDetail := &netproto.GroupDetail{
			Id: v.GroupId,
			Desc: v.Desc,
			Menus: v.MenuStr,
			Auths: v.AuthStr,
		}
		response.GroupList = append(response.GroupList, groupDetail)
	}

	response.ResponseSuccess(w)
}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetGroupAddRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	err = dbgroupservice.AddGroup(request.GroupName, request.Menus, request.Auths)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func GroupAuthUpdate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetGroupAuthUpdateRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	err = dbgroupservice.GroupAuthUpdate(request.GroupId, request.Menus, request.Auths)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	userManager.GetUsrSessionMgr().ReloadGroupAuth(r, request.GroupId)

	response.ResponseSuccess(w)
}

func AuthList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	response := &netproto.NetAuthListResponse{}

	auth_map := authManager.GetAuthManager().GetAuthMap()

	for _, vList := range auth_map {
		for _, v := range vList {
			auth := &netproto.Auth{
				Id: v.Id,
				Desc: v.Desc,
				Url: v.Url,
			}

			response.AuthList = append(response.AuthList, auth)
		}
	}

	menue_map := menuManager.GetAllMenuMap()
	for _, v := range menue_map {
		auth := &netproto.Auth{
			Id: v.Id,
			Desc: v.Desc,
			Url: "",
		}

		response.MenuList = append(response.MenuList, auth)
	}

	response.ResponseSuccess(w)
}

func GroupDelete(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetGroupDeleteRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	if (!userManager.GetUsrSessionMgr().IsAdmin(r)) {
		response.Msg = "no auth"
		response.ResponseError(w)
		return;
	}

	err = dbgroupservice.GroupDelete(request.GroupId)
	if err != nil {
		response.Msg = err.Error();
		response.ResponseError(w)
		return
	}

	userManager.GetUsrSessionMgr().ReloadGroupAuth(r, request.GroupId)

	response.ResponseSuccess(w)
}