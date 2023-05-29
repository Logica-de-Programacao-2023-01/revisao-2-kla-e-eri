package q1

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	for name, student := range studentData1 {
		mergedData[name] = student
	}

	for name, student := range studentData2 {
		if existingStudent, ok := mergedData[name]; ok {
			for subject, grade := range student.Subjects {
				existingStudent.Subjects[subject] = grade
			}
			mergedData[name] = existingStudent
		} else {
			mergedData[name] = student
		}
	}

	return mergedData
}
