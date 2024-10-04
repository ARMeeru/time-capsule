package utils

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, filenames []string, data interface{}) {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing template: %v", err)
		return
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error executing template: %v", err)
	}
}
