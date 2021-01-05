package dbuserresservice

import (
	"app/yut/manager/mysqlManager"
	"app/yut/proto/dbproto"
	"errors"
	"fmt"
	"time"
)

const table_name = "t_user_res"

func GetUserResList(username string, user_res_list *[]*dbproto.DBUserResInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s` where `username`='%s'",
		table_name, username), user_res_list)
	return err
}

func GetUserResById(id int, user_res *dbproto.DBUserResInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryOne(fmt.Sprintf("select * from `%s` where `id`=%d",
		table_name, id), user_res)
	return err
}

func UploadRes(username string, res_type int, res_name, res_desc string) error {
	proxy := mysqlManager.GetMysqlProxy()
	err := proxy.Insert(fmt.Sprintf("insert into `%s` (`username`,`res_type`,`res_name`,`res_desc`, `create_tm`) values('%s', %d,'%s','%s', %d)",
		table_name, username, res_type, res_name, res_desc, time.Now().Unix()))

	return err
}

func PageNateTotalCount(username string, res_type int) (int, error) {
	proxy := mysqlManager.GetMysqlProxy()
	where := fmt.Sprintf("`username`='%s' and `res_type`=%d", username, res_type)
	count, err := proxy.GetCount2("id", table_name, where)
	return count, err
}

func PageNateSearch(username string, res_type int, last_id int, limit_num int, resList *[]*dbproto.DBUserResInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s` where `username`='%s' and `res_type`=%d order by id asc limit %d,%d",
		table_name, username, res_type, last_id, limit_num), resList)

	return err
}

func DeleteRes(username string, id int) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Delete(fmt.Sprintf("delete from `%s` where id=%d and `username`='%s'",
		table_name, id, username))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}

func UpdateRes(id int, res_name, res_desc string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Update(fmt.Sprintf("update `%s` set `res_name`='%s'," +
		"`res_desc`='%s' " +
		"where `id`=%d",
		table_name, res_name, res_desc, id))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}