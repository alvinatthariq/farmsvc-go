package controllers

import (
	"github.com/alvinatthariq/farmsvc-go/domain"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type controller struct {
	gorm   *gorm.DB
	router *mux.Router
	domain domain.DomainItf
}

func Init(gorm *gorm.DB, router *mux.Router, domain domain.DomainItf) {
	var c *controller

	c = &controller{
		gorm:   gorm,
		router: router,
		domain: domain,
	}

	c.Serve()
}

func (c *controller) Serve() {
	// farm
	c.router.HandleFunc("/v1/farm", c.GetFarm).Methods("GET")
	c.router.HandleFunc("/v1/farm/{id}", c.GetFarmById).Methods("GET")
	c.router.HandleFunc("/v1/farm", c.CreateFarm).Methods("POST")
	c.router.HandleFunc("/v1/farm/{id}", c.UpdateFarm).Methods("PUT")
	c.router.HandleFunc("/v1/farm/{id}", c.DeleteFarm).Methods("DELETE")
}
