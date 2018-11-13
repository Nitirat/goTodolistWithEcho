package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	task struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}
)

var (
	tasks = map[int]*task{}
	seq   = 1
)

//----------
// Handlers
//----------

func createTask(c echo.Context) error {
	u := &task{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	tasks[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, tasks[id])
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func updateTask(c echo.Context) error {
	u := new(task)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	tasks[id].Status = u.Status
	tasks[id].Name = u.Name
	return c.JSON(http.StatusOK, tasks[id])
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(tasks, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/tasks", createTask)
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = ":8099"
	}
	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
