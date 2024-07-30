package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo/data"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func ConnectDB() {
	dbClient, err := data.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	db = dbClient.DB
	if err = db.Ping(); err != nil {
		log.Fatalf("error verifying ping of db: %v", err)
	}
	fmt.Println("Succesful connection")
}

type Task struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	DateLimit time.Time `json:"date_limit"`
}

func GetTasks(c echo.Context) error {
	var task Task
	var tasks []Task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return c.String(http.StatusInternalServerError, "get tasks method failed")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Title, &task.Status, &task.DateLimit)
		if err != nil {
			return c.String(http.StatusInternalServerError, "get rows process failed: "+err.Error())
		}
		log.Printf("\nID: %v\ntask: %v\nstatus: %v\ndate limit: %v", task.Id, task.Title, task.Status, task.DateLimit)
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return c.String(http.StatusInternalServerError, "error in rows during iteration: "+err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

type PostBody struct {
	Task      string    `json:"task"`
	DateLimit time.Time `json:"date_limit"`
}

func PostTask(c echo.Context) error {
	var body PostBody
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "body binding error")
	}
	log.Printf("task: %v, status: 'Not started', date_limit: %v", body.Task, body.DateLimit)

	stmt, err := db.Prepare("INSERT INTO tasks (task,status,datelimit) VALUES($1,$2,$3)")
	if err != nil {
		return c.String(http.StatusBadRequest, "post statement error: "+err.Error())
	}
	_, err = stmt.Exec(body.Task, "Not started", body.DateLimit)
	if err != nil {
		return c.String(http.StatusBadRequest, "post method failed "+err.Error())
	}
	return c.String(http.StatusOK, "Task added")
}

type PatchBody struct {
	Task      string    `json:"task"`
	Status    string    `json:"status"`
	DateLimit time.Time `json:"date_limit"`
}

func PatchTask(c echo.Context) error {
	var body PatchBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "body binding error")
	}
	id := c.Param("id")
	var task Task
	if err := db.QueryRow(`SELECT * FROM tasks WHERE id = $1`, id).Scan(&task.Id, &task.Title, &task.Status, &task.DateLimit); err != nil {
		c.String(http.StatusNotFound, "couldn't get the row: "+err.Error())
	}

	stmt, err := db.Prepare("UPDATE tasks SET task = $1, status = $2, datelimit = $3 WHERE id = $4")
	if err != nil {
		return c.String(http.StatusInternalServerError, "path statement error: "+err.Error())
	}
	if body.Task == "" {
		body.Task = task.Title
	}
	if body.Status == "" {
		body.Status = task.Status
	}
	if body.DateLimit.IsZero() {
		body.DateLimit = task.DateLimit
	}
	_, err = stmt.Exec(body.Task, body.Status, body.DateLimit, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "patch method couldn't be realized: "+err.Error())
	}
	return c.String(http.StatusOK, "Task saved")
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't delete the task: "+err.Error())
	}
	return c.String(http.StatusOK, "Task deleted")
}
