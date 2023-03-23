package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type Student struct {
	FirstName, LastName, University                string
	Test1Score, Test2Score, Test3Score, Test4Score int
}

type StudentStat struct {
	Student
	FinalScore float32
	Grade      Grade
}

func ParseCSV(filePath string) ([]Student, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []Student
	reader := csv.NewReader(file)
	reader.Read() // Skip header row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		test1, _ := strconv.Atoi(record[3])
		test2, _ := strconv.Atoi(record[4])
		test3, _ := strconv.Atoi(record[5])
		test4, _ := strconv.Atoi(record[6])

		students = append(students, Student{
			FirstName:  record[0],
			LastName:   record[1],
			University: record[2],
			Test1Score: test1,
			Test2Score: test2,
			Test3Score: test3,
			Test4Score: test4,
		})
	}
	return students, nil
}

func CalculateGrade(students []Student) []StudentStat {
	var gradedStudents []StudentStat
	for _, s := range students {
		finalScore := float32(s.Test1Score+s.Test2Score+s.Test3Score+s.Test4Score) / 4
		grade := F
		switch {
		case finalScore >= 70:
			grade = A
		case finalScore >= 50:
			grade = B
		case finalScore >= 35:
			grade = C
		}

		gradedStudents = append(gradedStudents, StudentStat{
			Student:    s,
			FinalScore: finalScore,
			Grade:      grade,
		})
	}
	return gradedStudents
}

func FindOverallTopper(gradedStudents []StudentStat) StudentStat {
	topper := gradedStudents[0]
	for _, gs := range gradedStudents {
		if gs.FinalScore > topper.FinalScore {
			topper = gs
		}
	}
	return topper
}

func FindTopperPerUniversity(gs []StudentStat) map[string]StudentStat {
	toppers := make(map[string]StudentStat)
	for _, studentStat := range gs {
		university := studentStat.University
		if _, ok := toppers[university]; !ok || studentStat.FinalScore > toppers[university].FinalScore {
			toppers[university] = studentStat
		}
	}
	return toppers
}
