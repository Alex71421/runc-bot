package greet

func exclaim(s string) string {
	return s + "!"
}

// Hello возвращает приветствие для name.
func Hello(name string) string {
	return exclaim("Hello, " + name)
}
