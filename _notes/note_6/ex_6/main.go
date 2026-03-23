package main

import (
	"fmt"
	"go/build/constraint"
	"strings"
)

func main() {
	names := []string{"bob", "jane", "tom", "karl", "ann"}

	result := Contains("tom", names)
	fmt.Println(result)

	res := Map(names, Upper)
	fmt.Println(res)
}

func Map[F, T any](s []F, f func(F) T) []T {
	result := make([]T, len(s))

	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

func Contains[S comparable](s S, ss []S) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func Upper(str string) string {
	constraint.
	return strings.ToUpper(str)
}
