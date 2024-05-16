package utility

import resType "github.com/aliamerj/aircup/server/internal/constant"

func SuccesRespnse(message string, resType resType.ResType, body interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"type":    resType,
		"body":    body,
	}
}
