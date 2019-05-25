package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	Select = "select"
	Create = "create"
	Insert = "insert"
	Update = "update"
	Delete = "delete"
	Exit   = "exit"
)

var (
	db = map[string]*[]Student{}
	arrayOfStudents []Student
	students []Student
	y int
	outputArr =[]string{"|  ID|               Fname|Age|  IsStudent|  IsWorker|  IsTeacher|Average| Experience|",
		"|  ID|               Fname|Age|"}

)


func main() {
	var login string
	var password string
	var logined bool
	var exit bool

	readedPassHash := readPasswordHash()

	for !logined {
		fmt.Print("Enter username: ")
		fmt.Scan(&login)
		fmt.Print("Enter pass: ")
		fmt.Scan(&password)
		if login == "root" && checkHash(password, readedPassHash) {
			logined = true
			fmt.Println("\nHello, root\n")
		} else {
			fmt.Println("\nInvalid credentials")
			fmt.Println("-------------------\n")
		}
	}

	for !exit {
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		commandStruct := strings.Split(command, " ")

		switch commandStruct[0] {
		case Select:
			// select * from tablename
			if len(commandStruct) < 4 {
				errCommand()
			} else {
				tableName := strings.TrimSpace(commandStruct[3])

				if db[tableName] == nil {
					tableNotExist()
				} else {
					columns := commandStruct[1]
					arrayOfStudents = *db[tableName]
					switch columns {

					case "*":
						if len(commandStruct) == 8 {
							if commandStruct[4] == "where" && commandStruct[5] == "age" {
								y, _ = strconv.Atoi(commandStruct[7])

								if commandStruct[6] == "<" {
									less()
								} else if commandStruct[6] == ">" {
									more()
								} else if commandStruct[6] == "==" {
									equal()
								} else {
									fmt.Println("Enter character correctly(<,>,==)")
									fmt.Println("---------------------------------")
								}
							}
						} else if len(commandStruct) == 4 {
							all()
						} else {
							errCommand()
						}
						break

					case "age", "Age":
						age()
						break

					case "average", "Average":
						average()
						break

					case "Fname", "fname":
						fname()
						break

					default:
						fmt.Println("\nWrong column name(", columns, ")")
						fmt.Println("---------------------" + strings.Repeat("-", len(columns)) + "\n")
						break
					}
				}
				break

			}
		case Create:
			tableName := strings.TrimSpace(commandStruct[1])
			emptySlice := []Student{}
			db[tableName] = &emptySlice
			fmt.Println( "\n" + tableName + " created \n")
			break

		case Insert:
			if commandStruct[1] == "into" {
				tableName := strings.TrimSpace(commandStruct[2])
				if db[tableName] == nil {
					tableNotExist()
				} else {
					x := 0
					for _, row := range arrayOfStudents {
						if x < row.ID {
							x = row.ID
						}
					}
					x += 1
					emptySlice := new(Student)
					emptySlice.Insert(x)
					students = append(students, *emptySlice)
					db[tableName] = &students
					arrayOfStudents = *db[tableName]
					all()
				}
			} else {
				errCommand()
			}
			break
		case Update:
			if len(commandStruct) < 3 {
				errCommand()
			}else {
				if len(commandStruct) < 4 && commandStruct[2] != "from" {
					errCommand()
				} else {
					update(commandStruct)
				}
			}
			break
		case Delete:
			if len(commandStruct) < 4 && commandStruct[2] != "from"{
				errCommand()
			}else {
				deleteIt(commandStruct)
			}
			break

		case Exit:
			os.Exit(1)

		default:
			if len(commandStruct) > 1 {
				errCommand()
			}
			break
		}
	}
}

func all()  {
	fmt.Println("\n" + strings.Repeat("-", len(outputArr[0])))
	fmt.Println(outputArr[0])
	fmt.Println(strings.Repeat("-", len(outputArr[0])))

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|%3d| %9v | %8v | %9v |  %1.2f | %9d |\n", row.ID, row.Fname, row.Age, row.IsStudent, row.IsWorker, row.IsTeacher, row.Average, row.Experience)
		fmt.Println(strings.Repeat("-", len(outputArr[0])))
	}

	fmt.Println("rows returned: ",len(arrayOfStudents), "\n")
}

