package server

import (
	"github.com/dollarkillerx/go-proxy-manager/proto/common"

	"sync"
)

type cache struct {
	mu        sync.Mutex
	taskCache map[string]*common.TaskInfo
}

func newCache() *cache {
	return &cache{taskCache: map[string]*common.TaskInfo{}}
}

func (c *cache) AddOrUpdateTask(task *common.TaskInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.taskCache[task.TaskID] = task
}

func (c *cache) GetTaskByDomain(domain string) *common.TaskInfo {
	c.mu.Lock()
	defer c.mu.Unlock()

	req := c.taskCache[domain]
	return req
}
