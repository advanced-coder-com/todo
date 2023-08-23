package jsondbprovider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	model "todo/Model"
)

const taskFilename = "tasks.json"

var _tasktsList model.Tasks

func AddTask(newTask model.Task) {

	var rawDataIn = _ReadFile()

	_Unmarshal(rawDataIn)

	_tasktsList.Tasks = append(_tasktsList.Tasks, newTask)

	_SaveFile()
}

func GetList() *model.Tasks {
	_Unmarshal(_ReadFile())
	return &_tasktsList
}

func _ReadFile() []byte {
	rawDataIn, err := ioutil.ReadFile(taskFilename)
	if err != nil {
		log.Fatal("Cannot load tasks:", err)
	}
	return rawDataIn
}

func _Unmarshal(rawDataIn []byte) {
	var err = json.Unmarshal(rawDataIn, &_tasktsList)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
}

func SetAsComplete(id int) {
	var tasksLists *model.Tasks = GetList()
	var taskToComplete *model.Task = &tasksLists.Tasks[id]
	taskToComplete.Complete = true
	_SaveFile()
}

func _SaveFile() {
	rawDataOut, err := json.MarshalIndent(&_tasktsList, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(taskFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
