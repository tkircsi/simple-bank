package token

import (
	"fmt"
	"time"

	"github.com/vk-rv/pvx"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *pvx.ProtoV2Local
	symmetricKey []byte
}

func NewPasetoMaker(simmetricKey string) (Maker, error) {
	if len(simmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:       pvx.NewPV2Local(),
		symmetricKey: []byte(simmetricKey),
	}
	return maker, nil
}

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
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
