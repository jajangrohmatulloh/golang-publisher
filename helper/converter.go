package helper

import (
	"encoding/json"
	"fmt"
	"errors"
)

func ToString(data any) (string, error) {
	bytes, err := json.Marshal(data)
        if err != nil {
                fmt.Println("Error marshalling:", err)
                return "", errors.New("failed to convert")
        }
        str := string(bytes)
        return str, nil
}