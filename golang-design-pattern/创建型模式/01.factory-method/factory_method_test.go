package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	tests := []struct {
		name           string
		s              int
		a, b           int
		actual, expect int
	}{
		{
			name:   "plus",
			a:      1,
			b:      2,
			expect: 3,
			s:      1,
		},
		{
			name:   "plus_error",
			a:      1,
			b:      2,
			expect: 4,
			s:      1,
		},
		{
			name:   "minus",
			a:      1,
			b:      2,
			expect: -1,
			s:      2,
		},
		{
			name:   "minus_error",
			a:      1,
			b:      2,
			expect: 4,
			s:      2,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(test.name, func(t *testing.T) {
			var o operator
			if tt.s == 1 {
				o = PlusOperatorFactory{}.Create()
			} else {
				o = MinusOperatorFactory{}.Create()
			}
			o.SetA(tt.a)
			o.SetB(tt.b)
			tt.actual = o.Result()
			if tt.actual != tt.expect {
               t.Fatalf("expect [%d] ,but actual is [%d]\n", tt.expect, tt.actual)
			}
		})
	}
}
