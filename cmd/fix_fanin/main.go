package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PatientData - структура для зберігання медичних даних пацієнта
type PatientData struct {
	ID        int
	BloodTest map[string]float64
	UrineTest map[string]float64
	ECG       string
}

type AnalysisResult struct {
	ID     int
	Result string
}

// generateRandomData - генерує випадкові медичні дані для прикладу
func generateRandomData(id int) PatientData {
	return PatientData{
		ID: id,
		BloodTest: map[string]float64{
			"glucose":    rand.Float64() * 5,
			"hemoglobin": rand.Float64() * 150,
		},
		UrineTest: map[string]float64{
			"protein": rand.Float64() * 0.3,
			"sugar":   rand.Float64() * 0.2,
		},
		ECG: "normal",
	}
}

// collectPatientData - збирає дані пацієнтів
func collectPatientData(bloodTests, urineTests chan<- PatientData) {
	go func() {
		for i := 1; i <= 5; i++ {
			data := generateRandomData(i)
			bloodTests <- data
			urineTests <- data
		}
		close(bloodTests)
		close(urineTests)
	}()
}

// analyzeBloodTest - аналізує кров
func analyzeBloodTest(in <-chan PatientData) <-chan AnalysisResult {
	out := make(chan AnalysisResult)
	go func() {
		for pd := range in {
			glucose := pd.BloodTest["glucose"]
			hemoglobin := pd.BloodTest["hemoglobin"]
			result := fmt.Sprintf("Blood Test - Glucose %.2f, Hemoglobin %.2f", glucose, hemoglobin)
			out <- AnalysisResult{ID: pd.ID, Result: result}
		}
		close(out)
	}()
	return out
}

// analyzeUrineTest - аналізує сечу
func analyzeUrineTest(in <-chan PatientData) <-chan AnalysisResult {
	out := make(chan AnalysisResult)
	go func() {
		for pd := range in {
			protein := pd.UrineTest["protein"]
			sugar := pd.UrineTest["sugar"]
			result := fmt.Sprintf("Urine Test - Protein %.2f, Sugar %.2f", protein, sugar)
			out <- AnalysisResult{ID: pd.ID, Result: result}
		}
		close(out)
	}()
	return out
}

// mergeResults - об'єднує результати різних аналізів
func mergeResults(cs ...<-chan AnalysisResult) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	combinedResults := make(map[int]string)

	output := func(c <-chan AnalysisResult) {
		for result := range c {
			if existing, found := combinedResults[result.ID]; found {
				combinedResults[result.ID] = existing + "; " + result.Result
			} else {
				combinedResults[result.ID] = result.Result
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		for _, combinedResult := range combinedResults {
			out <- combinedResult
		}
		close(out)
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bloodTests := make(chan PatientData)
	urineTests := make(chan PatientData)
	collectPatientData(bloodTests, urineTests)

	bloodTestData := analyzeBloodTest(bloodTests)
	urineTestData := analyzeUrineTest(urineTests)

	mergedResults := mergeResults(bloodTestData, urineTestData)

	for result := range mergedResults {
		fmt.Println(result)
	}
}
