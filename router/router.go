package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string // path
	HandlerFunc func(*gin.Context)
}
type routes struct {
	router *gin.Engine
}
type Routes []Route

func (r routes) UrlShortner(rg *gin.RouterGroup) {
	orderRouteGroup := rg.Group("/url")
	for _, route := range urlsShortner {
		switch route.Method {
		case http.MethodGet:
			orderRouteGroup.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouteGroup.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouteGroup.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouteGroup.DELETE(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			orderRouteGroup.POST(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGroup.POST(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	vl := r.router.Group(os.Getenv("API_VERSION"))
	r.UrlShortner(vl)
	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
