package secure

// SignClaims claims for sign
type SignClaims map[string]interface{}

// Signer sign data to make sure data is valid and not changed illegally
type Signer interface {
	// Sign to sign on data and returns the signed data as a string
	Sign(claims SignClaims) (string, error)

	// Validate the sign data and decode data to out interface
	Validate(in string) (SignClaims, error)
}

// IntegerIDHasher encode/decode integer id for secure reason
type IntegerIDHasher interface {
	Encode(n int64) (string, error)
	Decode(s string) (int64, error)
}
