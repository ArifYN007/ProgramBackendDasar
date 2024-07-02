package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	students             []model.Student
	studentStudyPrograms map[string]string
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students // TODO: replace this
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if len(id) == 0 || len(name) == 0 {
		return "", fmt.Errorf("ID or Name is undefined")
	}
	// sm = NewInMemoryStudentManager()
	// students := sm.students
	// fmt.Println(students)
	// var nameExist string
	var FoundUntilEnd bool //alangkah baiknya dibuat var saja tidak usah dikasih inisiasi awal
	// fmt.Println("foundultilEND :", FoundUntilEnd)
	for _, student := range sm.students {
		if student.ID == id || student.Name == name { //JIKA NAMA & ID ADA LOOPING BERHENTI
			FoundUntilEnd = true
			break
		} else {
			FoundUntilEnd = false
		}
	}

	if FoundUntilEnd { //jika nama ditemukan
		return "Login berhasil: " + name, nil
	} else {
		return "", fmt.Errorf("Login gagal: data mahasiswa tidak ditemukan")
	}
	// return "", nil // TODO: replace this
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	if len(id) == 0 || len(name) == 0 || len(studyProgram) == 0 {
		return "", fmt.Errorf("ID, Name or StudyProgram is undefined!")
	}
	var validProgram bool
	// fmt.Println("valid :", validProgram)
	// studentProg := sm.studentStudyPrograms [studyProgram]
	// // fmt.Println("studentprog", studentProg)
	if _, codeProg := sm.studentStudyPrograms[studyProgram]; !codeProg { //MENGECEK APABILA STUDYPROGRAM SAMA DALAM MAP, !CODEPROG ARTINYA TIDAK DITEMUKAN MAKA BERNILAI FALSE DAN PROGRAM DIdalmnya DIJALANKAN
		// fmt.Println("codeprog : ", codeProg)
		validProgram = false
		return "", fmt.Errorf("Study program " + studyProgram + " is not found")
	} else {
		validProgram = true
	}

	var IDFound bool
	// fmt.Println("found", IDFound)
	for _, student := range sm.students {
		if student.ID == id {
			IDFound = true
			break
		} else {
			IDFound = false
		}
	}

	if IDFound {
		return "", fmt.Errorf("Registrasi gagal: id sudah digunakan")
	}
	// fmt.Println("IDfound", IDFound)
	// fmt.Println("validProgram", validProgram)
	if validProgram && !IDFound { //JIKA PROGRAM VALID DAN ID TIDAK DITEMUKAN MAKA DIJJALANKAN PROGRAM , !IDFOUND KARENA IDFOUND SENDIRI FALSE JIKA ID BELUM DIGUNAKAN SEHINGGA AGAR DAPAT SAMA MENGGUNAKAN OPERATOR && DIBUAT SAMA
		newStudent := model.Student{ //AGAR TIPE DATA YANG DITAMBAHKAN SESUAI
			ID:           id,
			Name:         name,
			StudyProgram: studyProgram,
		}
		sm.students = append(sm.students, newStudent)
		return "Registrasi berhasil: " + name + " (" + studyProgram + ")", nil
	}

	return "", nil // TODO: replace this
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if len(code) == 0 {
		return "", fmt.Errorf("ID, Name or StudyProgram is undefined!")
	}
	if _, codeProg := sm.studentStudyPrograms[code]; !codeProg { //akan mencari input code dalam mapp, jika tidak ditemukan (!codeprogram), program didalamnya dijalankna
		return "", fmt.Errorf("Kode program studi tidak ditemukan")

	} else {
		return sm.studentStudyPrograms[code], nil  //KETIKA LAINNYA ATAU KODE SESUAI DENGAN DI MAP MAKA AKAN MENGEMBALIKAN VALUE DARI KEY(CODE) YANG DIBERIKAN
	}
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	var nameFound bool
	for i, student := range sm.students {
		if student.Name == name {
			nameFound = true
			if err := fn(&sm.students[i]); err != nil {    //MEMANGGIL FUNGSI MODIF KETIKA NAMA DITEMUKAN
				return "", err
			}
			break
		}
	}
	if !nameFound {
		return "", fmt.Errorf("Mahasiswa tidak ditemukan.")
	}
	return "Program studi mahasiswa berhasil diubah.", nil // TODO: replace this
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		if _, correct := sm.studentStudyPrograms[programStudi]; !correct {  //MEMERIKSA APAKAH INPUT PROGRAM STUDI SAMA
			return fmt.Errorf("Kode program studi tidak ditemukan")
		}
		s.StudyProgram = programStudi
		return nil // TODO: replace this
	}
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
