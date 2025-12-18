package pkg

func Foo() string {
	return "foo" + baz()
}

func Bar() string {
	return "bar"
}

func baz() string {
	return "baz"
}
