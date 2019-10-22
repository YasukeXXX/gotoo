package gotoo_test

import (
	"fmt"
	"github.com/YasukeXXX/gotoo"
)

func ExampleCopyFields() {
	type Source struct {
		Name string
		Email string
		Password string
	}

	type Target struct {
		Name string
		Email string
		Age int
	}

	t, err := gotoo.CopyFields(Source{"name", "email@example.com", "password"}, Target{})
	if err != nil {
		fmt.Print(err)
		return
	}
	if target, ok := t.(*Target); ok {
		fmt.Printf("%#v", target)
	}
	return
}
