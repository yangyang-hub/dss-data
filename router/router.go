package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type RegInfo struct {
	Method  string
	Uri     string
	Handler gin.HandlerFunc
}

var _RegInfo = make(map[string][]RegInfo)

func RegisterRoutes(router *gin.Engine) {
	defer log.Println("success register routes")
	for key, infos := range _RegInfo {
		party := router.Group(key)
		{
			for _, info := range infos {
				switch info.Method {
				case "Get":
					party.GET(info.Uri, info.Handler)
				case "Post":
					party.POST(info.Uri, info.Handler)
				case "Put":
					party.PUT(info.Uri, info.Handler)
				case "Delete":
					party.DELETE(info.Uri, info.Handler)
				case "Options":
					party.OPTIONS(info.Uri, info.Handler)
				}
			}
		}

	}
}
func RegisterHandler(method string, party string, uri string, handler gin.HandlerFunc) {
	info := RegInfo{method, uri, handler}
	infos, ok := _RegInfo[party]
	if !ok {
		infos = []RegInfo{}
	}
	infos = append(infos, info)
	_RegInfo[party] = infos
}
