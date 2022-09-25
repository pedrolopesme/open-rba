package memory

type Memory struct {
	data map[string][]interface{}
}

func NewMemory() *Memory {
	return &Memory{
		data: make(map[string][]interface{}),
	}
}

func (m *Memory) Insert(key string, row interface{}) error {
	if _, contains := m.data[key]; !contains {
		m.data[key] = make([]interface{}, 1)
	}

	m.data[key] = append(m.data[key], row)
	return nil
}
