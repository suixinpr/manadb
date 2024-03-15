package debug

import (
	"encoding/json"
	"fmt"
)

func Print(name string, x any) {
	j, _ := json.Marshal(x)
	fmt.Println(name, string(j))
	fmt.Println()
}
