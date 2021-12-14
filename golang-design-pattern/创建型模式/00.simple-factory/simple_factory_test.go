package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewApi(1)
	expect := "Hi Timothy"
	actual := api.Say("Timothy")
	if expect != actual {
		t.Fatalf("expect  [%s] ,but actual is [%s]\n", expect, actual)
	}
}
func TestType2(t *testing.T) {
	api := NewApi(2)
	expect := "hello Timothy"
	actual := api.Say("Timothy")
	if expect != actual {
		t.Fatalf("expect  [%s] ,but actual is [%s]\n", expect, actual)
	}
}

func TestApi(t *testing.T) {
	tests := []struct {
		t1     int
		name   string
		expect string
		actual string
	}{
		{
			name:   "type1",
			expect: "Hi Timothy",
			t1:     1,
		},
		{
			name:   "type2",
			expect: "hello Timothy",
			t1:     2,
		},
		{
			name:   "type3",
			expect: "",
			t1:     3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			api := NewApi(test.t1)
			if api == nil {
				return
			}
			test.actual = api.Say("Timothy")
			if test.actual != test.expect {
				t.Fatalf("expect  [%s] ,but actual is [%s]\n", test.expect, test.actual)
			}
		})

	}
}
