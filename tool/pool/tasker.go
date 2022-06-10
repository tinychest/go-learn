package pool

type Task func() (data interface{}, err error)

type Result struct {
	res interface{}
	err error
}

type Tasker interface {
	Total() int
	Task(index int) Task
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

func (t *tasker) Total() int {
	return t.total
}

func (t *tasker) Task(index int) Task {
	return t.getter(index)
}
