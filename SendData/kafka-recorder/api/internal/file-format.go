package internal

import (
	"encoding/binary"
	"os"

	"github.com/Shopify/sarama"
)

func WriteMessage(file *os.File, msg *sarama.ConsumerMessage) (err error) {
	if err = writeValue(file, msg.Key); err == nil {
		err = writeValue(file, msg.Value)
	}
	return err
}

func ReadMessage(file *os.File) (key, value []byte, err error) {
	if key, err = readValue(file); err == nil {
		value, err = readValue(file)
	}
	return key, value, err
}

func writeValue(file *os.File, value []byte) (err error) {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, uint32(len(value)))
	if _, err = file.Write(buffer); err == nil {
		_, err = file.Write(value)
	}
	return err
}

func readValue(file *os.File) (value []byte, err error) {
	buffer := make([]byte, 4)
	if _, err = file.Read(buffer); err == nil {
		length := binary.BigEndian.Uint32(buffer)
		value = make([]byte, length)
		_, err = file.Read(value)
	}
	return value, err
}
