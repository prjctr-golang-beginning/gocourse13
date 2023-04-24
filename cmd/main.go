package main

import "fmt"

// Інтерфейс Студентів
type Student interface {
	Name() string
	Grade() int
}

// Базовий тип Студентів
type BaseStudent struct {
	name  string
	grade int
}

func (s *BaseStudent) Name() string {
	return s.name
}

func (s *BaseStudent) Grade() int {
	return s.grade
}

// Декоратор, що додає функціональність "відмінника"
type HonorsDecorator struct {
	student Student
}

func (hd *HonorsDecorator) Name() string {
	return hd.student.Name() + ` відмінник`
}

func (hd *HonorsDecorator) Grade() int {
	if hd.student.Grade() < 90 {
		return 90
	}
	return hd.student.Grade()
}

// Декоратор, що додає функціональність "спортсмена"
type SportDecorator struct {
	student Student
}

func (hd *SportDecorator) Name() string {
	return hd.student.Name() + ` спортсмен`
}

func (hd *SportDecorator) Grade() int {
	if hd.student.Grade() < 90 {
		return 110
	}
	return hd.student.Grade()
}

// Головна функція
func main() {
	// Створення базового студента
	baseStudent := &BaseStudent{name: "John", grade: 80}
	fmt.Println("Base student:")
	fmt.Println(baseStudent.Name())
	fmt.Println(baseStudent.Grade())

	// Створення студента з функціональністю "відмінника"
	sportStudent := &SportDecorator{student: baseStudent}
	sportHonorsStudent := &HonorsDecorator{student: sportStudent}
	fmt.Println("Sport Honors student:")
	fmt.Println(sportHonorsStudent.Name())
	fmt.Println(sportHonorsStudent.Grade())

	// exercise:
	// Треба зробити агрегат, який будує керамічну чашку.
	// Для звичайних замовників вистачить білої чашки без ручки.
	// Для частих замовників треба із ручкою, і ще ддодати можливість обирати колір виробу.
	// Для преміальних замовників можна або інхрустувати золотом, або залишити відтіск пальців.
	//
	// Від кількості декораторів змінюється час виробництва, ціна, назва агрегату і кінцевий вироб.
}
