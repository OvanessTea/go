package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{}

type OrderService struct {
	logger Logger
}

func main() {
	consoleService := NewOrderService(ConsoleLogger{})
	consoleService.CreateOrder("hello!")

	fileService := NewOrderService(FileLogger{})
	fileService.CreateOrder("hello!")
}

func (FileLogger) Log(message string) {
	fmt.Println("FILE:", message)
}

func (ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

func NewOrderService(logger Logger) OrderService {
	return OrderService{
		logger: logger,
	}
}

func (s OrderService) CreateOrder(message string) {
	s.logger.Log(message)
}

// Ответ на вопрос:
// В текущей реализации Сервису не важно, как именно реализовано логирование.
// Ему важно лишь то, что при создании сервиса указывается структура, удовлетворяющая интерфейс Logger
// Это позволяет нам ,не меняя существующий код Сервиса, добавлять новые реализации
