package main
import "net/rpc"
import "net"
import "log"
import "github.com/amolb89/kvstore/server/serverlib"
import "fmt"

var Kvstore map [string]string

type RPCCall int

type KVError struct{
	Err string
}


func (e *KVError) Error() string {
	return fmt.Sprintf(e.Err)
}

func (c *RPCCall) Put(args *serverlib.Args, reply *string) error {
	_, ok := Kvstore[args.Key]
	if ok {
		return &KVError{"Key already present\n"}
	} else {
		Kvstore[args.Key] = args.Value
	}
	return nil
}

func (c *RPCCall) Get(key *string, reply *string) error {
	_, ok := Kvstore[*key]
	if ok {
		*reply = Kvstore[*key]
	} else {
		return &KVError{"Key not found\n"}
	}
	return nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp","0.0.0.0:42588")
	if err != nil {
		log.Fatal(err)
	}
	Kvstore = make(map[string]string)
	l, e := net.ListenTCP("tcp", add)
	if e != nil {
		log.Fatal("listen error:",e)	
	}
	call := new(RPCCall)
	rpc.Register(call)
	rpc.Accept(l)
}
