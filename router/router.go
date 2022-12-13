package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"log"
)

type RegInfo struct {
	Method  string
	Uri     string
	Handler iris.Handler
}

var _RegInfo = make(map[string][]RegInfo)

func RegisterRoutes(app *iris.Application) {
	defer log.Println("success register routes")
	for key, infos := range _RegInfo {
		party := app.Party(key, func(ctx iris.Context) {
			ctx.Next()
		})
		{
			for _, info := range infos {
				switch info.Method {
				case "Get":
					party.Get(info.Uri, info.Handler)
				case "Post":
					party.Post(info.Uri, info.Handler)
				case "Put":
					party.Put(info.Uri, info.Handler)
				case "Delete":
					party.Delete(info.Uri, info.Handler)
				case "Options":
					party.Options(info.Uri, info.Handler)
				}
			}
		}
	}
}
func RegisterHandler(method string, party string, uri string, handler context.Handler) {
	info := RegInfo{method, uri, handler}
	infos, ok := _RegInfo[party]
	if !ok {
		infos = []RegInfo{}
	}
	infos = append(infos, info)
	_RegInfo[party] = infos
}
