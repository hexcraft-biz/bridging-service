package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/hexcraft-biz/topic-management-service/config"
	"github.com/hexcraft-biz/topic-management-service/models"
	"github.com/hexcraft-biz/controller"
	"github.com/hexcraft-biz/model"
)

type EndpointTopicRels struct {
	*controller.Prototype
	Config *config.Config
}

func NewEndpointTopicRels(cfg *config.Config) *EndpointTopicRels {
	return &EndpointTopicRels{
		Prototype: controller.New("endpointTopicRels", cfg.DB),
		Config:    cfg,
	}
}

func (ctrl *EndpointTopicRels) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
	}
}

type ListEtrParams struct {
	EndpointID   string `form:"endpointId" binding:"omitempty,uuid`
	EndpointPath string `form:"endpointPath" binding:"omitempty,min=5,max=128"`
	TopicID      string `form:"topicId" binding:"omitempty,uuid`
	TopicName    string `form:"topicName" binding:"omitempty,min=5,max=256"`
	Limit        uint64 `form:"limit,default=20"`
	Offset       uint64 `form:"offset,default=0"`
}

func (ctrl *EndpointTopicRels) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		listParams := new(ListEtrParams)
		if err := c.ShouldBindQuery(listParams); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return
		}

		ctx, rdb, etrs := context.Background(), ctrl.Config.Redis, []*models.EndpointTopicRel{}

		rkey := listParams.EndpointID + "-" + listParams.EndpointPath + "-" + listParams.TopicID + "-" + listParams.TopicName + "-" + strconv.FormatUint(listParams.Offset, 10) + "-" + strconv.FormatUint(listParams.Limit, 10)
		sum := sha256.Sum256([]byte(rkey))
		hashKey := fmt.Sprintf("%x", sum)

		val, rErr := rdb.Get(ctx, hashKey).Result()
		if rErr == redis.Nil {
			pg := model.NewPagination(listParams.Offset, listParams.Limit)

			q := models.EtrListQuery{
				EndpointID:   listParams.EndpointID,
				EndpointPath: listParams.EndpointPath,
				TopicID:      listParams.TopicID,
				TopicName:    listParams.TopicName,
			}

			if ret, err := models.NewEndpointTopicRelsTableEngine(ctrl.DB).List(q, pg); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			} else {
				jsondata, _ := json.Marshal(ret)
				// exp to ENV
				if err := rdb.Set(ctx, hashKey, string(jsondata), 3600*time.Second).Err(); err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
					return
				}

				c.AbortWithStatusJSON(http.StatusOK, ret)
				return
			}
		} else if rErr != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": rErr.Error()})
			return
		} else {
			json.Unmarshal([]byte(val), &etrs)
			c.AbortWithStatusJSON(http.StatusOK, etrs)
			return
		}

	}
}

type TargetEndpointTopicRel struct {
	ID string `uri:"id" binding:"required"`
}

func (ctrl *EndpointTopicRels) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var targetETR TargetEndpointTopicRel
		if err := c.ShouldBindUri(&targetETR); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if entityRes, err := models.NewEndpointTopicRelsTableEngine(ctrl.DB).GetByID(targetETR.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			if entityRes == nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusOK, entityRes)
				return
			}
		}
	}
}

type createEndpointTopicRelParams struct {
	EndpointId string `json:"endpointId" binding:"required"`
	TopicId    string `json:"topicId" binding:"required"`
}

func (ctrl *EndpointTopicRels) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params createEndpointTopicRelParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if entityRes, err := models.NewEndpointTopicRelsTableEngine(ctrl.DB).Insert(params.EndpointId, params.TopicId); err != nil {
			if myErr, ok := err.(*mysql.MySQLError); ok && myErr.Number == 1062 {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": http.StatusText(http.StatusConflict)})
				return
			} else if ok && myErr.Number == 1452 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		} else {
			ctx := context.Background()
			ctrl.Config.Redis.FlushDB(ctx)

			c.AbortWithStatusJSON(http.StatusCreated, entityRes)
			return
		}
	}
}

func (ctrl *EndpointTopicRels) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		req, dbEngine := new(TargetEndpointTopicRel), models.NewEndpointTopicRelsTableEngine(ctrl.DB)

		if err := c.ShouldBindUri(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return
		}

		if rel, err := dbEngine.GetByID(req.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else if rel == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
			return
		} else if _, err := dbEngine.DeleteByID(req.ID); err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1451 {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		} else {
			ctx := context.Background()
			ctrl.Config.Redis.FlushDB(ctx)

			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"message": http.StatusText(http.StatusNoContent)})
			return
		}
	}
}
