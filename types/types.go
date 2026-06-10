package types

type Job struct {
	Host string
	Port int
}

type Result struct {
	Host    string
	Port    int
	Open    bool
	Service string
	Banner  string
}
