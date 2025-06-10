package util

import (
	"fmt"
	"time"
)

func GenerateID() string {
	return fmt.Sprintf("room_%d", time.Now().UnixNano())
}
