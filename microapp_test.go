package bytedance

import (
	"github.com/cute-angelia/go-utils/syntax/ijson"
	"github.com/cute-angelia/go-utils/utils/conf"
	"github.com/hacku7/bytedance/microapp"
	"log"
	"testing"
)

func getApp() *microapp.Component {
	conf.LoadConfigFile("./config.toml")
	return microapp.Load("microapp")
}

func TestGetAccessToken(t *testing.T) {
	microApp := getApp()
	accessTokenInfo := microApp.GetAccessToken()
	log.Println(ijson.Pretty(accessTokenInfo))
}
