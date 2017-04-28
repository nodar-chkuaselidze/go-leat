package leats

type LeatFn struct {
	Name string
	Fn   func(length int, hash, extend, byteFormat string) (string, string, error)
}

var LeatList = []LeatFn{
	LeatFn{"md5", LeatMd5},
}
