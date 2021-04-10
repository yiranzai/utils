package tool

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GetUUID 获取UUID
func GetUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
