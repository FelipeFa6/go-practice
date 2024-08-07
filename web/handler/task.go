package handler

import (
	"github.com/felipefa6/todo-go/view/task"
	"github.com/felipefa6/todo-go/model"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct { }

func (h TaskHandler) HandleTaskShow(c echo.Context) error{
    tsk := model.Task {
        ID: 1,
        Description: "Congrats",
        State: true,
    }
    return render(c, task.Show(tsk))
}