func age()  {
	fmt.Println("\n" + strings.Repeat("-", len(outputArr[1])))
	fmt.Println(outputArr[1])
	fmt.Println(strings.Repeat("-", len(outputArr[1])))

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
		fmt.Println(strings.Repeat("-", len(outputArr[1])))
	}

	fmt.Println("rows returned: ", len(arrayOfStudents), "\n")
}

func average()  {
	fmt.Println("\n-----------------------------------")
	fmt.Println("|  ID|               Fname|Average|")
	fmt.Println("-----------------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|  %1.2f |\n", row.ID, row.Fname, row.Average)
		fmt.Println("-----------------------------------")
	}

	fmt.Println("rows returned: ", len(arrayOfStudents), "\n")
}

func fname()  {
	fmt.Println("\n---------------------------")
	fmt.Println("|  ID|               Fname|")
	fmt.Println("---------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|\n", row.ID, row.Fname)
		fmt.Println("---------------------------")
	}

	fmt.Println("rows returned: ", len(arrayOfStudents), "\n")
}

func less()  {
	fmt.Println("\n" + strings.Repeat("-", len(outputArr[1])))
	fmt.Println(outputArr[1])
	fmt.Println(strings.Repeat("-", len(outputArr[1])))

	for _, row := range arrayOfStudents {
		if row.Age < y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			fmt.Println(strings.Repeat("-", len(outputArr[1])))
		}
	}
}

func more()  {
	fmt.Println("\n" + strings.Repeat("-", len(outputArr[1])))
	fmt.Println(outputArr[1])
	fmt.Println(strings.Repeat("-", len(outputArr[1])))

	for _, row := range arrayOfStudents {
		if row.Age > y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			fmt.Println(strings.Repeat("-", len(outputArr[1])))
		}
	}
}

func equal()  {
	fmt.Println("\n" + strings.Repeat("-", len(outputArr[1])))
	fmt.Println(outputArr[1])
	fmt.Println(strings.Repeat("-", len(outputArr[1])))

	for _, row := range arrayOfStudents {
		if row.Age == y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			fmt.Println(strings.Repeat("-", len(outputArr[1])))
		}
	}
}

func (s *Student) Insert(x int)  Student{
	age := 0
	fname := ""
	var who bool
	isSomeone := ""

		for fname == ""{
			fmt.Print("\nFname: ")
			fmt.Scan(&fname)
		}
		for age == 0{
			fmt.Print("Age: ")
			fmt.Scan(&age)
		}

		for !who {
			fmt.Print("Who is this person: ")
			fmt.Scan(&isSomeone)
			if isSomeone == "student"{
				s.IsStudent = true
				break
			}else if isSomeone == "teacher" {
				s.IsTeacher = true
				break
			}else if isSomeone == "worker" {
				s.IsWorker = true
				break
			}else {
				fmt.Println("\nenter right category (student|teache|worker)")
				fmt.Println("--------------------------------------------\n")
			}
		}
		if s.IsStudent == true {

			var average float32
			for average <= 0 {
				fmt.Print("Enter average: ")
				fmt.Scan(&average)
			}
			s.Average = average
		}else if s.IsTeacher == true || s.IsWorker == true{
			var experience int
			for experience <= 0 {
				fmt.Print("Enter experience: ")
				fmt.Scan(&experience)
			}
			s.Experience = experience
		}

		s.ID = x
		s.Fname = fname
		s.Age = age
		return *s

}

