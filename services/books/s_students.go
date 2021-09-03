package books

import (
	"time"
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/books"
)

// CreateStudent 新增学生信息
func (m *ManageBooksService) CreateStudent(s books.Students) (err error) {
	s.CreationTime = time.Now()
	err = global.GVA_DB.Create(&s).Error
	return err
}

// GetAssignStudentInfo 查询单个学生信息
func (m *ManageBooksService) GetAssignStudentInfo(s books.Students) (err error, studentInfo *books.Students) {
	var student books.Students
	err = global.GVA_DB.Where("student_id=?", s.ID).First(&student).Error
	return err, &student
}

// GetAllStudentsInfo 查询所有学生列表
func (m ManageBooksService) GetAllStudentsInfo() (err error, studentInfo []books.Students) {
	var students []books.Students
	err = global.GVA_DB.Find(&students).Error
	return err, students
}
