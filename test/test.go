package test

type FiberTestModel[T any] struct {
	description  string
	requestBody  *T
	expectStatus int
}
