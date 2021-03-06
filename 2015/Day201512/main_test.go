package Day201512

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(splitAndAdd("[1,2,3]")).Should(Equal(6))
	Expect(splitAndAdd(`{"a":2,"b":4}`)).Should(Equal(6))
	Expect(splitAndAdd(`[[[3]]]`)).Should(Equal(3))
	Expect(splitAndAdd(`{"a":{"b":4},"c":-1}`)).Should(Equal(3))
	Expect(splitAndAdd(`{"a":[-1,1]}`)).Should(Equal(0))
	Expect(splitAndAdd(`[-1,{"a":1}]`)).Should(Equal(0))
	Expect(splitAndAdd(`[]`)).Should(Equal(0))
	Expect(splitAndAdd(`{}`)).Should(Equal(0))

	var customData interface{}
	err := json.Unmarshal([]byte(Entry.PuzzleInput()), &customData)
	if err != nil {
		t.Error(err)
	}

	count := parseAndWalk(customData, 0, 0)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(count)
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
