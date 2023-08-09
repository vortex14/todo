package app

import (
	"github.com/google/uuid"
	"sync"
)

type InMemoryTodoRepository struct {
	data map[string]*Task
	once sync.Once
	mu   sync.RWMutex
}

func (r *InMemoryTodoRepository) Init() *InMemoryTodoRepository {
	r.once.Do(func() {
		r.data = make(map[string]*Task)
	})

	return r
}

func (r *InMemoryTodoRepository) GetList() []*Task {
	r.mu.RLock()
	var list []*Task
	for _, v := range r.data {
		list = append(list, v)
	}
	r.mu.RUnlock()
	return list
}

func (r *InMemoryTodoRepository) GetTask(id string) *Task {
	r.mu.RLock()
	task := r.data[id]
	r.mu.RUnlock()
	return task
}

func (r *InMemoryTodoRepository) DeleteItem(id string) error {
	if _, ok := r.data[id]; ok {
		r.mu.Lock()
		delete(r.data, id)
		r.mu.Unlock()
		return nil
	}

	return TaskNotFound
}

func (r *InMemoryTodoRepository) UpdateItem(task *Task) error {
	if _, ok := r.data[task.Id]; ok {
		r.mu.Lock()
		r.data[task.Id] = task
		r.mu.Unlock()
		return nil
	}

	return TaskNotFound
}

func (r *InMemoryTodoRepository) AddItem(task *Task) error {

	_uuid, _ := uuid.NewUUID()
	task.Id = _uuid.String()

	r.data[task.Id] = task
	return nil
}
