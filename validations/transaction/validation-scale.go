package transaction

import (
	"github.com/LerianStudio/plugin-sdk/validations/transaction/model"
	"math"
	"math/big"
	"strings"
)

// FindScale Function to find the scale for any value of a value
func FindScale(asset string, v float64, s int64) model.Amount {
	valueString := big.NewFloat(v).String()
	parts := strings.Split(valueString, ".")

	scale := s
	value := int64(v)

	if len(parts) > 1 {
		scale = int64(len(parts[1]))
		value = UndoScale(v, scale)

		if parts[1] != "0" {
			scale += s
		}
	}

	amount := model.Amount{
		Asset: asset,
		Value: value,
		Scale: scale,
	}

	return amount
}

// Scale func scale: (V * 10^ (S0-S1))
func Scale(v, s0, s1 int64) int64 {
	return int64(float64(v) * math.Pow(10, float64(s1)-float64(s0)))
}

// UndoScale Function to undo the scale calculation
func UndoScale(v float64, s int64) int64 {
	return int64(v * math.Pow(10, float64(s)))
}
