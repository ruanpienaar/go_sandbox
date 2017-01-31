package main

import "fmt" "os"

func main(){
  if len(os.Args) > 1 {
     fmt.Println(HelloWorld(os.Args[1]))
  } else {
     fmt.Println(HelloWorl(""))
  }
}

func HelloWorld(s string) string {

	if len(s) == 0 {
		s = "World"
	}

	return fmt.Sprintf("Hello, %s!", s)
}