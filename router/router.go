package router

import "github.com/gin-gonic/gin"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}
type routes struct {
	router *gin.Engine
}
type Routes []Route
