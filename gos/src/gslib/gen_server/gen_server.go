package gen_server

import (
	"errors"
	"fmt"
	"gslib/utils"
)

type Packet struct {
	method string
	args   []interface{}
}

type SignPacket struct {
	signal int
	reason string
}

type GenServer struct {
	name         string
	callback     GenServerBehavior
	cast_channel chan []interface{}
	call_channel chan []interface{}
	sign_channel chan SignPacket
}

var SIGN_STOP int = 1

type GenServerBehavior interface {
	Init(args []interface{}) (err error)
	HandleCast(args []interface{})
	HandleCall(args []interface{}) interface{}
	Terminate(reason string) (err error)
}

var ServerRegisterMap = utils.NewCMap()

func setGenServer(name string, instance *GenServer) {
	ServerRegisterMap.Set(name, instance)
}

func GetGenServer(name string) (*GenServer, bool) {
	v := ServerRegisterMap.Get(name)
	if v == nil {
		return &GenServer{}, false
	} else {
		return v.(*GenServer), true
	}
}

func delGenServer(name string) {
	ServerRegisterMap.Delete(name)
}

func Start(server_name string, module GenServerBehavior, args ...interface{}) (gen_server *GenServer) {
	gen_server, exists := GetGenServer(server_name)
	if !exists {
		cast_channel := make(chan []interface{}, 1024)
		call_channel := make(chan []interface{})
		sign_channel := make(chan SignPacket)

		gen_server = &GenServer{
			name:         server_name,
			callback:     module,
			cast_channel: cast_channel,
			call_channel: call_channel,
			sign_channel: sign_channel}

		gen_server.callback.Init(args)

		go loop(gen_server) // Enter infinity loop

		setGenServer(server_name, gen_server)
	} else {
		fmt.Println(server_name, " is already exists!")
	}
	return gen_server
}

func Stop(server_name, reason string) {
	if gen_server, exists := GetGenServer(server_name); exists {
		gen_server.sign_channel <- SignPacket{SIGN_STOP, reason}
	} else {
		fmt.Println(server_name, " not found!")
	}
}

func Call(server_name string, args ...interface{}) (result interface{}, err error) {
	if gen_server, exists := GetGenServer(server_name); exists {
		response_channel := make(chan interface{})
		defer func() {
			close(response_channel)
		}()
		args = append(args, response_channel)
		gen_server.call_channel <- args
		result = <-response_channel
	} else {
		fmt.Println(server_name, " not found!")
		err = errors.New("Server not found!")
	}
	return result, err
}

func Cast(server_name string, args ...interface{}) {
	if gen_server, exists := GetGenServer(server_name); exists {
		gen_server.cast_channel <- args
	} else {
		fmt.Println(server_name, " not found!")
	}
}

func loop(gen_server *GenServer) {
	defer func() {
		terminate(gen_server)
	}()

	for {
		select {
		case args, ok := <-gen_server.cast_channel:
			if ok {
				// utils.INFO("handle_cast: ", args)
				gen_server.callback.HandleCast(args)
			}
		case args, ok := <-gen_server.call_channel:
			if ok {
				// utils.INFO("handle_call: ", args)
				size := len(args)
				response_channel := args[size-1]
				result := gen_server.callback.HandleCall(args[0 : size-1])
				response_channel.(chan interface{}) <- result
			}
		case sign_packet, ok := <-gen_server.sign_channel:
			if ok {
				// utils.INFO("handle_sign: ", sign_packet)
				switch sign_packet.signal {
				case SIGN_STOP:
					gen_server.callback.Terminate(sign_packet.reason)
					return
				}
			}
		}
	}
}

func terminate(gen_server *GenServer) {
	close(gen_server.cast_channel)
	close(gen_server.call_channel)
	close(gen_server.sign_channel)
	delGenServer(gen_server.name)
}
