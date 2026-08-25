package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Service struct {
	speaker Speaker
}

type Robot struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func (Robot) Speak() {
	fmt.Println("Beep!")
}

// struct Dog has a method Speak
// So it comes Speaker interface automatically

func MakeSpeak(s Speaker) {
	s.Speak()
}

// Question 1
// MakeSpeak принимает только переменные, имеющие интерфейс Speaker
// В нашей реализации интерфейс Speaker имеет метод Speak()
// По правилам Go: любая структура, имеющая метод Speak() сразу становится интерфейсом Speaker
// d -> Dog{} -> Speaker{} -> MakeSpeak()

// Question 2
// Функции без разницы, какой struct в неё передается.
// Ей важно, чтобы struct, который был передан, имет интерфейс Speaker
// В нашем случае и Dog, и Cat имеют метод Speak()
// , что удовлетворяет требованиям для интерфейса Speaker

// Question 3
// Указывая speaker Speaker в структуре сервиса, мы указываем
// , что Сервис может хранить в себе любую структуру, удовлетворяющую Speaker
// В нашем случае, Service может хранить в себе и Dog{}, и Cat{}

func NewService(s Speaker) Service {
	return Service{
		speaker: s,
	}
}

// Question 4
// Правильный ответ: C) Service зависит от Speaker, а Dog просто предоставляет нужное поведение

// Question 5
// Нам не надо менять Service при добавлении новых структур. По сути, сервису без разницы,
// какие структуры в него передаются. Даже если они были добавлены постфактум.
// Сервису важно только одно: структура, которая в него передается, удовлетворяет требования интерфейса

// Lil' practice
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

func (f FileLogger) Log(message string) {
	fmt.Println("File:", message)
}

type LogService struct {
	logger Logger
}

func main() {
	// d := Dog{Name: "Jack"}
	// MakeSpeak(d)

	// c := Cat{Name: "Mira"}
	// MakeSpeak(c)

	// service := NewService(d)

	// _ = service

	// robot := Robot{}

	// service = NewService(robot)
	console := ConsoleLogger{}
	service := NewLogService(console)

	service.MakeLog("Hello!")

	file := FileLogger{}
	service = NewLogService(file)

	service.MakeLog("Hello!")
}

func NewLogService(l Logger) LogService {
	return LogService{
		logger: l,
	}
}

func (s LogService) MakeLog(message string) {
	s.logger.Log(message)
}
