package dbblogcategoryservice

import (
	"app/yut/manager/mysqlManager"
	"app/yut/proto/dbproto"
	"fmt"
)

const table_name = "t_blog_category"

func AddCategory(username string, category string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.Insert(fmt.Sprintf("insert into `%s` (`username`,`desc`) values('%s','%s')",
		table_name, username, category))

	return err
}

func GetCategoryList(username string, categoryList *[]*dbproto.DBBlogCategoryInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select `category_id`,`desc` from `%s` where username='%s'",
		table_name, username), categoryList)

	return err
}

