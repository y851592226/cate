package crypto

import (
	"encoding/hex"
	"encoding/json"
)

type Encryper interface {
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
}

type Encoder interface {
	Encode(i interface{}) ([]byte, error)
	Decode(data []byte, i interface{}) error
	EncodeToString(i interface{}) (string, error)
	DecodeString(s string, i interface{}) error
}

type encoder struct {
	key      []byte
	encryper Encryper
}

func NewEncoder(key string, encryper Encryper) Encoder {
	if len(key) < 16 {
		key = key + string(make([]byte, 16-len(key)))
	}
	return &encoder{
		key:      []byte(key),
		encryper: encryper,
	}
}

func (e *encoder) Encode(i interface{}) ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	ciphertext, err := e.encryper.Encrypt(e.key, data)
	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(ciphertext)), nil
}

func (e *encoder) Decode(data []byte, i interface{}) error {
	decoded, err := hex.DecodeString(string(data))
	if err != nil {
		return err
	}
	data, err = e.encryper.Decrypt(e.key, decoded)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}

func (e *encoder) EncodeToString(i interface{}) (string, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	ciphertext, err := e.encryper.Encrypt(e.key, data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func (e *encoder) DecodeString(s string, i interface{}) error {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	data, err := e.encryper.Decrypt(e.key, decoded)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}
