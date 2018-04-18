package sum_test

import (
	"fmt"

	"github.com/jbbarquero/gobenchmarking/sum"
)

//godoc -http=:6060 then go to http://localhost:6060/pkg/
func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)

}
