package bybit

type APIAuth interface {
	Token() (string, string)
}

type bybitAuth struct {
	AccessKey string
	SecretKey string
}

func BybitAuth(ak, sk string) APIAuth {
	return &bybitAuth{
		AccessKey: ak,
		SecretKey: sk,
	}
}

func (auth bybitAuth) Token() (string, string) {
	return auth.AccessKey, auth.SecretKey
}
