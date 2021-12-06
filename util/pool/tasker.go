package pool

type Task func() (data interface{}, err error)

type Tasker interface {
	GetTotal() int
	GetTask(index int) Task
}

type tasker struct {
	total  int
	getter func(index int) Task
}

func newTasker(total int, getter func(index int) Task) Tasker {
	return &tasker{
		total:  total,
		getter: getter,
	}
}

func (t *tasker) GetTotal() int {
	return t.total
}

func (t *tasker) GetTask(index int) Task {
	return t.getter(index)
}
