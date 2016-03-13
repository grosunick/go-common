package core

import "fmt"

// Структура базовой ошибки
type Error struct {
	Code    int
	Message string
}

// Выводит информацию об ошибке
func (this *Error) Error() string {
	return fmt.Sprintf("Code - %d, Message - %s", this.Code, this.Message)
}
