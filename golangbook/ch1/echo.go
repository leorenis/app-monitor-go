// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// SliceRagesTesteds is
func SliceRagesTesteds() {
	s := []string{"Abobora", "tomate", "manga", "Melon"}
	fmt.Println(s[0:3])
}

// Echo1 is
func Echo1() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // ineficient way. Bad
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%dns eleapsed \n", time.Since(start).Nanoseconds())
}

// Echo2 is
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg // ineficient way. Bad
		sep = " "
	}
	fmt.Println(s)
}

// ExerciseEcho2 is
func ExerciseEcho2() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		sep = " "
		s += sep + strconv.Itoa(index) + sep + arg // ineficient way. Bad
	}
	fmt.Println(s)
}

// Echo3 is
func Echo3() {
	started := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " ")) // Eficient Way :) GOOD
	//fmt.Println(os.Args[1:])                    // Println way
	//fmt.Println(os.Args[0:])                    // Exercise 1.1 | P. 31
	fmt.Printf("%dns eleapsed \n", time.Since(started).Nanoseconds())
}
