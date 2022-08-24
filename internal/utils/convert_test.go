package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDenomination(t *testing.T) {
	type fields struct {
		amount   string
		decimals int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "1 eth",
			fields: fields{amount: "1", decimals: 18},
			want:   "1000000000000000000",
		},
		{
			name:   "0.5 eth",
			fields: fields{amount: "0.5", decimals: 18},
			want:   "500000000000000000",
		},
		{
			name:   "0.001 eth",
			fields: fields{amount: "0.001", decimals: 18},
			want:   "1000000000000000",
		},
		{
			name:   "0.000025 eth",
			fields: fields{amount: "0.000025", decimals: 18},
			want:   "25000000000000",
		},
		{
			name:   "8.000025 eth",
			fields: fields{amount: "8.000025", decimals: 18},
			want:   "8000025000000000000",
		},
		{
			name:   "1 eth in gwei",
			fields: fields{amount: "1", decimals: 9},
			want:   "1000000000",
		},
		{
			name:   ".5 eth in gwei",
			fields: fields{amount: ".5", decimals: 9},
			want:   "500000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ToDenomination(tt.fields.amount, tt.fields.decimals)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, actual.String())
		})
	}
}
