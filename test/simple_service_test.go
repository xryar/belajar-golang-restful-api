package test

import (
	"belajar-golang-retsful-api/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
