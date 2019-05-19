package rediscli

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"../common"
	"github.com/gotk3/gotk3/gtk"
	"fmt"
	"regexp"
)

var ip string
var port string
var pwd string

var c redis.Conn

func ExecCmd(cmdStr string) (interface{},error){
	cmdStr = strings.TrimSpace(cmdStr)
	if c == nil{
		err := checkRedis()
		if err != nil{
			return nil,err
		}
	}
	reg,_ := regexp.Compile("\\s+")
	cmds := reg.Split(cmdStr,-1)

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


