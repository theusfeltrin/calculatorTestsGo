package main

import "testing"

func ShouldSumCorrect(t *testing.T) {
	teste := sum(3, 2)
	resultado := 5.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	teste := sum(3, 2)
	resultado := 4.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSubCorrect(t *testing.T) {
	teste := sub(3, 2)
	resultado := 1.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSubIncorrect(t *testing.T) {
	teste := sub(3, 2)
	resultado := 4.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldMultCorrect(t *testing.T) {
	teste := mult(3, 2)
	resultado := 6.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldMultIncorrect(t *testing.T) {
	teste := mult(3, 2)
	resultado := 9.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldDivideCorrect(t *testing.T) {
	teste := divide(3, 2)
	resultado := 1.5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldDivideIncorrect(t *testing.T) {
	teste := sum(3, 2)
	resultado := 6.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
