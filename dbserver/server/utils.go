package main

import (
	"strings"
	"strconv"
	"fmt"
)

func getError(err string) string {
	s := "\n" + err + "\n"
	s += "---------------------;"
	return s
}

func handleCmd(cmd string) string {
	_cmd := strings.TrimSpace(cmd)
	commandStruct := strings.Split(_cmd, " ")

	switch commandStruct[0] {
	case Select:

		if len(commandStruct) < 4 {
			return getError("command not recognize")
		} else {
			tableName := strings.TrimSpace(commandStruct[3])
			arrayOfStudents = *db[tableName]

			if db[tableName] == nil {
				return getError("table not exist")
			} else {
				columns := commandStruct[1]
				switch columns {

				case "*":
					if len(commandStruct) == 8 {
						if commandStruct[4] == "where" && commandStruct[5] == "age" {
							y, _ = strconv.Atoi(commandStruct[7])

							if commandStruct[6] == "<" {
								return less()
							} else if commandStruct[6] == ">" {
								return more()
							} else if commandStruct[6] == "==" {
								return equal()
							} else {
								s := "Enter character correctly(<,>,==)" + "\n"
								s += "---------------------------------" + "\n"
								return s + ";"
							}
						}
					} else if len(commandStruct) == 4 {
						return all()
					} else {
						return getError("command not recognize")
					}
					break

				case "age", "Age":
					return age()
					break

			 	case "average", "Average":
					return average()
					break

				 case "Fname", "fname":
					return fname()
					break

				default:
					s := "\nWrong column name("+columns+ ")"
					s += "---------------------" + strings.Repeat("-", len(columns)) + "\n"
					return s

				}
			}
			break

		}
	case Create:
		tableName := strings.TrimSpace(commandStruct[1])

		if db[tableName] != nil {
			return getError("table already exist")
		}else {
			emptySlice := []Student{}
			db[tableName] = &emptySlice
			return  tableName + " created;"
		}
	case Insert:
		if len(commandStruct) < 4 {
			return getError("command not recognize")
		}

		if commandStruct[1] == "into" {
			tableName := strings.TrimSpace(commandStruct[2])
			if db[tableName] == nil {
				return getError("table not exist")
			} else {
				x := 1
				students := *db[tableName]

				for _, row := range students {
					if x != row.ID {
						break
					}else {
						x += 1
					}
				}

				fields := strings.Split(commandStruct[3], ",")
				age, _ := strconv.Atoi(fields[1])
				average,_ := strconv.ParseFloat(fields[2],32)
				exp, _ := strconv.Atoi(fields[3])
				emptySlice := &Student{
					ID: x,
					Fname: fields[0],
					Age: age,
					IsStudent: false,
					IsWorker: false,
					IsTeacher: false,
					Average: float32(average),
					Experience: exp,
				}

				fmt.Printf("%v\n", *emptySlice)


				students = append(students, *emptySlice)
				db[tableName] = &students
				arrayOfStudents = *db[tableName]
				return all()
			}
		} else {
			return getError("command not recognize")
		}
		break

	case Update:
		if len(commandStruct) < 3 {
			return getError("command not recognize")
		}else {
			if len(commandStruct) < 5 && commandStruct[2] != "from" {
				return getError("command not recognize")

			} else {
				tableName := strings.TrimSpace(commandStruct[3])
				students := *db[tableName]

				x, _ := strconv.Atoi(commandStruct[1])

				for i, row := range students {
					if x == row.ID {
						students := *db[tableName]
						students = append(students[:i], students[i+1:]...)

						fields := strings.Split(commandStruct[4], ",")
						age, _ := strconv.Atoi(fields[1])
						average,_ := strconv.ParseFloat(fields[2],32)
						exp, _ := strconv.Atoi(fields[3])
						emptySlice := &Student{
							ID:         x,
							Fname:      fields[0],
							Age:        age,
							IsStudent:  false,
							IsWorker:   false,
							IsTeacher:  false,
							Average:    float32(average),
							Experience: exp,
						}
							students = append(students, *emptySlice)
							db[tableName] = &students
							arrayOfStudents = *db[tableName]
							return all()

					}
				}
			}
		}
		break

	case Delete:
		if len(commandStruct) < 4 && commandStruct[2] != "from"{
			return getError("command not recognize")
		}else {
			return deleteIt(commandStruct)
		}
		break
	default:
		if len(commandStruct) > 1 {
			return getError("command not recognize")
		}

		return "NAN;"
	}

	return "Nan;"
}


