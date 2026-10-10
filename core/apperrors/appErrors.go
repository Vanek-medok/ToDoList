package apperrors

import "errors"

var ErrTaskNotFound = errors.New("Задача не найдена")
var ErrInvalidTaskID = errors.New("неккоректный ID задачи")
