package dbpagehelper

import (
	"app/yut/dbservice/dbblogservice"
	"app/yut/dbservice/dbuserresservice"
	"app/yut/proto/dbproto"
	"errors"
)

const PER_PAGE_NUM  = 5

type DBPageHelper struct {
	CurPage int

	totalCount int
	totalPage int
	hasInit bool
}

func NewPageHelper(cur_page int) *DBPageHelper {
	return &DBPageHelper{
		CurPage: cur_page,
		totalPage: 0,
		totalCount: 0,
		hasInit: false,
	}
}

func (this *DBPageHelper) TotalPage() int {
	if this.hasInit {
		return this.totalPage
	}

	page_count := this.totalCount / PER_PAGE_NUM
	e_count := this.totalCount % PER_PAGE_NUM
	if e_count > 0 {
		page_count += 1
	}

	this.totalPage = page_count
	this.hasInit = true

	return page_count
}


func (this *DBPageHelper) SearchBlog(username string, blogList *[]*dbproto.DBBlogAllInfo) error {
	total_count, err := dbblogservice.PageNateTotalCount(username)
	if err != nil {
		return err
	}

	this.totalCount = total_count

	total_page := this.TotalPage()
	if total_page < this.CurPage {
		return errors.New("page not exists")
	}

	err = dbblogservice.PageNateSearchUserBlog(username, (this.CurPage-1)*PER_PAGE_NUM, PER_PAGE_NUM, blogList)
	return err
}

func (this *DBPageHelper) SearchUserRes(username string, res_type int, resList *[]*dbproto.DBUserResInfo) error {
	total_count, err := dbuserresservice.PageNateTotalCount(username, res_type)
	if err != nil {
		return err
	}

	this.totalCount = total_count

	total_page := this.TotalPage()
	if total_page < this.CurPage {
		return errors.New("page not exists")
	}

	err = dbuserresservice.PageNateSearch(username, res_type, (this.CurPage-1)*PER_PAGE_NUM, PER_PAGE_NUM, resList)
	return err
}