package main

import (
	"fmt"
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
	s.Average = total/ float64(len(s.Grades))
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
	var name string 
	for true{
		fmt.Println("Enter Your Name Please")
		_, err := fmt.Scanln( &name)
		if err != nil{
			fmt.Println("Wrong Input. Please Enter Again")
		} else 
			{break }
		}
		return name
}

func receiveCourse() map[string]float64{
	grades := make(map[string]float64)
	courses := 0
	for true{
		fmt.Println("How many courses are you taking")
		_, err := fmt.Scanln(&courses)
		if err != nil || courses < 1{
			fmt.Println("Wrong Input Please Try Again.")
		} else {break }
	}
	for true {
		course := ""
		value := 0.0
		fmt.Println("Enter Course name ")
		_, err := fmt.Scanln( &course)

		if err != nil{
			fmt.Println("Wrong INput Please Try Again.")
		} else { 
			fmt.Printf("Enter Points you got for %s\n", course)
			_, err := fmt.Scanln( &value)
				if err != nil || value < 0 || value > 100{
					fmt.Println("Wrong INput Please Try Again. 3")
				} else {
					grades[course] = value
					courses --
				}
			}
			if courses == 0{
					break
				}
		} 

		return grades 
	}
