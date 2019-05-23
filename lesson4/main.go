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
)

var (
	db = map[string]*[]Student{}
	arrayOfStudents []Student
	students []Student
	y int

)


func main() {
	var login string
	var password string
	var logined bool
	//db := map[string]*[]Student{}
	var exit bool

	readedPassHash := readPasswordHash()

	for !logined {
		fmt.Println("Enter username: ")
		fmt.Scan(&login)
		fmt.Println("Enter pass:")
		fmt.Scan(&password)
		if login == "root" && checkHash(password, readedPassHash) {
			logined = true
			fmt.Println("Hello, root")
		} else {
			fmt.Println("Invalid credentials")
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
				fmt.Println("command now recognize")
			} else {
				tableName := strings.TrimSpace(commandStruct[3])

				if db[tableName] == nil {
					fmt.Println("table not exits")
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
								}
							}
						} else if len(commandStruct) == 4 {
							all()
						} else {
							fmt.Println("command not recognize")
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
						fmt.Println("Wrong column name(", columns, ")")
						break
					}

				}
				break
			}
		case Create:
			tableName := strings.TrimSpace(commandStruct[1])
			emptySlice := []Student{}
			db[tableName] = &emptySlice
			fmt.Println("table created: " + tableName)
			break

		case Insert:
			if commandStruct[1] == "into" {
				tableName := strings.TrimSpace(commandStruct[2])
				if db[tableName] == nil {
					fmt.Println("table not exits")
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
				fmt.Println("command not recognize")
			}
			break
		case Update:
			if len(commandStruct) < 4 {
				fmt.Println("command not recognize")
			}else {
				x,_ := strconv.Atoi(commandStruct[1])
				tableName := strings.TrimSpace(commandStruct[3])
				if db[tableName] == nil {
					fmt.Println("table not exits")
				} else {
					for i, row := range arrayOfStudents {
						if x == row.ID {
							fmt.Printf("\n|  ID|               Fname|Age|Average |\n")
							fmt.Println("------------------------------------")
							fmt.Printf("|%4d|%20s|%3d|%2.2f|\n", row.ID, row.Fname, row.Age, row.Average)
							students = append(students[:i], students[i+1:]...)
							emptySlice := new(Student)
							emptySlice.Insert(x)
							students = append(students, *emptySlice)
							db[tableName] = &students
							arrayOfStudents = *db[tableName]
							all()
						}
					}
				}
			}

			break
		default:
			if len(commandStruct) > 1 {
				fmt.Println("command not recognize")
			}
			break
		}
	}

}

func all()  {

	fmt.Printf("\n|  ID|               Fname|Age|Average |\n")
	fmt.Println("------------------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|%3d|%2.2f|\n", row.ID, row.Fname, row.Age, row.Average)
	}

	fmt.Println("rows returned: ",len(arrayOfStudents))
}

func age()  {

	fmt.Printf("\n|  ID|               Fname|Age|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
	}

	fmt.Println("rows returned: ", len(arrayOfStudents))
}

func average()  {

	fmt.Printf("\n|  ID|               Fname|Average|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|%2.2f|\n", row.ID, row.Fname, row.Average)
	}

	fmt.Println("rows returned: ", len(arrayOfStudents))
}

func fname()  {

	fmt.Printf("\n|  ID|               Fname|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		fmt.Printf("|%4d|%20s|\n", row.ID, row.Fname)
	}

	fmt.Println("rows returned: ", len(arrayOfStudents))
}

func less()  {
	fmt.Printf("\n|  ID|               Fname|Age|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		if row.Age < y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
		}
	}
}

func more()  {
	fmt.Printf("\n|  ID|               Fname|Age|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		if row.Age > y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
		}
	}
}

func equal()  {
	fmt.Printf("\n|  ID|               Fname|Age|\n")
	fmt.Println("------------------------------")

	for _, row := range arrayOfStudents {
		if row.Age == y {
			fmt.Printf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
		}
	}
}

func (s *Student) Insert(x int)  Student{
	age := 0
	var fname string= ""

	for fname == ""{
		fmt.Println("Fname: ")
		fmt.Scan(&fname)
	}
	for age == 0{
		fmt.Println("Age: ")
		fmt.Scan(&age)
	}

	s.ID = x
	s.Fname = fname
	s.Age = age
	return *s
}
