package ured

import (
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/michelia/ulog"
)

type (
	CmdAction = radix.CmdAction
	Cmd       = radix.Cmd
	FlatCmd   = radix.FlatCmd
)

type Config struct {
	Addr        string
	Passwd      string
	DbNum       int
	DialTimeout int // 建立连接, 默认为10s
	PoolSize    int // 连接池大小, 默认为10
}

type Red struct {
	pool *radix.Pool
}

// Do 简单封装了 pool.Do + Cmd
func (r *Red) Do(rcv interface{}, cmd string, args ...string) error {
	return r.pool.Do(radix.Cmd(rcv, cmd, args...))
}

// DoFlat 简单封装了 pool.Do + FlatCmd
func (r *Red) DoFlat(rcv interface{}, cmd, key string, args ...interface{}) error {
	return r.pool.Do(radix.FlatCmd(rcv, cmd, key, args...))
}

// DoPipeline 简单封装了 pool.Do + DoPipeline +
func (r *Red) DoPipeline(cmds ...CmdAction) error {
	return r.pool.Do(radix.Pipeline(cmds...))
}

func New(slog ulog.Logger, c Config) *Red {
	dialTimeout := 10
	if c.DialTimeout != 0 {
		dialTimeout = c.DialTimeout
	}
	poolSize := 10
	if c.PoolSize != 0 {
		poolSize = c.PoolSize
	}
	// this is a ConnFunc which will set up a connection which is authenticated
	// and has a 1 minute timeout on all operations
	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(c.DialTimeout*time.Second),
			radix.DialAuthPass(c.Passwd),
			radix.DialSelectDB(c.DbNum),
			radix.DialReadTimeout(time.Second*5),
			radix.DialWriteTimeout(time.Second*5),
		)
	}
next:
	pool, err := radix.NewPool("tcp", c.Addr, c.PoolSize, PoolConnFunc(customConnFunc))
	if err != nil {
		slog.Error().Caller().Err(err).Msg("connect redis err, wait 3 second and reconnect")
		time.Sleep(time.Second * 3)
		goto next
	}
	return &Red{pool}
}
