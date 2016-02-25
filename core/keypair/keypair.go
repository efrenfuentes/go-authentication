// Package keypair generates private and public RSA keys
package keypair

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GenerateKeys generates private and public RSA keys
func GenerateKeys(bits int) (string, string) {
	priv, _ := rsa.GenerateKey(rand.Reader, bits)

	// Get der format. priv_der []byte
	privDer := x509.MarshalPKCS1PrivateKey(priv)

	privBlk := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDer,
	}

	// Resultant private key in PEM format.
	privPem := string(pem.EncodeToMemory(&privBlk))

	// Public Key generation
	pub := priv.PublicKey
	pubDer, _ := x509.MarshalPKIXPublicKey(&pub)

	pubBlk := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubDer,
	}

	// Resultant public key in PEM format.
	pubPem := string(pem.EncodeToMemory(&pubBlk))

	return privPem, pubPem
}
