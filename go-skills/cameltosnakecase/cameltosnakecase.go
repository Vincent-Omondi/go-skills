package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}


func CamelToSnakeCase(s string) string{
	if s[len(s)-1] >= 'a' && s[len(s)-1] <= 'z'{
		for i :=  range s{
			if i > 0 && s[i] >= 'A' && s[i] <= 'Z'{
				if s[i-1] != '_'{
					s = s[:i] + "_" + s[i:]
				}
			}
		}	
	}
	return s
}
