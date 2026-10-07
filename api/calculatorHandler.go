package api

import (
	"encoding/json"
	"net/http"
)

type CalculatorHandler struct{}

// Calculate godoc
// @Summary Calculate
// @Description Perform a mathematical operation
// @Tags calculator
// @Accept json
// @Produce json
// @Param request body Request true "Calculator request"
// @Success 200 {object} map[string]float64
// @Failure 400 {string} string
// @Failure 405 {string} string
// @Router /calculate [post]
func (h *CalculatorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	a, ok := request.Number1.(float64)
	if !ok {
		http.Error(w, "Parameter 'a' must be a number", http.StatusBadRequest)
		return
	}

	b, ok := request.Number2.(float64)
	if !ok {
		http.Error(w, "Parameter 'b' must be a number", http.StatusBadRequest)
		return
	}

	var result float64

	switch request.Operation {

	case "addition":
		result = a + b

	case "subtraction":
		result = a - b

	case "multiplication":
		result = a * b

	case "division":
		if b == 0 {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}

		result = a / b

	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	response := map[string]float64{
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}