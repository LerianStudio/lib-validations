package transaction

import (
	"errors"
	"math"
	"plugin-sdk/validations/transaction/model"
	"strings"
)

// ValidateSendSourceAndDistribute Validate send and distribute totals making all the necessary calculation
// returning model.Responses
func ValidateSendSourceAndDistribute(send model.Send) (*model.Responses, error) {
	response := &model.Responses{
		Total:        send.Value,
		Asset:        send.Asset,
		From:         make(map[string]model.Amount),
		To:           make(map[string]model.Amount),
		Sources:      make([]string, 0),
		Destinations: make([]string, 0),
		Aliases:      make([]string, 0),
	}

	var (
		sourcesTotal      int64
		destinationsTotal int64
	)

	t := make(chan int64)
	ft := make(chan map[string]model.Amount)
	sd := make(chan []string)

	go calculateTotal(send.Source.From, send, t, ft, sd)
	sourcesTotal = <-t
	response.From = <-ft
	response.Sources = <-sd
	response.Aliases = append(response.Aliases, response.Sources...)

	go calculateTotal(send.Distribute.To, send, t, ft, sd)
	destinationsTotal = <-t
	response.To = <-ft
	response.Destinations = <-sd
	response.Aliases = append(response.Aliases, response.Destinations...)

	for _, source := range response.Sources {
		if _, ok := response.To[source]; ok {
			return nil, errors.New("transaction value mismatch when validate send, source and distribute structs")
		}
	}

	for _, destination := range response.Destinations {
		if _, ok := response.From[destination]; ok {
			return nil, errors.New("transaction value mismatch when validate send, source and distribute structs")
		}
	}

	if math.Abs(float64(response.Total)-float64(sourcesTotal)) != 0 {
		return nil, errors.New("transaction value mismatch when validate send, source and distribute structs")
	}

	if math.Abs(float64(sourcesTotal)-float64(destinationsTotal)) != 0 {
		return nil, errors.New("transaction value mismatch when validate send, source and distribute structs")
	}

	return response, nil
}

// calculateTotal Calculate total for sources/destinations based on shares, amounts and remains
func calculateTotal(fromTos []model.FromTo, send model.Send, t chan int64, ft chan map[string]model.Amount, sd chan []string) {
	fmto := make(map[string]model.Amount)
	scdt := make([]string, 0)

	total := model.Amount{
		Asset: send.Asset,
		Scale: 0,
		Value: 0,
	}

	remaining := model.Amount{
		Asset: send.Asset,
		Scale: send.Scale,
		Value: send.Value,
	}

	for i := range fromTos {
		if fromTos[i].Share != nil && fromTos[i].Share.Percentage != 0 {
			percentage := fromTos[i].Share.Percentage

			percentageOfPercentage := fromTos[i].Share.PercentageOfPercentage
			if percentageOfPercentage == 0 {
				percentageOfPercentage = 100
			}

			shareValue := float64(send.Value) * (float64(percentage) / float64(percentageOfPercentage))
			amount := FindScale(send.Asset, shareValue, send.Scale)

			normalize(&total, &amount, &remaining)
			fmto[fromTos[i].Account] = amount
		}

		if fromTos[i].Amount != nil && fromTos[i].Amount.Value > 0 && fromTos[i].Amount.Scale > -1 {
			amount := model.Amount{
				Asset: fromTos[i].Amount.Asset,
				Scale: fromTos[i].Amount.Scale,
				Value: fromTos[i].Amount.Value,
			}

			normalize(&total, &amount, &remaining)
			fmto[fromTos[i].Account] = amount
		}

		if IsNilOrEmpty(&fromTos[i].Remaining) {
			total.Value += remaining.Value

			fmto[fromTos[i].Account] = remaining
			fromTos[i].Amount = &remaining
		}

		scdt = append(scdt, fromTos[i].Account)
	}

	ttl := total.Value
	if total.Scale > send.Scale {
		ttl = Scale(total.Value, total.Scale, send.Scale)
	}

	t <- ttl
	ft <- fmto
	sd <- scdt
}

func IsNilOrEmpty(s *string) bool {
	return s == nil || strings.TrimSpace(*s) == "" || strings.TrimSpace(*s) == "null" || strings.TrimSpace(*s) == "nil"
}

// normalize func that normalize scale from all values
func normalize(total, amount, remaining *model.Amount) {
	if total.Scale < amount.Scale {
		if total.Value != 0 {
			v0 := Scale(total.Value, total.Scale, amount.Scale)

			total.Value = v0 + amount.Value
		} else {
			total.Value += amount.Value
		}

		total.Scale = amount.Scale
	} else {
		if total.Value != 0 {
			v0 := Scale(amount.Value, amount.Scale, total.Scale)

			total.Value += v0

			amount.Value = v0
			amount.Scale = total.Scale
		} else {
			total.Value += amount.Value
			total.Scale = amount.Scale
		}
	}

	if remaining.Scale < amount.Scale {
		v0 := Scale(remaining.Value, remaining.Scale, amount.Scale)

		remaining.Value = v0 - amount.Value
		remaining.Scale = amount.Scale
	} else {
		v0 := Scale(amount.Value, amount.Scale, remaining.Scale)

		remaining.Value -= v0
	}
}
