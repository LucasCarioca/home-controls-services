package routes

import (
	"github.com/LucasCarioca/home-controls-services/pkg/models"
	"net/http"
	"strconv"

	"github.com/LucasCarioca/home-controls-services/pkg/services"
	"github.com/gin-gonic/gin"
)

//SwitchRouter router for switch crud routes
type SwitchRouter struct {
	s *services.SwitchService
}

//CreateSwitchRequest payload structure for creating new switches
type CreateSwitchRequest struct {
	Name  string `json:"name" binding:"required"`
	State string `json:"state" binding:"required"`
}

//UpdateSwitchRequest payload structure for updating switches
type UpdateSwitchRequest struct {
	DesiredState *string `json:"desired-state"`
	State        *string `json:"state"`
}

func (r *SwitchRouter) getSwitch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid id",
		})
		return
	}

	sw, err := r.s.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "could not find switch",
			"error":   "SWITCH_NOT_FOUND",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, sw)
}

func (r *SwitchRouter) deleteSwitch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid id",
		})
		return
	}

	err = r.s.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "could not delete switch",
			"error":   "SWITCH_DELETE_FAILED",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "switch delete successfully",
	})
}

func (r *SwitchRouter) getAllSwitches(ctx *gin.Context) {
	switches := r.s.GetAll()
	ctx.JSON(http.StatusOK, switches)
}

func (r *SwitchRouter) createSwitch(ctx *gin.Context) {
	var data CreateSwitchRequest
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing or incorrect fields received",
			"error":   "CREATE_SWITCH_PAYLOAD_INVALID",
			"details": err.Error(),
		})
		return
	}
	if state, ok := models.SwitchStates[data.State]; ok {
		sw := r.s.Create(data.Name, state)
		ctx.JSON(http.StatusOK, sw)
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "state value must be CLOSED or OPEN",
	})
}

func (r *SwitchRouter) updateSwitch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid id",
		})
		return
	}

	var data UpdateSwitchRequest
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing or incorrect fields received",
			"error":   "CREATE_SWITCH_PAYLOAD_INVALID",
			"details": err.Error(),
		})
		return
	}

	if data.DesiredState != nil && data.State != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Updating desired state and state at the same time is forbidden",
		})
		return
	}

	if data.DesiredState != nil {
		if desiredState, ok := models.SwitchStates[*data.DesiredState]; ok {
			sw, err := r.s.UpdateDesiredStateByID(id, desiredState)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "missing or incorrect fields received",
					"details": err.Error(),
				})
			}
			ctx.JSON(http.StatusOK, sw)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "state value must be CLOSED or OPEN",
		})
		return
	}

	if data.State != nil {
		if state, ok := models.SwitchStates[*data.State]; !ok {
			sw, err := r.s.UpdateStateByID(id, state)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "missing or incorrect fields received",
					"details": err.Error(),
				})
			}
			ctx.JSON(http.StatusOK, sw)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "state value must be CLOSED or OPEN",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Nothing to update",
	})
}

// NewSwitchRouter router for the switch endpoint
func NewSwitchRouter(router *gin.Engine) {
	r := SwitchRouter{
		s: services.NewSwitchService(),
	}
	router.Use(checkKeyMiddleware)
	router.GET("/api/v1/switches", r.getAllSwitches)
	router.POST("/api/v1/switches", r.createSwitch)
	router.PUT("/api/v1/switches/:id", r.updateSwitch)
	router.GET("/api/v1/switches/:id", r.getSwitch)
	router.DELETE("/api/v1/switches/:id", r.deleteSwitch)
}
