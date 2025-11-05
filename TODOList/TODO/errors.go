package todo

import "errors"

var ErrTaskIsAlreadyExist = errors.New("Task is already exist")
var ErrTaskNotFound = errors.New("Task not found")
