package encryption

type Encryptor interface {
	Encrypt(password string) (string, error)
	Compare(hash, password string) (bool, error)
}

type EncryptorConstructor func() Encryptor
type Method string

var implementations = make(map[Method]EncryptorConstructor)

func GetInstance(method Method) Encryptor {
	return implementations[method]()
}
