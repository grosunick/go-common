package worker

// Base task handler interface
type IHandler interface {
	// Runs task excecution
	Run() (interface{}, error)
}

// Многопоточный обработчик заданий
type IMultyThreadWorker interface {
	// Добавляет задание в очередь
	Add(cmd IHandler)
	// Запускает обработку заданий
	Run()
	// Уведомляет обработчик о завершении выполнения задания
	notify(error error)
}
