package conf

type EtcdCli struct {
	Endpoints   []string
	DialTimeout int64 // unit second
}
