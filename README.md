# Lib Validations

The lib-validation provides a set of reusable validation functions designed to ensure data integrity, consistency, and proper formatting.

## 📦 Install

```bash
go get -u github.com/LerianStudio/lib-validations
```

## The Lib

### File validation.go

### ``ValidateSendSourceAndDistribute(send model.Send) (*model.Responses, error)``

Validates the consistency of the transaction by checking if the total value matches between sources and destinations.

### ``calculateTotal(fromTos []model.FromTo, send model.Send, t chan int64, ft chan map[string]model.Amount, sd chan []string)``

Computes the total amount for sources and destinations based on shares, amounts, and remaining values.

### ``IsNilOrEmpty(s *string) bool``
Checks if a string is empty, nil, or contains invalid values like "null" or "nil".

### ``normalize(total, amount, remaining *model.Amount)``

Normalizes the scale of all values to maintain consistency across calculations.

### File validation-scale.go

### ``FindScale(asset string, v float64, s int64) model.Amount``

Determines the appropriate scale for a given value, adjusting it accordingly.

### ``Scale(v, s0, s1 int64) int64``

Converts a value from one scale to another using the formula: V * 10^(S0 - S1).

### ``UndoScale(v float64, s int64) int64``

Reverts a scaled value back to its original integer representation.