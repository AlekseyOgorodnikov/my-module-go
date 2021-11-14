package myModule

import "fmt"

func Version(s string) string {
	fmt.Println("Version 1.0.0")
	message := s + ", Hello world!"
	return message
}
