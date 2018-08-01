package data

import (
	"logmonitor/base/pub"
	"time"
)

var globalStat map[string]*SrvLogStat

func init() {
	globalStat = map[string]*SrvLogStat{}
}

type SrvLogStat struct {
	srvName              string
	currentLogCount      int
	lastTriggerTimeStamp int64
	logStatCache         *pub.LRUCache
}

func SrvStat(srvname string) *SrvLogStat {
	_, ok := globalStat[srvname]
	if !ok {
		newLogStat := &SrvLogStat{
			srvName:         srvname,
			currentLogCount: 0,
			logStatCache:    pub.NewLRUCache(60),
		}
		globalStat[srvname] = newLogStat
	}
	return globalStat[srvname]
}

func GlobalSrvList() []string {
	retlist := []string{}
	for k, _ := range globalStat {
		retlist = append(retlist, k)
	}
	return retlist
}

func (obj *SrvLogStat) Set() {
	if obj == nil {
		return
	}
	obj.currentLogCount++
}

func (obj *SrvLogStat) StatSet() {
	if obj == nil {
		return
	}
	currentTimeStr := time.Now().Local().Format("15:04")
	obj.logStatCache.Set(currentTimeStr, obj.currentLogCount)
	obj.currentLogCount = 0
}

func (obj *SrvLogStat) GetStatMap() ([]int, error) {
	retlist := []int{}
	StatMap, err := obj.logStatCache.GetAllWithoutUpdate()
	if err != nil {
		return nil, err
	}
	for _, v := range StatMap {
		retlist = append(retlist, v.(int))
	}
	return retlist, nil
}

func (obj *SrvLogStat) LastTriggerTimeStamp() int64 {
	return obj.lastTriggerTimeStamp
}

func (obj *SrvLogStat) SetLastTriggerTimeStamp(v int64) {
	obj.lastTriggerTimeStamp = v
}
