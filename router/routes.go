package router

import (
	"net/http"
	"url-shortner/constant"
	"url-shortner/controller"
)

var urlsShortner = Routes{
	Route{"url short service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
}
