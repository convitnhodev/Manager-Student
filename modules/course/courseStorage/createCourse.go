package courseStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/course/courseModel"
)

func (db *mongoStore) CreateNewCourse(ctx context.Context, data *courseModel.Course) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(courseModel.NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
