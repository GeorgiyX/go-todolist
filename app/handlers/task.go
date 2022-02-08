package handlers

import (
	"github.com/gin-gonic/gin"
	"go-todolist/app/usecases"
	"strconv"
)

type TaskHandler struct {
	TaskHandler usecases.ITaskUseCase
}

func CreateTaskHandler(url string,
	taskUseCase usecases.ITaskUseCase,
	router *gin.RouterGroup) *TaskHandler {
	handler := &TaskHandler{
		TaskHandler: taskUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.GET("", handler.Get)
	urlGroup.POST("", handler.Update)
	urlGroup.PATCH("/:id/toggle", handler.Toggle)
	urlGroup.DELETE("/:id", handler.Delete)

	return handler
}

func (handler *TaskHandler) Get(context *gin.Context) {

}

func (handler *TaskHandler) Update(context *gin.Context) {

}

func (handler *TaskHandler) Toggle(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return
	}

}

func (handler *TaskHandler) Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return
	}

}
