package account

import (
	"github.com/LoCCS/bliss/params"
	"github.com/LoCCS/bliss"
	"github.com/LoCCS/bliss/sampler"
	"fmt"
	"hash"
)

const (
	AddressLength = 52 // or 43 if use base64
	BlissVersion  = params.BLISS_B_4
)

type Address struct {
	pk  bliss.PublicKey
	txt [AddressLength]byte
}

func (addr *Address) ToString() string { return string(addr.txt[:]) }

func newSeed(str string) []uint8 {
	str_len := uint32(len(str))
	var seed_size uint32
	if str_len < sampler.SHA_512_DIGEST_LENGTH {
		seed_size = sampler.SHA_512_DIGEST_LENGTH
	} else {
		seed_size = str_len
	}

	seed := make([]uint8, seed_size)
	copy(seed[:str_len], ([]uint8)(str))
	return seed
}

func newPrivateKey(str string) (*sampler.Entropy, *bliss.PrivateKey, error) {
	seed := newSeed(str)
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		return nil, nil, err
	}

	sk, err := bliss.GeneratePrivateKey(BlissVersion, entropy)
	if err != nil {
		return nil, nil, err
	} else {
		return entropy, sk, nil
	}
}

func GenerateAddress(passphrase string) (*Address, error) {
	_, sk, err := newPrivateKey(passphrase)
	if err != nil {
		return nil, fmt.Errorf("Error: bad passphrase.")
	}

	pk := sk.PublicKey()
	tmp_txt := ([]byte)((hash.Sha3_256(pk.Encode())).ToBase32Hex()) // or ToBase64URL
	var txt [AddressLength]byte
	for i := AddressLength - 1; i >= 0; i-- {
		txt[i] = tmp_txt[i]
	}

	return &Address{
		pk:  *pk,
		txt: txt,
	}, nil
}

func (addr *Address) Authentication(passphrase string) (bool, error) {
	_, sk, err := newPrivateKey(passphrase)
	if err != nil {
		return false, fmt.Errorf("Error: bad passphrase.")
	}

	tmp_pk := sk.PublicKey()
	return string((tmp_pk.Encode())[:]) == string((addr.pk.Encode())[:]), nil
}


