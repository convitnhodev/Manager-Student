package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type AddStudentToCourseStore interface {
	CreateNewStudent(ctx context.Context, data *studentModel.Student, location string) error
	FindStudent(ctx context.Context, conditions interface{}, location string) (*studentModel.Student, error)
}

type addStudentToCourseBiz struct {
	store AddStudentToCourseStore
}

func NewAddStudentToCourseBiz(store AddStudentToClassStore) *addStudentToCourseBiz {
	return &addStudentToCourseBiz{store: store}
}

func (biz *addStudentToCourseBiz) AddStudentToCourse(ctx context.Context, data *studentModel.Student) error {
	student, err := biz.store.FindStudent(ctx, bson.M{"id": data.Id}, "student")
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if student == nil {
		managerLog.WarningLogger.Println("Student is not exist")
		return solveError.ErrEntityNotExisted("Student", nil)
	}

	student, err = biz.store.FindStudent(ctx, bson.M{"id": data.Id, "course_id": data.CourseId}, "student_course")
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if student != nil {
		managerLog.WarningLogger.Println("Student exist")
		return solveError.ErrEntityExisted("Student in this course", nil)
	}

	managerLog.InfoLogger.Println("Check student ok, can create currently user")
	if err := biz.store.CreateNewStudent(ctx, data, "student_course"); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create student ok")
	return nil
}
