package internal

import (
	"context"
	"log"
	"runtime"
	"sync"
)

type Manager struct {
	ctx    context.Context
	wg     sync.WaitGroup
	jobsCh chan ObjectInfo
	hasher HashMaker
}

func NewManager(ctx context.Context, path string, hasher HashMaker) (*Manager, error) {
	jobsCh := Walk(path)

	return &Manager{
		ctx:    ctx,
		wg:     sync.WaitGroup{},
		jobsCh: jobsCh,
		hasher: hasher,
	}, nil
}

func (m *Manager) Loop() {
	defer close(m.jobsCh)
	for i := 0; i < runtime.NumCPU(); i++ {
		m.wg.Add(1)
		go m.worker()
	}
	m.wg.Wait()
}

func (m *Manager) worker() {
	for {
		select {
		case <-m.ctx.Done():
			m.wg.Done()
			return
		case object, ok := <-m.jobsCh:
			if ok {
				m.log(object)
			}
		}
	}
}

func (m *Manager) log(obj ObjectInfo) {
	if obj.Err != nil {
		log.Println("error", obj.Err)
		return
	}

	res, err := m.hasher.HashFile(obj.Path)
	if err != nil {
		log.Println("error", obj.Err)
		return
	}

	log.Printf("file: %s, hash: %s", obj.Path, res)
}
