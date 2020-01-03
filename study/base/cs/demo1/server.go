
package main
 
import (
	"flag"
	"github.com/astaxie/beego/logs"
	"net"
)
 
var port string;
func init(){
	flag.StringVar(&port,"port","localhost:80",":port")
}
 
func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"server.log","color":true}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(2)
	flag.Parse()
	listener,err:=net.Listen("tcp",port); if err!=nil{
		panic(err)
	}
	for{
		conn,err:=listener.Accept();if err!=nil{
			logs.Debug("err at accept:%v",err)
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	buffer:=make([]byte,1024)
	for{
		n,err:=conn.Read(buffer);if err!=nil{
			return
		}
		logs.Debug(string(buffer[:n]))
	}
}