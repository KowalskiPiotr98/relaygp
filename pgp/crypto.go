package pgp

import (
	"github.com/ProtonMail/gopenpgp/v3/crypto"
	"github.com/ProtonMail/gopenpgp/v3/profile"
)

const (
	publicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEZvGFihYJKwYBBAHaRw8BAQdAsWxDDWjxjmzGn4VOIs3Tdn4BdTIWV/IZowIS
qkOt+W20JFBpb3RyIEtvd2Fsc2tpIDxwaW90ci45OEBpY2xvdWQuY29tPoiTBBMW
CgA7FiEEAtqfe58GdE7ZN26ZtOU23Ke7hasFAmbxhYoCGwMFCwkIBwICIgIGFQoJ
CAsCBBYCAwECHgcCF4AACgkQtOU23Ke7has8RAEA5hf+ijcFRvcLTmM0k+PyrjV2
aTKvSP6fpcpmA18Rb2YA/Ai9tL+q4IbZ0DZ4ycRC+uLMKZuPaC+D2tFCjs+hlGkM
uDgEZvGFihIKKwYBBAGXVQEFAQEHQFp1iim404ZxpNiGV3ieK9mJYMa32wvDbGOl
XMPnFjRVAwEIB4h4BBgWCgAgFiEEAtqfe58GdE7ZN26ZtOU23Ke7hasFAmbxhYoC
GwwACgkQtOU23Ke7hav/tgD/RHO0t4yvxAeaaGTbkqyVSTMNGZBn7EiVKTO9typg
bqwA/RPCS4teLwr+G+W/8Ni+o/KbE/oY/0LgQN7xrfHAsN8A
=jpJU
-----END PGP PUBLIC KEY BLOCK-----`
)

func Encrypt(message []byte) ([]byte, error) {
	key, err := crypto.NewKeyFromArmored(publicKey)
	if err != nil {
		return nil, err
	}
	pgp := crypto.PGPWithProfile(profile.RFC9580())
	encHandle, err := pgp.Encryption().Recipient(key).New()
	if err != nil {
		return nil, err
	}
	pgpMessage, err := encHandle.Encrypt([]byte("my message"))
	if err != nil {
		return nil, err
	}
	return pgpMessage.ArmorBytes()
}
