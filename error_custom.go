package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func main() {

}
