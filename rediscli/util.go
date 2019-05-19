package rediscli

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"../common"
	"github.com/gotk3/gotk3/gtk"
	"fmt"
)

var ip string
var port string
var pwd string

var c redis.Conn

func ExecCmd(cmdStr string) (interface{},error){
	if c == nil{
		err := checkRedis()
		if err != nil{
			return nil,err
		}
	}
	cmds := strings.Split(cmdStr," ")
	// todo 过滤空
	cmd := cmds[0]
	args := cmds[1:]
	argIns := make([]interface{},0)
	for _,argIn := range args{
		argIns = append(argIns, argIn)
	}
	return c.Do(cmd,argIns...)
}

func checkRedis() error{
	ipEntry := common.ComponentPool["ipEntry"]
	portEntry := common.ComponentPool["portEntry"]
	ip,_ := ipEntry.(*gtk.Entry).GetText()
	port,_ := portEntry.(*gtk.Entry).GetText()
	//pwd := redis.DialPassword(pwd)
	fmt.Println(ip+":"+port)
	conn, err := redis.Dial("tcp", ip+":"+port)
	if err != nil{
		return err
	}
	c = conn
	return nil
}

func defaultRedis() error{
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil{
		return err
	}
	c = conn
	return nil
}


