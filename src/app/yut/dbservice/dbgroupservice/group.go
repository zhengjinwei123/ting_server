package dbgroupservice

import (
	"app/yut/manager/mysqlManager"
	"app/yut/proto/dbproto"
	"errors"
	"fmt"
)

const table_name = "t_group"

func GetMenuList(groupId int, group_info *dbproto.DBGroupInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryOne(fmt.Sprintf("select * from `%s` where id='%d'",
		table_name, groupId), group_info)
	if err != nil {
		return err
	}

	return nil
}

func GetGroupList(group_list *[]*dbproto.DBGroupInfo) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.QueryList(fmt.Sprintf("select * from `%s`",
		table_name), group_list)
	return err
}

func GetAuthList(groupId int, group_info *dbproto.DBGroupInfo) error {
	return GetMenuList(groupId, group_info)
}

func AddGroup(group_name, menus, auths string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err := proxy.Insert(fmt.Sprintf("insert into `%s` (`desc`,`menus`,`auths`) values('%s','%s','%s')",
		table_name, group_name, menus, auths))

	return err
}

func GroupAuthUpdate(group_id int, menus, auths string) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Update(fmt.Sprintf("update `%s` set `menus`='%s',`auths`='%s' where `id`=%d",
		table_name, menus, auths, group_id))
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("update failed: data not changed")
	}
	return nil
}

func GroupDelete(group_id int) error {
	proxy := mysqlManager.GetMysqlProxy()

	err, n := proxy.Delete(fmt.Sprintf("delete from `%s` where `id`=%d",
		table_name, group_id))
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("update failed: data not changed")
	}
	return nil
}