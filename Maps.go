package main

import "fmt"

func main() {
	var populationMap map[string]float64 = make(map[string]float64)

	populationMap["India"] = 10800000000
	populationMap["USA"] = 58000000
	populationMap["China"] = 160700000000

	fmt.Println(populationMap)
	delete(populationMap, "China")

	fmt.Println(populationMap)

	var emptyMap1 map[string]string
	fmt.Println(emptyMap1)

	var emptyMap2 map[string]string
	fmt.Println(emptyMap2)

	emptyMap1 = map[string]string{"key1": "value1"}
	emptyMap2 = make(map[string]string)
	emptyMap2["key2"] = "value2"

	fmt.Println(emptyMap1)
	fmt.Println(emptyMap2)

}
