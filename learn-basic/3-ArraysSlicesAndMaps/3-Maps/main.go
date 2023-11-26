package main

import (
	"fmt"
)

func main()  {
	

	// Maps declaration options
	// #########################
	
	fmt.Println("// Maps")
	fmt.Println("// #####")
	fmt.Println()

	fmt.Println("// Maps declaration options")

	// Option 1 
	var map_1 map[string]string //nil map
	map_1 = make(map[string]string) // init map 
	map_1["en"] = "English"
	fmt.Println("map_1 = ", map_1)
	fmt.Println()
	
	// Option 2
	var map_2 map[string]string = make(map[string]string) 
	map_2["he"] = "Hebrew"
	fmt.Println("map_2 = ", map_2)
	fmt.Println()
	
	// Option 3
	map_3 := make(map[string]string)
	map_3["fr"] = "French"
	fmt.Println("map_3 = ", map_3)
	fmt.Println()
	
	// Option 4
	map_4 := map[string]string{
		"en": "English", 
		"fr": "French",
	}
	fmt.Println("map_4 = ", map_4)
	fmt.Println("len(map_4) = ", len(map_4))
	fmt.Println()
	
	// Getting a key
	fmt.Println("// Getting a key")
	fmt.Println("keyVal, isFound := map_4[\"fr\"]")
	keyVal, isFound := map_4["fr"]
	fmt.Printf("isFound = %v, keyVal = %q\n", isFound, keyVal)
	fmt.Println("keyVal, isFound := map_4[\"hh\"]")
	keyVal, isFound = map_4["hh"]
	fmt.Printf("isFound = %v, keyVal = %q\n", isFound, keyVal)
	fmt.Println()
	
	// Adding a key
	fmt.Println("// Adding a key")
	fmt.Println("map_4[\"it\"] = \"Italian\"")
	map_4["it"] = "Italian"
	fmt.Println("map_4 = ", map_4)
	fmt.Println("keyVal, isFound := map_4[\"it\"]")
	keyVal, isFound = map_4["it"]
	fmt.Println(isFound, keyVal)
	fmt.Println()
	
	// Updating a key
	fmt.Println("// Updating a key")
	fmt.Println("map_4 =", map_4)
	map_4["en"] = "English Language"
	fmt.Println("map_4 =", map_4)
	fmt.Println()
	
	// Deleting a key
	fmt.Println("// Deleting a key")
	fmt.Println("map_4 =", map_4)
	delete(map_4,"en")
	fmt.Println("map_4 =", map_4)
	fmt.Println()
	
	// Adding a key
	fmt.Println("// Adding a key")
	map_4["en"] = "English"
	fmt.Println("map_4 =", map_4)
	fmt.Println()
	
	// Iterating a map
	fmt.Println("// Iterating a map")
	fmt.Println("map_4 =", map_4)
	for key, value := range(map_4){
		fmt.Println(key, "=>", value)
	}
	fmt.Println()
	
	// Truncating a map
	fmt.Println("// Truncating a map")
	
	// Option 1 - delete by iterating 
	fmt.Println("// Option 1 - delete by iterating ")
	for key := range map_4 {
		fmt.Printf("delete(map_4,%s)\n", key)
		delete(map_4,key)
	}
	
	fmt.Println("map_4 =", map_4)
	fmt.Println()
	
	
	// Option 2 - re init1
	fmt.Println("// Option 2 - re init")
	map_4 = map[string]string{
		"en": "English", 
		"fr": "French",
		"it": "Italian",
	}
	fmt.Println("map_4 =", map_4)
	fmt.Println()
	
	
	// Option 3 - re init2
	fmt.Println("// Option 3 - re init2")
	fmt.Println("map_4 =", map_4)
	fmt.Println("map_4 = make(map[string]string)")
	map_4 = make(map[string]string)
	fmt.Println(map_4)
	fmt.Println()
}



