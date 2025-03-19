package routes

import (
	"github.com/gin-gonic/gin"
)

func Controllers(r *gin.Engine) {
	r.POST("/schools", CreateSchool)
	r.GET("/schools", GetSchools)
	r.GET("/schools/:id", GetSchoolByID)
	r.PUT("/schools/:id", UpdateSchool)
	r.DELETE("/schools/:id", DeleteSchool)
}