func (s *Student)UpdateIt() Student{
	age := 0
	fname := ""
	var who bool
	isSomeone := ""
	choise :=""

	fmt.Println("Choose what you want to Update")
	fmt.Scanln(&choise)

	switch choise {
	case "fname":
		for fname == ""{
			fmt.Print("\nFname: ")
			fmt.Scan(&fname)
		}
		s.Fname = fname
		break
	case "age":
		for age == 0{
			fmt.Print("Age: ")
			fmt.Scan(&age)
		}
		s.Age = age
		break
	case "who":
		for !who {
			fmt.Print("Who is this person: ")
			fmt.Scan(&isSomeone)
			if isSomeone == "student"{
				s.IsStudent = true
				break
			}else if isSomeone == "teacher" {
				s.IsTeacher = true
				break
			}else if isSomeone == "worker" {
				s.IsWorker = true
				break
			}else {
				fmt.Println("\nenter right category (student|teache|worker)")
				fmt.Println("--------------------------------------------\n")
			}
		}
		if s.IsStudent == true {

			var average float32
			for average <= 0 {
				fmt.Print("Enter average: ")
				fmt.Scan(&average)
			}
			s.Average = average
		}else if s.IsTeacher == true || s.IsWorker == true{
			var experience int
			for experience <= 0 {
				fmt.Print("Enter experience: ")
				fmt.Scan(&experience)
			}
			s.Experience = experience
		}
		break

	case "all":
		for fname == ""{
			fmt.Print("\nFname: ")
			fmt.Scan(&fname)
		}
		for age == 0{
			fmt.Print("Age: ")
			fmt.Scan(&age)
		}

		for !who {
			fmt.Print("Who is this person: ")
			fmt.Scan(&isSomeone)
			if isSomeone == "student"{
				s.IsStudent = true
				break
			}else if isSomeone == "teacher" {
				s.IsTeacher = true
				break
			}else if isSomeone == "worker" {
				s.IsWorker = true
				break
			}else {
				fmt.Println("\nenter right category (student|teache|worker)")
				fmt.Println("--------------------------------------------\n")
			}
		}
		if s.IsStudent == true {

			var average float32
			for average <= 0 {
				fmt.Print("Enter average: ")
				fmt.Scan(&average)
			}
			s.Average = average
		}else if s.IsTeacher == true || s.IsWorker == true{
			var experience int
			for experience <= 0 {
				fmt.Print("Enter experience: ")
				fmt.Scan(&experience)
			}
			s.Experience = experience
		}
		s.Fname = fname
		s.Age = age
		break
	}
	return *s
}

func update(commandStruct []string)  {
	x, _ := strconv.Atoi(commandStruct[1])
	tableName := strings.TrimSpace(commandStruct[3])
	if db[tableName] == nil {
		tableNotExist()
	} else {
		for i, row := range arrayOfStudents {
			if x == row.ID {
				fmt.Println("\n" + strings.Repeat("-", len(outputArr[0])))
				fmt.Println(outputArr[0])
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				fmt.Printf("|%4d|%20s|%3d| %9v | %8v | %9v |  %1.2f | %9d |\n", row.ID, row.Fname, row.Age, row.IsStudent, row.IsWorker, row.IsTeacher, row.Average, row.Experience)
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				students = append(students[:i], students[i+1:]...)
				row.UpdateIt()

				students = append(students, row)
				db[tableName] = &students
				arrayOfStudents = *db[tableName]
				all()
			}
		}
	}
}

func deleteIt(commandStruct []string)  {
	x,_ := strconv.Atoi(commandStruct[1])
	tableName := strings.TrimSpace(commandStruct[3])
	if db[tableName] == nil {
		tableNotExist()
	} else {
		for i, row := range arrayOfStudents {
			if x == row.ID {
				fmt.Println("\n" + strings.Repeat("-", len(outputArr[0])))
				fmt.Println(outputArr[0])
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				fmt.Printf("|%4d|%20s|%3d| %9v | %8v | %9v |  %1.2f | %9d |\n", row.ID, row.Fname, row.Age, row.IsStudent, row.IsWorker, row.IsTeacher, row.Average, row.Experience)
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				students = append(students[:i], students[i+1:]...)
				db[tableName] = &students
				arrayOfStudents = *db[tableName]
				all()
			}
		}
	}
}

func errCommand()  {
	fmt.Println("\ncommand not recognize")
	fmt.Println("---------------------\n")
}

func tableNotExist()  {
	fmt.Println("\ntable not exits")
	fmt.Println("---------------\n")
}