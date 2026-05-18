package local

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func generateFileName(originalName string) string {

	extension := filepath.Ext(originalName)

	return fmt.Sprintf(
		"%s-%d%s",
		uuid.New().String(),
		time.Now().Unix(),
		extension,
	)
}
