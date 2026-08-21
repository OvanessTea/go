package main

func main() {

}

// -----------------------------
// file of same area of responsibility should be placed in same package
// -----------------------------
// otherwise files with different respons should have self pkgs
// user/
//     user.go
// email/
//     email.go
// payment/
//     payment.go

// user       → users
// email      → email
// payment    → payments
// -----------------------------
// pkg has to have understandable respons
// user
// product
// order
// payment
// notification
// -----------------------------
// BAD PRACTICE
// utils/
//     string.go
//     date.go
//     http.go
//     json.go
//     password.go
//     user.go

// utils.DoSomething() - Misundestanding
// other examples:
// utils.HashPassword()
// utils.FormatDate()
// utils.ParseJSON()
// utils.ValidateEmail()
// utils.GenerateToken()
// These are usefull functions but they have different areas of respons
// auth/
//     password.go
//     token.go

// email/
//     validation.go

// json/
//     ...
// -----------------------------
// No need to create pkg for single file
// project/
// ├── user/
// │   └── user.go
// ├── name/
// │   └── name.go
// ├── age/
// │   └── age.go
// ├── email/
// │   └── email.go
// All that could be placed in same pkg: user
// -----------------------------
//         encapsulation
//         package user
// ┌─────────────────────────┐
// │                         │
// │ internal implementation │
// │                         │
// │  validateName()         │
// │  validateEmail()        │
// │  normalizeName()        │
// │                         │
// │        ↓                │
// │                         │
// │      NewUser()          │ ← API
// │      GetEmail()         │ ← API
// │                         │
// └───────────┬─────────────┘
//             │
//             ▼
//       external code
// -----------------------------
// import cycle
// user → order
//   ↑     ↓
//   └─────┘
// it's forbidden in Go
// dependencies have to be linear
// user
//   │
//   ▼
// notification
