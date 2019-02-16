package simpletimed

import (
	"fmt"
	"time"
)

var h24 = time.Hour * 24

// c formats a timestamp as HH:MM
func c(t time.Time) string {
	return fmt.Sprintf("%.2d:%.2d", t.Hour(), t.Minute())
}
