package conf

type Redis struct {
	Network      string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Expire       int
}
