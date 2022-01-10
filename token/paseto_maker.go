package token

import (
	"fmt"
	"time"

	"github.com/vk-rv/pvx"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *pvx.ProtoV4Local
	symmetricKey *pvx.SymKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d", chacha20poly1305.KeySize)
	}
	symK := pvx.NewSymmetricKey([]byte(symmetricKey), pvx.Version4)
	maker := &PasetoMaker{
		paseto:       pvx.NewPV4Local(),
		symmetricKey: symK,
	}
	return maker, nil
}

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return m.paseto.Encrypt(m.symmetricKey, payload)
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey).ScanClaims(&payload)
	if err != nil {

		if err == ErrExpiredToken {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
