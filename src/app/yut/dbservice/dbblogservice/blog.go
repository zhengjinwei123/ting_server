package dbblogservice

import (
	"app/yut/manager/mysqlManager"
	"app/yut/proto/dbproto"
	"errors"
	"fmt"
	"time"
)

const table_name = "t_blog"

func AddBlog(username string, category_id int, content string, blog_name string, blog_type int) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.Insert(fmt.Sprintf("insert into `%s` (`username`,`category_id`, `name`,`content`,`create_tm`,`type`) values('%s', '%d', '%s','`%s`',%d,%d)",
		table_name, username, category_id, blog_name, content, time.Now().Unix(), blog_type))

	return err
}

func DelBlog(id int) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Delete(fmt.Sprintf("delete from `%s` where `id`=%d",
		table_name, id))
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("update failed: data not changed")
	}
	return nil
}


func UpdateBlog(blog_id int, content string, blog_name string, category_id int) error {
	proxy := mysqlManager.GetMysqlProxy()
	err, n := proxy.Update(fmt.Sprintf("update `%s` set `name`='%s', `content`='%s', `category_id`=%d where `id`=%d",
		table_name, blog_name, content, category_id, blog_id))
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("update failed: data not changed")
	}
	return nil
}

func GetBlog(blog_id int, blogInfo *dbproto.DBBlogAllInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	sql := fmt.Sprintf("select C.*, D.`nickname` as nickname from (select A.*, B.`desc` as article_cls FROM (select * from t_blog where id=%d) as A " +
		"LEFT JOIN (select `category_id`,`desc`,`username` from t_blog_category) as B ON A.`category_id`=B.`category_id` " +
		"WHERE B.`category_id` IS NOT NULL) as C JOIN (select `nickname`,`username` from t_user) as D ON C.`username`=D.`username`", blog_id)
	err := proxy.QueryOne(sql, blogInfo)

	fmt.Println(sql)

	return err
}

func GetUserBlogsByStatus(username string, status int, blogList *[]*dbproto.DBBlogAllInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s` where `username`='?' and `status`=%d",
		table_name, username, status), blogList)
	return err
}


func OneKeyPublish(username string) error {
	proxy := mysqlManager.GetMysqlProxy()
	err, _ := proxy.Update(fmt.Sprintf("update `%s` set `status`=%d, `publish_tm`=%d, " +
		"`url`=concat('/pup/blog/', cast(`id` as char)) where `username`='%s' and `status`=%d",
		table_name, 1, time.Now().Unix(), username, 0))
	if err != nil {
		return err
	}

	return nil
}

func OneKeyClose(username string) error {
	proxy := mysqlManager.GetMysqlProxy()
	err, _ := proxy.Update(fmt.Sprintf("update `%s` set `status`=%d " +
		"where `username`='%s' and `status`=%d",
		table_name, 0, username, 1))
	if err != nil {
		return err
	}

	return nil
}


func GetUserBlogList(username string, blogList *[]*dbproto.DBBlogAllInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s` where `username`='?'",
		table_name, username), blogList)
	return err
}

func PublishBlog(blog_id int, url string, status int) error {
	proxy := mysqlManager.GetMysqlProxy()
	err, n := proxy.Update(fmt.Sprintf("update `%s` set `status`=%d, `publish_tm`=%d, `url`='%s' where `id`=%d",
		table_name, status, time.Now().Unix(), url, blog_id))
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("update failed: data not changed")
	}
	return nil
}

func PageNateTotalCount(username string) (int, error) {
	proxy := mysqlManager.GetMysqlProxy()
	count, err := proxy.GetCount2("id", table_name, "`username`='" + username + "'")
	return count, err
}

func PageNateSearchUserBlog(username string, last_id int, limit_num int, blogList *[]*dbproto.DBBlogAllInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s` where `username`='%s' order by id asc limit %d,%d",
		table_name, username, last_id, limit_num), blogList)

	return err
}