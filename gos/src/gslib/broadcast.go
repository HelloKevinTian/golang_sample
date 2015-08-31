package gslib

import (
	"gslib/gen_server"
)

type Broadcast struct {
	channels map[string](map[string]bool)
}

/*
   GenServer Callbacks
*/
func (self *Broadcast) Init(args []interface{}) (err error) {
	self.channels = make(map[string](map[string]bool))
	return nil
}

func (self *Broadcast) HandleCast(args []interface{}) {
	method_name := args[0].(string)
	if method_name == "JoinChannel" {
		self.handleJoinChannel(args[1].(string), args[2].(string))
	} else if method_name == "LeaveChannel" {
		self.handleLeaveChannel(args[1].(string), args[2].(string))
	} else if method_name == "Publish" {
		self.handlePublish(args[1].(*BroadcastMsg))
	}
}

func (self *Broadcast) HandleCall(args []interface{}) interface{} {
	return nil
}

func (self *Broadcast) Terminate(reason string) (err error) {
	self.channels = nil
	return nil
}

/*
   Callback Handlers
*/

func (self *Broadcast) handleJoinChannel(playerId, channel string) {
	if v, ok := self.channels[channel]; ok {
		v[playerId] = true
	} else {
		m := map[string]bool{}
		m[playerId] = true
		self.channels[channel] = m
	}
}

func (self *Broadcast) handleLeaveChannel(playerId, channel string) {
	if v, ok := self.channels[channel]; ok {
		delete(v, playerId)
	}
}

func (self *Broadcast) handlePublish(msg *BroadcastMsg) {
	channel := msg.Channel
	if v, ok := self.channels[channel]; ok {
		for id, _ := range v {
			if _, ok := gen_server.GetGenServer(id); ok {
				gen_server.Cast(id, "broadcast", msg)
			} else {
				delete(self.channels[channel], id)
			}
		}
	}
}
