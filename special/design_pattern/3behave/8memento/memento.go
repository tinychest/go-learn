package _memento

type InputText struct {
	content string
}

func (in *InputText) Append(content string) {
	in.content += content
}

func (in *InputText) GetText() string {
	return in.content
}

func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

// Snapshot 快照，用于存储数据快照
// 对于快照来说，指能不能被外部（不同包）修改，只能获取数据，满足封装的特性
type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}