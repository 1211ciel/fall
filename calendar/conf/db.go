package conf

type Db struct {
	Addr          string
	MaxIdleConnes int
	MaxOpenConns  int
	MaxLifTime    int // unit hour
	Debug         bool
}
