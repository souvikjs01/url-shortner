package router

import (
	"net/http"
	"url-shortner/constant"
	"url-shortner/controller"
)

var urlsShortner = Routes{
	Route{"Url short service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectUrl},
}
