package api

import (
	"github.com/devlucky/fakelink/src/links"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(config *Config) *httprouter.Router {
	router := httprouter.New()
	router.OPTIONS("/*path", InjectConfig(config, CORS))
	router.GET("/links/:slug", InjectConfig(config, GetLink))
	router.POST("/links", InjectConfig(config, PostLink))

	return router
}

func inMemoryConf() *Config {
	return &Config{
		Template:  templates.Get(),
		LinkStore: links.NewInMemoryStore(),
	}
}