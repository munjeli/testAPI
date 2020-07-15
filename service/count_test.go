package service

import (
	"encoding/json"
	"testing"
)

var (
	validString = "cat dog cat not! Pickels Jam!!"
)

func validResult() []byte {
	w := WordCount{
		Count: 5,
		WordMap: map[string]int{
			"cat":     2,
			"dog":     1,
			"not":     1,
			"pickels": 1,
			"jam":     1,
		},
	}
	b, err := json.Marshal(w)
	if err != nil {
		return []byte{}
	}
	return b
}

func TestCountWordsInText(t *testing.T) {
	tests := []struct {
		desc    string
		str     string
		want    []byte
		wantErr bool
	}{
		{
			desc:    "test with valid string",
			str:     validString,
			want:    validResult(),
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			result, err := CountWordsInText(test.str)
			if err != nil && !test.wantErr {
				t.Errorf("wantErr: %v, got: %v", test.wantErr, err)
			}
			if string(result) != string(test.want) && !test.wantErr {
				t.Errorf("unexpected error: got: %v, want: %v", string(result), string(test.want))
			}
		})
	}
}