func all() string {
	s := "\n" + strings.Repeat("-", len(outputArr[0])) + "\n"
	s += outputArr[0] + "\n"
	s += strings.Repeat("-", len(outputArr[0])) + "\n"

	for _, row := range arrayOfStudents {
		s += fmt.Sprintf("|%4d|%20s|%3d| %9v | %8v | %9v |  %1.2f | %9d |\n", row.ID, row.Fname, row.Age, row.IsStudent, row.IsWorker, row.IsTeacher, row.Average, row.Experience)
		s+= strings.Repeat("-", len(outputArr[0])) + "\n"
	}

	s += "rows returned: "+ fmt.Sprint(len(arrayOfStudents))+ "\n;"
	return s
}

func age() string  {
	s := "\n" + strings.Repeat("-", len(outputArr[1])) + "\n"
	s += outputArr[1] + "\n"
	s += strings.Repeat("-", len(outputArr[1])) + "\n"

	for _, row := range arrayOfStudents {
		s += fmt.Sprintf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
		s+= strings.Repeat("-", len(outputArr[1])) + "\n;"
	}

	s += "rows returned: "+ string(len(arrayOfStudents))+ ";"
	return s
}

func less() string {
	s := "\n" + strings.Repeat("-", len(outputArr[1])) + "\n"
	s +=  outputArr[1] + "\n"
	s += strings.Repeat("-", len(outputArr[1])) + "\n"

	for _, row := range arrayOfStudents {
		if row.Age < y {
			s += fmt.Sprintf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			s += strings.Repeat("-", len(outputArr[1]))
		}
	}
	s+=";"
	return s
}

func more() string  {
	s := "\n" + strings.Repeat("-", len(outputArr[1])) + "\n"
	s += outputArr[1] + "\n"
	s += strings.Repeat("-", len(outputArr[1])) + "\n"

	for _, row := range arrayOfStudents {
		if row.Age > y {
			s += fmt.Sprintf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			s += strings.Repeat("-", len(outputArr[1])) + "\n"
		}
	}
	s+=";"
	return s
}

func equal() string  {
	s := "\n" + strings.Repeat("-", len(outputArr[1])) + "\n"
	s += outputArr[1] + "\n"
	s += strings.Repeat("-", len(outputArr[1]))	+ "\n"

	for _, row := range arrayOfStudents {
		if row.Age == y {
			s += fmt.Sprintf("|%4d|%20s|%3d|\n", row.ID, row.Fname, row.Age)
			s += strings.Repeat("-", len(outputArr[1]))+ "\n"
		}
	}
	s+=";"
	return s
}

func average()  string{
	s := "\n-----------------------------------\n"
	s += "|  ID|               Fname|Average|\n"
	s += "-----------------------------------\n"

	for _, row := range arrayOfStudents {
		s += fmt.Sprintf("|%4d|%20s|  %1.2f |\n", row.ID, row.Fname, row.Average)
		s += "-----------------------------------\n"
	}

	s += "rows returned: " + string(len(arrayOfStudents)) + "\n;"

	return s
}

func fname()  string{

	s := "\n---------------------------\n"
	s += "|  ID|               Fname|\n"
	s += "---------------------------\n"

	for _, row := range arrayOfStudents {
		s += fmt.Sprintf("|%4d|%20s|\n", row.ID, row.Fname)
		s += "---------------------------\n"
	}

	s += "rows returned: " + string(len(arrayOfStudents)) + "\n;"

	return s
}

func deleteIt(commandStruct []string)  string{
	x,_ := strconv.Atoi(commandStruct[1])
	tableName := strings.TrimSpace(commandStruct[3])

	if db[tableName] == nil {
		return getError("table not exits")
	} else {
		arrayOfStudents = *db[tableName]

		for i, row := range arrayOfStudents {
			if x == row.ID {
				fmt.Println("\n" + strings.Repeat("-", len(outputArr[0])))
				fmt.Println(outputArr[0])
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				fmt.Printf("|%4d|%20s|%3d| %9v | %8v | %9v |  %1.2f | %9d |\n", row.ID, row.Fname, row.Age, row.IsStudent, row.IsWorker, row.IsTeacher, row.Average, row.Experience)
				fmt.Println(strings.Repeat("-", len(outputArr[0])))
				arrayOfStudents = append(arrayOfStudents[:i], arrayOfStudents[i+1:]...)
				db[tableName] = &arrayOfStudents
				return all()
			}
		}
	}
	return ""
}
