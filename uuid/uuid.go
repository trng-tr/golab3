package uuid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// generation des uuid

func GenerateUuid(str string) string {
	str = strings.ReplaceAll(strings.TrimSpace(str), " ", "-")
	str = str[:4]

	var uuid string = uuid.NewString()[:6]
	return fmt.Sprintf("%s-%s", str, uuid)
}
