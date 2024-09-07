package helper

import (
        "encoding/json"
        "errors"
)

func ToString(data any) (string, error) {
        bytes, err := json.Marshal(data)
        if err != nil {
                return "", errors.New("failed to convert")
        }
        str := string(bytes)
        return str, nil
}