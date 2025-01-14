package fyf

func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	i := 0
	for i < n {
		i++
		if students[0] == sandwiches[0] {
			if len(students) > 1 {
				students = students[1:]
				sandwiches = sandwiches[1:]
				n = len(students)
				i = 0
			} else {
				return 0
			}
		} else {
			if len(students) > 1 {
				students = append(students[1:], students[0])
			}
			if len(students) == 1 {
				return 1
			}
		}

	}
	return len(students)
}
