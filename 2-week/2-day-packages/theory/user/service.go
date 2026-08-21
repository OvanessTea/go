package user

// code of the same package could have accessions
// for unexported entities from another code

func changeEmail(u *User, email string) {
	u.email = email
}
