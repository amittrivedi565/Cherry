package builder

import (
	"Cherry/dsl"
	"fmt"
)

func BuildDrop(obj dsl.Drop) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", obj.Entity)
}
