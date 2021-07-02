package conf

type Redis struct {
	Addr         string // 127.0.0.1:6379
	Pwd          string
	Active       int // unit ?
	Idle         int // unit ?
	DialTimeout  int // unit millisecond
	ReadTimeout  int // unit millisecond
	WriteTimeout int // unit millisecond
	IdleTimeout  int // unit second
	Expire       int // unit minute
}
