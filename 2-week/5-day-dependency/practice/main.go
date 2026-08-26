package main

import "fmt"

// Ситуация 1
// Я считаю, что интерфейс здесь нужен, потому что нет гарантий,
// что не появится новая структура, из-за чего придется переписывать логику зависимостей
// Если не использовать interface Logger, тогда придется жестко прописывать
// , что OrderService зависит от структуры ConsoleLogger
// Но если 100% гарантии, что больше структур не появился, что можно обойтись без интерфейса.
// Но мне такое не нравится.

// Ситуация 2
// Здесь интерфейс 100% нужен, потому что мы уже имеем две структуры + 3 в потенциале.
// Не за чем делать сложную зависимость у OrderService. Гораздо проще и правильнее вынести зависимость
// в отдельный интерфейс и сделать:
// OrderService {payment PaymentProcessor}, где PaymentProcessor - interface с методом Pay(amount int)
// , чтобы наши струтктуры удовлетворяли интерфейс.
// Получается зависимость: OrderService -> PaymentProcessor <- Stripe|PayPal|...
// конкретные реализации фукнции Pay() каждая структура описывает для самой себя.
// Ни интерфейсу, ни сервису не интересна конкретная реализация.
// type OrderService struct {payment PaymentProcessor}
// type PaymentProcessor interface {Pay(amount int)}
// func (s Stripe) Pay(amount int) {...}
// func (p PayPal) Pay(amount int) {...}
// func NewOrderService(p PaymentProcessor) OrderService {...}

// Ситуация 3
// Я считаю, что в данной реализации интерфейс излишний.
// Я не вижу никакой нужды в использовании интерфейса здесь.
// Структура получает метод и может его использовать.

// Ситуация 4
type OrderService struct {
	payment PaymentProcessor
}

type PaymentProcessor interface {
	Pay(amount int)
}

type StripeProcessor struct{}

func (StripeProcessor) Pay(amount int) {
	fmt.Println("Pay with Stripe:", amount)
}

type PayPalProcessor struct{}

func (PayPalProcessor) Pay(amount int) {
	fmt.Println("Pay with PayPal:", amount)
}

func NewOrderService(p PaymentProcessor) OrderService {
	return OrderService{
		payment: p,
	}
}

func main() {
	stripe := StripeProcessor{}

	service := NewOrderService(stripe)
	service.CreateOrder()
}

func (s OrderService) CreateOrder() {
	s.payment.Pay(100)
}
