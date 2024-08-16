package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct{
	Name string 
	Grades map[string]float64
	Average float64
}

func (s *Student) calculateAverage(){
	total := 0.0
	fmt.Println()
	fmt.Println("-----------------------")

	fmt.Println("Name: ", s.Name)
	for subject, grade := range s.Grades {
		total += grade 
		fmt.Printf("%s :%f\n", subject, grade)
	}
	s.Average = total / float64(len(s.Grades))
	fmt.Printf("Average is: %.2f\n", s.Average)
}

func main(){
	fmt.Println("Hello there!")
	name := receiveName()
	courses := receiveCourse()
	student := Student{Name: name, Grades: courses, Average: 0.0}
	student.calculateAverage()
}

func receiveName() string{
	reader := bufio.NewReader(os.Stdin)
	var name string 
	for true {
		fmt.Println("Enter Your Name Please")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("Wrong Input. Please Enter Again")
		} else {
			break
		}
	}
	return name
}

func receiveCourse() map[string]float64{
	reader := bufio.NewReader(os.Stdin)
	grades := make(map[string]float64)
	courses := 0
	for true {
		fmt.Println("How many courses are you taking")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		numCourses, err := strconv.Atoi(input)
		if err != nil || numCourses < 1 {
			fmt.Println("Wrong Input Please Try Again.")
		} else {
			courses = numCourses
			break
		}
	}
	for courses > 0 {
		fmt.Println("Enter Course name ")
		course, _ := reader.ReadString('\n')
		course = strings.TrimSpace(course)

		fmt.Printf("Enter Points you got for %s\n", course)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value < 0 || value > 100 {
			fmt.Println("Wrong Input Please Try Again.")
		} else {
			grades[course] = value
			courses--
		}
	} 

	return grades 
}
