package model

// Responses is the struct result of transaction validation process
type Responses struct {
	Total        int64
	Asset        string
	From         map[string]Amount
	To           map[string]Amount
	Sources      []string
	Destinations []string
	Aliases      []string
}

// Amount is the struct designed to represent the amount of an operation.
type Amount struct {
	Asset     string `json:"asset,omitempty" validate:"required" example:"BRL"`
	Value     int64  `json:"value,omitempty" validate:"required" example:"1000"`
	Scale     int64  `json:"scale,omitempty" validate:"gte=0" example:"2"`
	Operation string `json:"operation,omitempty"`
}

// Share is the struct designed to represent the sharing fields of an operation.
type Share struct {
	Percentage             int64 `json:"percentage,omitempty" validate:"required"`
	PercentageOfPercentage int64 `json:"percentageOfPercentage,omitempty"`
}

// Send is the struct designed to represent the sending fields of an operation.
type Send struct {
	Asset      string     `json:"asset,omitempty" validate:"required" example:"BRL"`
	Value      int64      `json:"value,omitempty" validate:"required" example:"1000"`
	Scale      int64      `json:"scale,omitempty" validate:"gte=0" example:"2"`
	Source     Source     `json:"source,omitempty" validate:"required"`
	Distribute Distribute `json:"distribute,omitempty" validate:"required"`
}

// FromTo is the struct designed to represent the from/to fields of an operation.
type FromTo struct {
	Account         string         `json:"account,omitempty" example:"@person1"`
	Amount          *Amount        `json:"amount,omitempty"`
	Share           *Share         `json:"share,omitempty"`
	Remaining       string         `json:"remaining,omitempty" example:"remaining"`
	Rate            *Rate          `json:"rate,omitempty"`
	Description     string         `json:"description,omitempty" example:"description"`
	ChartOfAccounts string         `json:"chartOfAccounts" example:"1000"`
	Metadata        map[string]any `json:"metadata" validate:"dive,keys,keymax=100,endkeys,nonested,valuemax=2000"`
	IsFrom          bool           `json:"isFrom,omitempty" example:"true"`
}

// Source is the struct designed to represent the source fields of an operation.
type Source struct {
	Remaining string   `json:"remaining,omitempty" example:"remaining"`
	From      []FromTo `json:"from,omitempty" validate:"singletransactiontype,required,dive"`
}

// Distribute is the struct designed to represent the distribution fields of an operation.
type Distribute struct {
	Remaining string   `json:"remaining,omitempty"`
	To        []FromTo `json:"to,omitempty" validate:"singletransactiontype,required,dive"`
}

// Rate is the struct designed to represent the rate fields of an operation.
type Rate struct {
	From       string `json:"from" validate:"required" example:"BRL"`
	To         string `json:"to" validate:"required" example:"USDe"`
	Value      int64  `json:"value" validate:"required" example:"1000"`
	Scale      int64  `json:"scale" validate:"gte=0" example:"2"`
	ExternalID string `json:"externalId" validate:"uuid,required" example:"00000000-0000-0000-0000-000000000000"`
}
