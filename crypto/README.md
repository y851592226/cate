# crypto

> Encryption and decryption

### Simple Example

```go
type Payload struct {
			A int
		}
key := "1234567890abaaaa"
encoder := crypto.NewEncoder(key, aes.AES{})
payload := Payload{A: 12}
data, err := encoder.Encode(payload)
```

### doc

```go
type Encoder interface {
	Encode(i interface{}) ([]byte, error)
	Decode(data []byte, i interface{}) error
}

func NewEncoder(key string, encryper Encryper) Encoder

type Encryper interface {
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
}
```

