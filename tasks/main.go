package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Patient - структура для зберігання даних пацієнта
type Patient struct {
	ID      int
	Name    string
	Age     int
	Country string
	Health  string
}

// generateRandomPatients - генерує велику кількість пацієнтів
func generateRandomPatients(n int) []Patient {
	patients := make([]Patient, n)
	for i := 0; i < n; i++ {
		patients[i] = Patient{
			ID:      i,
			Name:    fmt.Sprintf("Patient%d", i),
			Age:     rand.Intn(100),
			Country: getRandomCountry(),
			Health:  getRandomHealthStatus(),
		}
	}
	return patients
}

// getRandomHealthStatus - випадково вибирає стан здоров'я
func getRandomHealthStatus() string {
	statuses := []string{"Good", "Fair", "Poor"}
	return statuses[rand.Intn(len(statuses))]
}

// getRandomHealthStatus - випадково вибирає стан здоров'я
func getRandomCountry() string {
	countires := []string{"Spain", "France", "USA", "Japan"}
	return countires[rand.Intn(len(countires))]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Генерація даних
	numOfPatients := 1000 // Кількість пацієнтів для генерації
	patients := generateRandomPatients(numOfPatients)

	for i, patient := range patients {
		fmt.Printf("Пацієнт %d: %#v\n", i, patient)
	}
}
