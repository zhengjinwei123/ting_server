package dbuserservice

import (
	"app/yut/manager/mysqlManager"
	"app/yut/proto/dbproto"
	"errors"
	"fmt"
)

const table_name = "t_user"

func GetUser(username string, password string, user_info *dbproto.DBUserInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryOne(fmt.Sprintf("select * from `%s` where username='%s' and password='%s'",
		table_name, username, password), user_info)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByName(username string, user_info *dbproto.DBUserInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryOne(fmt.Sprintf("select * from `%s` where username='%s'",
		table_name, username), user_info)
	return err
}

func RegisterUser(username string, password string, email string, group_id int) error {
	proxy := mysqlManager.GetMysqlProxy()

	count, err := proxy.GetCount(fmt.Sprintf("select * from `%s` where username='%s'", table_name, username))
	if err != nil {
		return errors.New(err.Error())
	}

	if count > 0 {
		return errors.New("user " + username + " has  registered")
	}

	err = proxy.Insert(fmt.Sprintf("insert into `%s` (`username`,`password`,`email`,`group_id`) values('%s','%s','%s',%d)",
		table_name, username, password, email, group_id))

	return err
}

func UserUpdate(username string, email string, group_id int) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Update(fmt.Sprintf("update `%s` set `email`='%s',`group_id`=%d where `username`='%s'",
		table_name, email, group_id, username))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}

func UserDelete(username string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Delete(fmt.Sprintf("delete from `%s` where `username`='%s'",
		table_name, username))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}

func UserUpdatePassword(username string, password string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Update(fmt.Sprintf("update `%s` set `password`='%s' where `username`='%s'",
		table_name, password, username))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}

func UserUpdateProfile(username string, nickname string, sex int, userdesc string, wx_image string, zf_image string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Update(fmt.Sprintf("update `%s` set `nickname`='%s'," +
		"`userdesc`='%s'," +
		"`wx_rcode_img`='%s'," +
		"`zf_rcode_img`='%s'," +
		"`sex`=%d where `username`='%s'",
		table_name, nickname, userdesc, wx_image, zf_image, sex, username))

	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("data not changed")
	}

	return nil
}

func GetUserProfile(username string, user_info *dbproto.DBUserProfile) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryOne(fmt.Sprintf("select `nickname`,`userdesc`,`wx_rcode_img`,`zf_rcode_img`,`sex` from `%s` where username='%s'",
		table_name, username), user_info)
	return err
}


func GetUserList(user_list *[]*dbproto.DBUserInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s`",
		table_name), user_list)
	return err
}