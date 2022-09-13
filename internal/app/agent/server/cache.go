package server

import (
	"github.com/dollarkillerx/go-proxy-manager/proto/agent"

	"sync"
)

type cache struct {
	mu        sync.Mutex
	taskCache map[string]*agent.AddTaskReq
}

func newCache() *cache {
	return &cache{taskCache: map[string]*agent.AddTaskReq{}}
}

func (c *cache) AddOrUpdateTask(task *agent.AddTaskReq) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.taskCache[task.Domain] = task
}

func (c *cache) GetTaskByDomain(domain string) *agent.AddTaskReq {
	c.mu.Lock()
	defer c.mu.Unlock()

	req := c.taskCache[domain]
	return req
}
