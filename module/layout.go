package module

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	"../common"
	"../rediscli"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"strings"
)

func Layout() *gtk.Box{

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// ip
	ipbox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	ipLabel,_ := gtk.LabelNew("ip")
	ipLabel.SetSizeRequest(common.DEFAULT_COM_WIDTH,common.DEFAULT_COM_HEIGHT)

	ipEntry,_ := gtk.EntryNew()
	ipEntry.SetWidthChars(20)
	ipEntry.SetPlaceholderText("127.0.0.1")
	ipEntry.SetText("127.0.0.1")
	common.ComponentPool["ipEntry"] = ipEntry

	ipbox.Add(ipLabel)
	ipbox.Add(ipEntry)
	ipbox.SetMarginTop(10)
	// port
	portbox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	portLabel,_ := gtk.LabelNew("port")
	portLabel.SetSizeRequest(common.DEFAULT_COM_WIDTH,common.DEFAULT_COM_HEIGHT)

	portEntry,_ := gtk.EntryNew()
	portEntry.SetWidthChars(20)
	portEntry.SetText("6379")
	common.ComponentPool["portEntry"] = portEntry

	portbox.Add(portLabel)
	portbox.Add(portEntry)
	portbox.SetMarginTop(10)

	// cmd
	cmdbox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	cmdLabel,_ := gtk.LabelNew("cmd:")
	cmdLabel.SetSizeRequest(common.DEFAULT_COM_WIDTH,common.DEFAULT_COM_HEIGHT)

	cmdEntry,_ := gtk.EntryNew()
	cmdEntry.SetWidthChars(50)


	cmdbox.Add(cmdLabel)
	cmdbox.Add(cmdEntry)
	cmdbox.SetMarginTop(10)

	// show
	//showbox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,10)
	tag,_ := gtk.TextTagTableNew()
	buf,_ := gtk.TextBufferNew(tag)
	showView,_ := gtk.TextViewNewWithBuffer(buf)
	showView.SetSizeRequest(400,300)
	showView.SetIndent(5)
	showView.SetEditable(false)
	common.ComponentPool["showView"] = showView
	common.ComponentPool["showViewBuf"] = buf

	//showbox.Add(showView)

	box.Add(ipbox)
	box.Add(portbox)
	box.Add(cmdbox)

	adjust,_ := gtk.AdjustmentNew(0,0,0,10,10,10)
	scroll,_ := gtk.ScrolledWindowNew(adjust,nil)
	scroll.SetSizeRequest(400,300)
	scroll.SetBorderWidth(0)
	scroll.Add(showView)
	scroll.SetVExpand(true)
	scroll.SetMarginBottom(10)
	scroll.SetMarginEnd(10)
	scroll.SetMarginStart(10)
	scroll.SetMarginTop(10)
	scroll.SetFocusHAdjustment(adjust)
	screen,_ := scroll.GetScreen()
	screen.ForceFloating()

	box.Add(scroll)


	// 绑定事见
	enterEvent := func(cmdStr string){
		insertKey(cmdStr,buf)
		res,err := rediscli.ExecCmd(cmdStr)
		if err != nil{
			insertValue(err.Error(),buf)
			insertValue("",buf)
			return
		}

		values := make([]string,0)
		values,err1 := redis.Strings(res,err)
		if err1 != nil{
			resStr,_ := redis.String(res,err)
			values = append(values,resStr )
		}


		cmdEntry.SetText("")
		showViewIn := common.ComponentPool["showViewBuf"]
		buf := showViewIn.(*gtk.TextBuffer)


		for _,el := range values{
			insertValue(el,buf)
		}
		insertValue("",buf)

	}
	common.AppWindow.SetFocusChild(cmdEntry)
	cmdEntry.Connect("key_press_event", func(widget gtk.IWidget, event *gdk.Event) {
		eventkey := gdk.EventKey{event}
		v := eventkey.KeyVal()
		if(v == 65293){
			cmdStr,_ := cmdEntry.GetText()
			cmdStr = strings.TrimSpace(cmdStr)
			if(cmdStr == ""){
				return
			}
			// 执行enter
			enterEvent(cmdStr)
			// 记录上一次指令
			if len(common.LastCmdStrs)==0 || common.LastCmdStrs[len(common.LastCmdStrs)-1]!=cmdStr{
				common.LastCmdStrs = append(common.LastCmdStrs, cmdStr)
			}
			common.LastCmdIndex = len(common.LastCmdStrs)-1
			// 跳到最后
			scrollToBottom()
		}
		// 上键
		if(v == 65362){
			fmt.Println("up")
			common.AppWindow.SetFocusChild(cmdEntry)
			cmdEntry.SetText(getLastCmdStr())

		}
		// 下键
		if(v == 65364){
			fmt.Println("down")
			common.AppWindow.SetFocusChild(cmdEntry)
			cmdEntry.SetText(getNextCmdStr())
		}
	},nil)


	return box
}

func insertKey(key string,buf *gtk.TextBuffer){
	//buf.InsertAtCursor("\n")
	buf.InsertAtCursor(">>")
	buf.InsertAtCursor(key)
	buf.InsertAtCursor("\n")
}

func insertValue(value string,buf *gtk.TextBuffer){
	//buf.InsertAtCursor("\n")
	buf.InsertAtCursor("    ")
	buf.InsertAtCursor(value)
	buf.InsertAtCursor("\n")
}


func getLastCmdStr() string{
	common.LastCmdStrLock.Lock()
	defer common.LastCmdStrLock.Unlock()
	if(len(common.LastCmdStrs) == 0){
		return ""
	}
	cmdStr := common.LastCmdStrs[common.LastCmdIndex]
	common.LastCmdIndex--
	if(common.LastCmdIndex<=0){
		common.LastCmdIndex = 0
	}
	return cmdStr
}

func getNextCmdStr() string{
	common.LastCmdStrLock.Lock()
	defer common.LastCmdStrLock.Unlock()
	common.LastCmdIndex++
	if(common.LastCmdIndex>len(common.LastCmdStrs)-1){
		common.LastCmdIndex = len(common.LastCmdStrs)-1
		return ""
	}
	cmdStr := common.LastCmdStrs[common.LastCmdIndex]
	return cmdStr
}

func scrollToBottom(){
	common.ScrollLock.Lock()
	defer common.ScrollLock.Unlock()

	showViewIn := common.ComponentPool["showView"]
	showViewBufIn := common.ComponentPool["showViewBuf"]

	showView := showViewIn.(*gtk.TextView)
	showViewBuf := showViewBufIn.(*gtk.TextBuffer)
	iter := showViewBuf.GetEndIter()
	mark := showViewBuf.CreateMark("bottomMark",iter,true)
	showView.ScrollToMark(mark,0,true,1,0)

}

