package utils

import (
	"bytes"
	"encoding/binary"
)

/**
 * 将一个int64数值类型转换为【】byte类型
 */
func IntToByte(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
