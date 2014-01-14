package main
import "net/rpc"
import "log"
import "fmt"
import "github.com/amolb89/kvstore/server/serverlib"
import "strconv"

func main () {
	client, err := rpc.Dial("tcp", "localhost" + ":42588")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	args := new(serverlib.Args)
	for i := 1; i<20; i+=2 {
		args.Key = strconv.Itoa(i)
		args.Value = strconv.Itoa(i)
		err = client.Call("RPCCall.Put",args, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
	//for i:= 1; i<20; i+=4 {
	//	args.Key = strconv.Itoa(i)
	//	args.Value = strconv.Itoa(i)
	//	err = client.Call("RPCCall.Put",args, &reply)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	for i:=1; i<20; i+=2 {
		key := strconv.Itoa(i)
		err = client.Call("RPCCall.Get",&key,&reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}		
	
}
