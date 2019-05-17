package rediscli

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

var ip string
var port string
var pwd string

var c redis.Conn

func ExecCmd(cmdStr string) (interface{},error){
	if c == nil{
		checkRedis()
	}
	cmds := strings.Split(cmdStr," ")
	cmd := cmds[0]
	args := cmds[1:]
	return c.Do(cmd,args)
}

func checkRedis() error{
	pwd := redis.DialPassword(pwd)
	conn, err := redis.Dial("tcp", ip+":"+port,pwd)
	if err != nil{
		return err
	}
	c = conn
	return nil
}


