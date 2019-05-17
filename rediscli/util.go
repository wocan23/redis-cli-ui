package rediscli

import "github.com/garyburd/redigo/redis"

func ExecCmd(cmd string){

}

func checkRedis() error{
	c, err := redis.Dial("tcp", "localhost:6379")
	redis.Dial()
	if err != nil{

	}

}


