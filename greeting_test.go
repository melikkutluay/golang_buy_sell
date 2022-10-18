package main

import (
	"fmt"
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Hello %s", name)
	result := sayHello(name)

	if result != expected {
		t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, got -> %v", name, expected, result)
	}
}
