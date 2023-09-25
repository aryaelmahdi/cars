package helper

import (
	"fmt"
	"time"
)

func RenameFile(fileName string) string {
	return fmt.Sprint("image", time.Now().UnixNano())
}
