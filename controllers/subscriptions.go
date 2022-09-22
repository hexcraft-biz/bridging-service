package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/controller"
	"github.com/hexcraft-biz/topic-management-service/config"
)

type Subscriptions struct {
	*controller.Prototype
	Config *config.Config
}

func NewSubscription(cfg *config.Config) *Subscriptions {
	return &Subscriptions{
		Prototype: controller.New("subscriptions", cfg.DB),
		Config:    cfg,
	}
}

func (ctrl *Subscriptions) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
	}
}

// TODO GCP pubsub
