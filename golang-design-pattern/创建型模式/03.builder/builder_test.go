package builder

import "testing"

func TestBuilder(t *testing.T) {
	tests := []struct{
		name string
		expect interface{}
		actual interface{}
		builder Builder
	}{
		{
			name: "stringBuilder",
			expect: "123",
			builder: &StringBuilder{},
		},
		{
			name: "intBuilder",
			expect: 6,
			builder: &IntBuilder{},
		},
	}
	var director *Director
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			director = NewDirector(test.builder)
			director.Construct()
			test.actual = test.builder.GetResult()
			if test.actual != test.expect {
				t.Fatalf("expect result is %v,but actual is %v\n",test.expect,test.actual)
			}
		})
	}
}
