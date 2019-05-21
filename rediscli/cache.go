package rediscli

import (
	"os"
	"encoding/json"
	"sync"
	"io/ioutil"
	"../common"
)




var cacheMap = make(map[string]string,0)
var cacheLock = new(sync.RWMutex)

func init(){
	readCacheFile()
}

func checkFile(){
	dirExist,_ := Exists(common.CACHE_PATH)
	if !dirExist{
		os.MkdirAll(common.CACHE_PATH,os.ModePerm)
	}
	fileExist,_ := Exists(common.CACHE_PATH+common.CACHE_FILE)
	if !fileExist{
		os.Create(common.CACHE_PATH+common.CACHE_FILE)
	}

}

func readCacheFile() error{
	checkFile()
	bs,err := ioutil.ReadFile(common.CACHE_PATH+common.CACHE_FILE)
	if err != nil{
		return err
	}
	err = json.Unmarshal(bs,&cacheMap)
	return err
}

func PutCache(key string, value string){
	cacheLock.Lock()
	defer cacheLock.Unlock()
	checkFile()
	f, _ := os.OpenFile(common.CACHE_PATH+common.CACHE_FILE, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	cacheMap[key] = value
	json.Marshal(cacheMap)
	f.WriteString(key)

}

func GetCache(key string)string{
	cacheLock.RLock()
	defer cacheLock.Unlock()
	return cacheMap[key]
}



// 判断文件/文件夹是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}




