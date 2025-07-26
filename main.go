package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type BMIResult struct {
	Weight           float64
	Height           float64
	BMI              float64
	IdealWeight      float64
	WeightDifference float64
	IsOverWeight     bool
	IsUnderWeight    bool
	HasResult        bool
}

func calculateBMI(weight, height float64) BMIResult {
	heightInMeters := height / 100
	bmi := weight / (heightInMeters * heightInMeters)
	idealWeight := (heightInMeters * heightInMeters) * 22
	weightDifference := weight - idealWeight
	
	return BMIResult{
		Weight:           weight,
		Height:           height,
		BMI:              bmi,
		IdealWeight:      idealWeight,
		WeightDifference: weightDifference,
		IsOverWeight:     weightDifference > 0,
		IsUnderWeight:    weightDifference < 0,
		HasResult:        true,
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	
	var result BMIResult
	
	if r.Method == "POST" {
		weightStr := r.FormValue("weight")
		heightStr := r.FormValue("height")
		
		weight, weightErr := strconv.ParseFloat(weightStr, 64)
		height, heightErr := strconv.ParseFloat(heightStr, 64)
		
		if weightErr == nil && heightErr == nil && weight > 0 && height > 0 {
			result = calculateBMI(weight, height)
		}
	}
	
	err := tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}