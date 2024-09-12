package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yash-sojitra/todo/internal/models"
	"github.com/yash-sojitra/todo/internal/utils"
	"github.com/yash-sojitra/todo/services"
)

type PathInfo struct {
	Path string `json:"path"`
	Description string `json:"description"`
}

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *taskHandler {
	return &taskHandler{taskService: taskService}
}

func (h *taskHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/hello", h.Hello).Methods("GET")
	router.HandleFunc("/info", h.Info).Methods("GET")
	router.HandleFunc("/tasks", h.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", h.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
}

func (h *taskHandler) Info(w http.ResponseWriter, r *http.Request) {

	JSONResponse := []PathInfo {
		{
			Path: "GET http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/hello",
			Description: "gives hello world as a message in response",
		},
		{
			Path: "GET http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/tasks",
			Description: "gives all tasks present in the database.",
		},
		{
			Path: "GET http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/tasks/1",
			Description: "gives all task whose ID is 1",
		},
		{
			Path: "POST http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/tasks",
			Description: "creates task with following body {\"title\": \"api todo\",\"description\": \"add auth\"}",
		},
		{
			Path: "DELETE http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/tasks/1",
			Description: "deletes task specified by id",
		},
		{
			Path: "PUT http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/tasks/1",
			Description: "Updates the task with new given task to the task mentioned by id",
		},
		{
			Path: "GET http://ec2-43-204-97-110.ap-south-1.compute.amazonaws.com/api/v1/info",
			Description: "Gives all these info",
		},
	}


	if err := utils.WriteJSON(w,http.StatusOK,&JSONResponse); err != nil {
		utils.WriteError(w, http.StatusBadRequest,err)
		return
	}
}

func (h *taskHandler) Hello(w http.ResponseWriter, r *http.Request) {
	if err := utils.WriteJSON(w,http.StatusOK,"hello world"); err != nil {
		utils.WriteError(w, http.StatusBadRequest,err)
		return
	}
}

func (h *taskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	
	tasks ,err := h.taskService.GetAllTasks(); 
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	if err := utils.WriteJSON(w,http.StatusOK, &tasks); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

}
func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request)  {

	var task models.Task
	if err := utils.ParseJSON(r, &task); err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	if err := h.taskService.CreateTask(&task); err != nil {
		utils.WriteError(w,http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *taskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.taskService.GetTaskByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

}
func (h *taskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var task models.Task
	if err := utils.ParseJSON(r, &task); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	task.ID = id

	if err := h.taskService.UpdateTask(&task); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
func (h *taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.taskService.DeleteTask(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
