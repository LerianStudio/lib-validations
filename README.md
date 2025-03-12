# Lib Validations

The lib-validation provides a set of reusable validation functions designed to ensure data integrity, consistency, and proper formatting.

## 📦 Install

```bash
go get -u github.com/LerianStudio/lib-validations
```

## Features 💡

- *Validation of Send, Source, and Distribution:* Ensures that transaction values match across different structures.

- *Total Calculation:* Computes total amounts for sources and destinations based on shares, amounts, and remaining values.

- *Scale Normalization:* Adjusts scales to ensure consistency in calculations.

- *Scale Detection:* Identifies and adjusts the scale of a given value.

- *Scale Conversion:* Transforms values between different scales.

- *Scale Reversal:* Restores a scaled value to its original integer format.


## Usage 🚀

### 📌 Scales functions 

```go
package main

import (
	"fmt"
	"github.com/LerianStudio/lib-validations/transaction"
)

func main() {
	amount := transaction.FindScale("BRL", 123.456, 2)
	fmt.Println(amount) // {BRL 123456 5 }

	scaledValue := transaction.Scale(123456, 5, 2)
	fmt.Println(scaledValue) // 123

	originalValue := transaction.UndoScale(123.456, 2)
	fmt.Println(originalValue) // 12345
}

```

### 📌 Validation functions

```go
package main

import (
	"fmt"
	"github.com/LerianStudio/lib-validations/transaction"
	"github.com/LerianStudio/lib-validations/transaction/model"
)

func main() {
	send := model.Send{
		Asset: "USD",
		Scale: 2,
		Value: 10000,
		Source: model.Source{
			From: []model.FromTo{
				{Account: "account_1", Amount: &model.Amount{Asset: "USD", Scale: 2, Value: 5000}},
				{Account: "account_2", Amount: &model.Amount{Asset: "USD", Scale: 2, Value: 5000}},
			},
		},
		Distribute: model.Distribute{
			To: []model.FromTo{
				{Account: "account_3", Amount: &model.Amount{Asset: "USD", Scale: 2, Value: 10000}},
			},
		},
	}

	response, err := transaction.ValidateSendSourceAndDistribute(send)
	if err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation successful:", response)
	}
}

```

## Functions description 📋

```FindScale(asset string, v float64, s int64) model.Amount```
- Determines the appropriate scale for a given value, adjusting it accordingly.

```Scale(v, s0, s1 int64) int64```
- Converts a value from one scale to another using the formula: V * 10^(S0 - S1).

```UndoScale(v float64, s int64) int64```
- Reverts a scaled value back to its original integer representation.

```IsNilOrEmpty(s *string) bool```
- Checks if a string is empty, nil, or contains invalid values like "null" or "nil".

```normalize(total, amount, remaining *model.Amount)```
- Normalizes the scale of all values to maintain consistency across calculations.

```FindScale(asset string, v float64, s int64) model.Amount```
- Determines the appropriate scale for a given value, adjusting it accordingly.

```Scale(v, s0, s1 int64) int64```
- Converts a value from one scale to another using the formula: V * 10^(S0 - S1).

```UndoScale(v float64, s int64) int64```
- Reverts a scaled value back to its original integer representation.
