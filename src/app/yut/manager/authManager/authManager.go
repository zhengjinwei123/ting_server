package authManager

import (
	"app/yut/proto"
	"app/yut/utils"
	"errors"
	"flag"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"strconv"
	"strings"
)

var authFile = flag.String("authconfig", "../settings/menu.json", "")
var requestConf = newRequestJsonConfig()

var g_authM *authManager = nil

type requestJsonConfig struct {
	RequestList [][]string `json:"requestlist"`
}

func newRequestJsonConfig() *requestJsonConfig {
	return &requestJsonConfig{}
}

func newAuthManager() *authManager {
	authM := &authManager{}
	authM.hasInit = false
	authM.reqMap = make(map[int]*proto.ReqUrl, 0)
	authM.reqGroup = make(map[int][]*proto.ReqUrl, 0)

	return authM
}

func GetAuthManager() *authManager {
	if g_authM == nil {
		g_authM = newAuthManager()
	}
	return g_authM
}

type authManager struct {
	reqMap map[int]*proto.ReqUrl
	reqGroup map[int][]*proto.ReqUrl
	hasInit bool
}

func (this *authManager) Load() error {
	if this.hasInit == true {
		return nil
	}
	if err := utils.LoadJsonConfig(*authFile, requestConf); err != nil {
		l4g.Error("load requestUrl config %v failed: %v \n", *authFile, err)
		return err
	}

	for _, vList := range requestConf.RequestList {
		for _, v := range vList {
			vv := strings.Split(v, ",")

			if len(vv) != 3 {
				l4g.Error("requestUrl init: valid len of menu[%s] \n", v)
				continue;
			}

			id, _ := strconv.Atoi(vv[0])
			url := strings.TrimSpace(vv[1])
			desc := vv[2]

			url = "/api" + url

			_, exists := this.reqMap[id];
			if exists {
				l4g.Error("requestUrl init:repeated of url: %s \n", url)
				continue
			}

			reqUrl := &proto.ReqUrl{
				Id: id,
				Url: url,
				Desc: desc,
			}
			this.reqMap[id] = reqUrl

			mapId := int(id / 1000);
			_, exists = this.reqGroup[mapId]
			if !exists {
				this.reqGroup[mapId] = make([]*proto.ReqUrl, 0)
			}
			this.reqGroup[mapId] = append(this.reqGroup[mapId], reqUrl)
		}
	}

	l4g.Debug("request url init sucess")
	return nil
}

func (this *authManager) GetAuthById(id int) (*proto.ReqUrl, error) {
	_, exists := this.reqMap[id];
	if exists {
		return this.reqMap[id], nil
	}
	return nil, errors.New(fmt.Sprintf("GetUrlById not found id: %d",id))
}

func (this *authManager) GetAuthMap() map[int][]*proto.ReqUrl {
	return this.reqGroup
}