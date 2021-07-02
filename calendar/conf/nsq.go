package conf

type NsqCli struct {
	Addresses []string // localhost:4150
}

type NsqServer struct {
	Addr string // localhost:4150
}
