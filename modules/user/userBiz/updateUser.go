package userBiz

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/hasher"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"

	"go.mongodb.org/mongo-driver/bson"
)

type UpdateUserStore interface {
	UpdateUser(ctx context.Context, conditions interface{}, data interface{}) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type updateBusiness struct {
	store  UpdateUserStore
	hasher hasher.HasherInfo
}

func NewUpdateBusiness(store UpdateUserStore, hasher hasher.HasherInfo) *updateBusiness {
	return &updateBusiness{store: store, hasher: hasher}
}

func (biz *updateBusiness) UpdateUser(ctx context.Context, conditions interface{}, data *userModel.User) error {

	user, err := biz.store.FindUser(ctx, bson.M{"username": data.UserName})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
			return solveError.ErrDB(err)
		}
	}
	if user == nil {
		managerLog.WarningLogger.Println("User's not new")
		return solveError.ErrEntityExisted("User is not exist", nil)
	}

	managerLog.InfoLogger.Println("Check user ok, can create currently user")

	if err := biz.store.UpdateUser(ctx, conditions, data); err != nil {
		return err
	}
	return nil
}
