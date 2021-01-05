package blogservice

import (
	"app/yut/dbservice/dbblogcategoryservice"
	"app/yut/dbservice/dbblogservice"
	"app/yut/dbservice/dbpagehelper"
	"app/yut/manager/userManager"
	"app/yut/proto/dbproto"
	"app/yut/proto/netproto"
	"app/yut/utils"
	"app/yut/utils/orm"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"net/http"
)

func GetBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetPubBlogGetResponse{}
	ctx := r.Context()
	blogId, ok := ctx.Value("id").(int)
	if !ok {
		response.ResponseError(w)
		return
	}

	var dbblog dbproto.DBBlogAllInfo
	err := dbblogservice.GetBlog(blogId, &dbblog)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	if dbblog.Status == 0 {
		response.Msg = "forbidden visit"
		response.ResponseError(w)
		return
	}

	response.UpdateTm = dbblog.UpdateTm
	response.PublishTm = dbblog.PublishTm
	response.ArticleCls = dbblog.ArticleCls
	response.ArticleName = dbblog.BlogName
	response.Content = utils.Stripslashes(dbblog.Content)
	response.Type = dbblog.BlogType
	response.UserName = dbblog.UserName
	response.NickName = dbblog.NickName

	response.ResponseSuccess(w)
}

func AddBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetBlogAddRequest{}
	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	err = dbblogservice.AddBlog(username, request.CategoryId, utils.Addslashes(request.Content), request.BlogName, request.BlogType)

	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	response.ResponseSuccess(w)
}

func GetUserCategories(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetGetUserBlogCategoryListResponse{}
	username := userManager.GetUsrSessionMgr().GetUserName(r)

	var categories []*dbproto.DBBlogCategoryInfo
	err := dbblogcategoryservice.GetCategoryList(username, &categories)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	for _, v := range categories {
		data := netproto.NetBlogCategoryDetail{
			CategoryId: v.CategoryId,
			Category: v.Category,
		}

		response.CategoryList = append(response.CategoryList, &data)
	}

	response.ResponseSuccess(w)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetBlogUpdateRequest{}
	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}
	err = dbblogservice.UpdateBlog(request.BlogId, utils.Addslashes(request.Content), request.BlogName, request.CategoryId)

	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	response.ResponseSuccess(w)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	request := &netproto.NetBlogDeleteRequest{}
	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}
	err = dbblogservice.DelBlog(request.BlogId)

	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	response.ResponseSuccess(w)

}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetBlogCategoryAddRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	err = dbblogcategoryservice.AddCategory(username, request.Category)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	response.ResponseSuccess(w)
}

func OneKeyPublish(w http.ResponseWriter, r *http.Request) {
	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)

	err := dbblogservice.OneKeyPublish(username)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func OneKeyClose(w http.ResponseWriter, r *http.Request) {
	response := &netproto.NetResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)

	err := dbblogservice.OneKeyClose(username)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}

func GetBlogPageNateSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetGetBlogPagenateSearchRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}
	pageHelper := dbpagehelper.NewPageHelper(request.CurPage)

	response := &netproto.NetGetBlogPagenateListResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)
	var dbBlogList []*dbproto.DBBlogAllInfo
	err = pageHelper.SearchBlog(username, &dbBlogList)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	response.CurPage = pageHelper.CurPage;
	response.TotalPage = pageHelper.TotalPage()
	for _, v := range dbBlogList {
		detail := netproto.NetBlogAllDetail{
			Id: v.Id,
			Status: v.Status,
			BlogType: v.BlogType,
			UserName: v.UserName,
			BlogName: v.BlogName,
			BlogUrl: v.BlogUrl,
			CategoryId: v.CategoryId,
			Content: utils.Stripslashes(v.Content),
			CreateTm: v.CreateTm,
			UpdateTm: v.UpdateTm,
			PublishTm: v.PublishTm,
		}

		response.UserBlogList = append(response.UserBlogList, &detail)
	}
	response.ResponseSuccess(w)
}

func GetBlogList(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	response := &netproto.NetGetUserBlogListResponse{}

	username := userManager.GetUsrSessionMgr().GetUserName(r)

	var dbBlogList []*dbproto.DBBlogAllInfo
	err := dbblogservice.GetUserBlogList(username, &dbBlogList)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return;
	}

	for _, v := range dbBlogList {
		detail := netproto.NetBlogAllDetail{
			Id: v.Id,
			Status: v.Status,
			BlogType: v.BlogType,
			UserName: v.UserName,
			BlogName: v.BlogName,
			BlogUrl: v.BlogUrl,
			CategoryId: v.CategoryId,
			Content: utils.Stripslashes(v.Content),
			CreateTm: v.CreateTm,
			UpdateTm: v.UpdateTm,
			PublishTm: v.PublishTm,
		}

		response.UserBlogList = append(response.UserBlogList, &detail)
	}

	response.ResponseSuccess(w)
}

func PublishBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	request := &netproto.NetBlogPublishRequest{}

	err := orm.UnmarshalHttpValues(request, r.PostForm)
	if err != nil {
		l4g.Error("UnmarshalHttpValues error: [%v] %v \n",r.PostForm, err)
		return
	}

	response := &netproto.NetResponse{}

	url := fmt.Sprintf("/pup/blog/%d", request.BlogId)

	err = dbblogservice.PublishBlog(request.BlogId, url, request.Status)
	if err != nil {
		response.Msg = err.Error()
		response.ResponseError(w)
		return
	}

	response.ResponseSuccess(w)
}