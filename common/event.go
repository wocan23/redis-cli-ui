package common



func Subscribe(id string,f func()){
	EventPool[id+"_f"] = f
}

func send(id string){
	f := EventPool[id]
	f()
}
