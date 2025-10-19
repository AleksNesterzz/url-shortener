package validator

type Validator interface {
	IsValid(string) bool
}

//TODO : add validator logic to url validation
