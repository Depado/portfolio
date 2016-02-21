package utils

import (
	"html/template"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

// InitAssetsTemplates initializes the router to use the rice boxes.
// r is our main router, tbox is our template rice box, abox is our assets box
// and names are the file names of the templates to load
func InitAssetsTemplates(r *gin.Engine, tbox, abox *rice.Box, names ...string) error {
	var err error

	if tbox != nil {
		var tmpl string
		var message *template.Template
		for _, x := range names {
			if tmpl, err = tbox.String(x); err != nil {
				return err
			}
			if message, err = template.New(x).Parse(tmpl); err != nil {
				return err
			}
			r.SetHTMLTemplate(message)
		}
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	if abox != nil {
		r.StaticFS("/static", abox.HTTPBox())
	} else {
		r.Static("/static", "assets")
	}
	return nil
}
