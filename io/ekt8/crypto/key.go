package crypto

const (
	PublicKeyLength  = 128
	PrivateKeyLength = 256
)

type PublicKey [PublicKeyLength]byte
type PrivateKey [PrivateKeyLength]byte
