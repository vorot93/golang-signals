package signals

import (
	"math/rand"
	"sync"
)

type SignalCB func(data interface{})

type Signal struct {
	m sync.Mutex
	slots map[int64]SignalCB
}

func (self *Signal) Connect(cb SignalCB) int64 {
	self.m.Lock()
	defer self.m.Unlock()
	if self.slots == nil {
		self.slots = map[int64]SignalCB{}
	}
	for {
		var i = rand.Int63()
		var _, exists = self.slots[i]
		if exists {
			continue
		}

		self.slots[i] = cb
		return i
	}
}

func (self *Signal) Disconnect(v int64) bool {
	self.m.Lock()
	defer self.m.Unlock()
	if self.slots == nil {
		self.slots = map[int64]SignalCB{}
	}
	var _, exists = self.slots[v]
	if exists {
		return false
	}
	delete(self.slots, v)
	return true
}

func (self *Signal) Emit(data interface{}) {
	for _, cb := range self.slots {
		cb(data)
	}
}
