package models


type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
}

type TaskStore struct {
	tasks map[string]Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]Task),
	}
}

func (s *TaskStore) GetTasks() []Task {
    tasks := make([]Task, 0, len(s.tasks))
    for _, task := range s.tasks {
        tasks = append(tasks, task)
    }
    return tasks
}

func (s *TaskStore) GetTask(id string) (Task, bool) {
    task, exists := s.tasks[id]
    return task, exists
}

func (s *TaskStore) CreateTask(task Task) Task {
    s.tasks[task.ID] = task
    return task
}

func (s *TaskStore) UpdateTask(task Task) (Task, bool) {
    if _, exists := s.tasks[task.ID]; !exists {
        return Task{}, false
    }
    s.tasks[task.ID] = task
    return task, true
}

func (s *TaskStore) DeleteTask(id string) bool {
    if _, exists := s.tasks[id]; !exists {
        return false
    }
    delete(s.tasks, id)
    return true
}