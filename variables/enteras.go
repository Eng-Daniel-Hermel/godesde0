package variables

import "fmt"

func MuestroEnteros() {
	var intcomum int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomum = ", intcomum)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intcomum = ", intde64)
}
