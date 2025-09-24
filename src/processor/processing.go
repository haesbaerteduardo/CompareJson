package processor

import (
	"encoding/json"
	"os"
	"strings"
)

func ExtractIDs(filename string) (map[string]bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	ids := make(map[string]bool)
	extractIDsRecursive(jsonData, ids)
	return ids, nil
}

func extractIDsRecursive(data interface{}, ids map[string]bool) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if isIDField(key) {
				if str, ok := value.(string); ok && str != "" {
					ids[str] = true
				}
			}
		}
		for _, value := range v {
			extractIDsRecursive(value, ids)
		}
	case []interface{}:
		for _, item := range v {
			extractIDsRecursive(item, ids)
		}
	}
}

func isIDField(fieldName string) bool {
	idFields := []string{
		"id", "resourceId", "sid", "uuid", "identifier",
		"taskId", "processId", "elementId", "_id",
	}

	fieldLower := strings.ToLower(fieldName)
	for _, idField := range idFields {
		if strings.Contains(fieldLower, idField) {
			return true
		}
	}
	return false
}
