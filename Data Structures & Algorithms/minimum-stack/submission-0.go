type MinStack struct {
	arr []int
}

func Constructor() MinStack {
	return MinStack{arr: []int{}}
}

func (s *MinStack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *MinStack) Pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

func (s *MinStack) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.arr) == 0 {
		return 0
	}

	m := s.arr[0]

	for i := 1; i < len(s.arr); i++ {
		m = min(m, s.arr[i])
	}

	return m
}
