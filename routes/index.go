package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)


func index(rnd render.Render) {
	rnd.HTML(200, "write", model)
}