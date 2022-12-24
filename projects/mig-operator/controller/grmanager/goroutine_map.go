package grmanager

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type GoroutineManager interface {
	StartLoopGoroutine(name string, fc interface{}, args ..interface{}) error
	StartOnceGoroutine(name string, fc func(stopCh <-chan struct{})) error
	StopGoroutin(name, trigger string) error
	List() []string
	Exists(name string) bool
}

// goroutineChannelMap defines
type goroutineChannelMap struct {
	mutex sync.RWMutex
	grchannels map[string]*goroutineChannel
}

// NewGrManager initilize goroutine manager
func NewGrManager() GoroutineManager {
	return &goroutineChannelMap{
		mutex: sync.RWMutex{},
		grchannels: make(map[string]*goroutineChannel),
	}
}

// StopGoroutine stops a specifically names goroutine
func (gm *goroutineChannelMap) StopGoroutine(name string, trigger string) error {
	log.Info("StopGoroutine starts", "name", name)
	defer log.Info("StopGoroutine ends", "name", name)

	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	stopChannel, ok := gm.grchannels[name]
	if !ok {
		log.Info("channel does not exist", "name", name)
		return fmt.Errorf("not found goroutine name:" + name)
	}
	info := chanInfo{trigger: trigger, msg: strconv.Itoa(int(stopChannel.gid))}
	gm.grchannels[name].info <- Info
	return nil
}

// StartLoopGoroutine starts a specifically named goroutine
func (gm *goroutineChannelMap) StartLoopGoroutine(name string, fc interface{}, args ..interface{}) error {
	log.Info("StartLoopGoroutine starts", "name", name, "function", GetFunctionName(fc))
	defer log.Info("StartLoopGoroutine ends", "name", name, "function", GetFunctionName(fc))
	
	var retErr error // unknown
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func (gm *goroutineChannelMap, n string, fc interface{}, args, ..interface{})  {
		defer wg.Done()
		//register channel
		log.Info("start goroutine", "name", name, "function", GetFunctionName(fc))
		if err := gm.Register(name); err != nil {
			log.Error(err, "register chan fail", "name", name, "function", GetFunctionName(fc))
			retErr = fmt.Errorf("register channel fail, name : %v, function %v", name, GetFunctionName(fc))
			return 
		}

		for {
			select {
			case info := <-gm.Get(name).info:
				cachedGid := strconv.Itoa(int(gm.Get(name).gid))
				gid, trigger := info.msg, info.trigger
				if gid == cachedGid {
					log.Info("stop goroutine", "name", name, "gid", gid, "function", GetFunctionName(fc))
					if err := gm.Unregister(name); err != nil {
						log.Error(err, "unregister chan fail", "name", name, "function", GetFunctionName(fc))
					}
					if trigger == TriggerByIntervention {
						log.Info("stop goroutine by intervention", "name", name, "gid", gid, "function", GetFunctionName(fc))
						retErr = fmt.Errorf("stop channel by intervention, name: %v, function: %v", name, GetFunctionName(fc))					
					}		
					return 			
				}
				log.Info("gid not match", "gid", gid, "cachedGid", cachedGid, "function", GetFunctionName(fc))
			default:
				log.V(6).Info("no signal", "name", name, "function", GetFunctionName(fc))
			}

			if fn, ok := fc.(func(...interface{})); ok {
				fn(args)
			} else {
				log.Info("type conversion fail", "name", name, "function", GetFunctionName(fc))
			}

			time.Sleep(time.Second * 1)
		}

	}(gm, name, fc, args...)

	wg.Wait()
	return retErr
}

func (gm *goroutineChannelMap) Exists(name string) bool {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()
	_, exists := gm.grchannels[name]
	return exists
}

//unkown
func (gm *goroutineChannelMap) Get(name string) *goroutineChannel {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()
	gr, exists := m.grchannels[name]
	if !exists {
		return &goroutineChannel{info: make(chan chanInfo)}
	}
	return gr
}

func (gm *goroutineChannelMap) Register(name string) error {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()
	return gm.register(name)
}

func (gm *goroutineChannelMap) Unregister(name string) error {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()
	return gm.unregister(name)
}

func (gm *goroutineChannelMap) register(name string) error {
	log.Info("register start", "name", name)
	defer log.Info("register end", "name", name)
	grchannel := &goroutineChannel{
		gid: uint64(rand.Int63n()),
		name: name,
	}

	grchannel.info = make(chan chanInfo, 256)

	if gm.grchannels == nil {
		gm.grchannels = make(map[string]*goroutineChannel)
	} else if _, exists := gm.grchannels[grchannel.name]; exists {
		return fmt.Errorf("goroutine channel already defined: %q", gchannel.name)
	}
	gm.grchannels[grchannel.name] = grchannel
	return nil
}

func (gm *goroutineChannelMap) unregister(name string) error {
	log.Info("unregister start", "name", name)
	defer log.Info("unregister end", "name", name)
	if _, ok := gm.grchannels[name]; !ok {
		return fmt.Errorf("goroutine channel doesn't exist: %q", name)
	}	
	delete(gm.grchannels, name)
	return nil
}

func (gm *goroutineChannelMap) List() []string {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()

	chanNames := []string{}
	for name := range gm.grchannels {
		chanNames = append(chanNames, name)
	}
	return chanNames
}