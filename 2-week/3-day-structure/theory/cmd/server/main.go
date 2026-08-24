package main

// cmd - usually includes entrypoints of an app
// cmd/
// ├── server/
// │   └── main.go
// ├── worker/
// │   └── main.go
// └── migrate/
//     └── main.go
// server  → launch HTTP API
// worker  → lauch background worker
// migrate → Execute DB migrations
// These are separated programms

// internal - separated programms could use busness logic from internal/ dir
//      internal
//    ↙    ↓     ↘
// api   worker  migrate
// internal code could be used only in permitted parent tree
// external code can't import that package

// unexported - isolation for package layer
// internal - isolation for application layer

// github.com/company/shop/
// ├── cmd/
// │   └── server/
// │       └── main.go
// │
// └── internal/
//     └── user/
//         └── user.go
// We can use internal in main.go
// import "github.com/company/shop/internal/user"
// 'cause it's the same project
// But we can't import that internal in:
// github.com/someone/other-project/
// It's forbidden

// It's useful in Backend 'cause BE usually has a lot of features that
// shouldn't be used as Public API

// Project structure:
// myapp/
// ├── go.mod
// │
// ├── cmd/
// │   └── server/
// │       └── main.go - build app & launch server
// │
// └── internal/
//     ├── user/
//     │   ├── user.go - internal logic of a user
//     │   └── service.go
//     │
//     └── notification/
//         └── notification.go - everything about notifications

// ----------------------------------------------------------------

// Very often the interface determines which package needs behavior.
// User tells:
// type Notifier interface {
//     Notify(message string)
// }
// Notification gives specific realization
// That's why user is not depend on specific type: EmailNotifier

// Dependency inversion
// Bad practice: user -> EmailNotifier
// Better: user -> Notifier <- email
// user determines contract:
// type Notifier interface {
//     Notify(message string)
// }
// And the concrete implementation satisfies this contract.

// We shouldn't use interface for everything
// interface useful when:
// - Overrides the implementation
// - Reduces coupling;
// - Simplifies testing
// - Separates business logic from infrastructure.

func main() {

}
