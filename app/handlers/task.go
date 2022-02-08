package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-todolist/app/models"
	"go-todolist/app/usecases"
	"go-todolist/utils/constants"
	"net/http"
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
	var limit, offset int
	var err error
	rawParam, ok := context.GetQuery("limit")
	if !ok {
		limit = constants.DEFAULT_LIMIT
	} else {
		limit, err = strconv.Atoi(rawParam)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	rawParam, ok = context.GetQuery("offset")
	if !ok {
		offset = constants.DEFAULT_OFFSET
	} else {
		offset, err = strconv.Atoi(rawParam)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	tasks, code := handler.TaskHandler.All(offset, limit)
	if code != constants.OK_CODE {
		context.AbortWithStatus(code)
		return
	}

	context.JSON(http.StatusOK, tasks)
	return
}

func (handler *TaskHandler) Update(context *gin.Context) {
	task := &models.Task{}

	err := easyjson.UnmarshalFromReader(context.Request.Body, task)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	createdTask, code := handler.TaskHandler.Create(task)
	if code != constants.OK_CODE {
		context.AbortWithStatus(code)
		return
	}

	context.JSON(http.StatusCreated, createdTask)
	return
}

func (handler *TaskHandler) Toggle(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return
	}

	code := handler.TaskHandler.Toggle(id)
	if code != constants.OK_CODE {
		context.AbortWithStatus(code)
		return
	}

	context.Status(http.StatusOK)
	return
}

func (handler *TaskHandler) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return
	}

	code := handler.TaskHandler.Delete(id)
	if code != constants.OK_CODE {
		context.AbortWithStatus(code)
		return
	}

	context.Status(http.StatusOK)
	return
}
