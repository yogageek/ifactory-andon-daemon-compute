package db

import (
	"encoding/json"
)

type Test struct {
	A string
	B string
	C string
}

func (t *Test) UnmarshalJSON(data []byte) error {
	type testAlias Test
	test := &testAlias{
		A: "default A",
		B: "default B",
		C: "default C",
	}

	_ = json.Unmarshal(data, test)

	*t = Test(*test)
	return nil
}

var example []byte = []byte(`[{"A": "1", "C": "3"}, {"A": "4"}]`)

// func main() {
// 	out := &[]Test{}
// 	_ = json.Unmarshal(example, &out)
// 	fmt.Print(out)
// }
