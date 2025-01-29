package halstead

import (
	"halstead/colors"
	"halstead/tokenizer"
	"math"
)

// Analyze source code with a specific parser
func AnalyzeSourceCode(filepath string) (Halstead, error) {

	tokens := tokenizer.Tokenize(filepath)

	metrics := GetHalsteadMetrics(tokens)

	return metrics, nil
}

type Halstead struct {
	// Number of distinct operators
	N1 int
	// Number of distinct operands
	N2 int
	// Total number of operators
	n1 int
	// Total number of operands
	n2 int

	// Program vocabulary
	N int
	// Program length
	n int
	// Calculated program length
	Np float64
	// Calculated program volume
	V float64
	// Calculated program difficulty
	D float64
	// Calculated program effort
	E float64
	// Calculated program time
	T float64
	// Calculated program bugs
	B float64
}

func (h *Halstead) Calculate() {
	h.N = h.N1 + h.N2
	h.n = h.n1 + h.n2

	h.Np = float64(h.N1) * (float64(h.N2) / 2)
	h.V = float64(h.N) * (math.Log2(float64(h.N)))
	h.D = (float64(h.N1) / 2) * (float64(h.n2) / float64(h.N2))
	h.E = h.D * h.V
	h.T = h.E / 18   // 18 is the average number of bugs per hour
	h.B = h.V / 3000 // 3000 is the average volume of a bug
}

func (h *Halstead) Print() {
	colors.BLUE.Println("--------------------- Halstead Metrics ---------------------")
	colors.ORANGE.Printf("Number of distinct operators \t(N1): %d\n", h.N1)
	colors.PURPLE.Printf("Number of distinct operands \t(N2): %d\n", h.N2)
	colors.ORANGE.Printf("Total number of operators \t(n1): %d\n", h.n1)
	colors.PURPLE.Printf("Total number of operands \t(n2): %d\n", h.n2)
	colors.GREEN.Printf("Program vocabulary \t\t(N): %d\n", h.N)
	colors.GREEN.Printf("Program length \t\t\t(n): %d\n", h.n)
	colors.YELLOW.Printf("Calculated program length \t(Np): %.2f\n", h.Np)
	colors.YELLOW.Printf("Calculated program volume \t(V): %.2f\n", h.V)
	colors.RED.Printf("Calculated program difficulty \t(D): %.2f\n", h.D)
	colors.RED.Printf("Calculated program effort \t(E): %.2f\n", h.E)
	colors.BROWN.Printf("Calculated program time \t(T): %.2f\n", h.T)
	colors.BROWN.Printf("Calculated program bugs \t(B): %.2f\n", h.B)
	colors.BLUE.Println("------------------------------------------------------------")
}

func GetHalsteadMetrics(tokens []tokenizer.Token) Halstead {

	h := Halstead{}

	operators := make(map[string]bool)
	operands := make(map[string]bool)

	for _, token := range tokens {
		switch token.Type {
		case tokenizer.OPERATOR:
			operators[token.Value] = true
			h.n1++
		case tokenizer.OPERAND:
			operands[token.Value] = true
			h.n2++
		}
	}

	h.N1 = len(operators)
	h.N2 = len(operands)

	h.Calculate()

	return h
}
