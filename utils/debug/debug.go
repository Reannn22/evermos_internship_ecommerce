package debug

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(prefix string, v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s: %s\n", prefix, string(b))
}
