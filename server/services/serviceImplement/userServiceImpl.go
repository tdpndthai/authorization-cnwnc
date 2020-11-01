package serviceImplement

import (
	"admin-go/data"
	"admin-go/models"
	"admin-go/repository/repositoryImplement"
	"admin-go/services"
)

type UserServiceImpl struct {
}

func NewUserService() services.IUserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) Login(loginModel *models.LoginModel) map[string]interface{} {
	context := data.NewContext()
	defer context.Close()
	c := context.DbCollection("Users")
	userRepo := repositoryImplement.NewUserRepo(c)

	var user *models.UserModel
	user = userRepo.Login(loginModel)

	var resp = map[string]interface{}{"status": false, "message": "logged failed"}
	if user == nil {
		resp = map[string]interface{}{"status": false, "message": "Email or password failed"}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Đăng nhập thành công xin chào  "}
		listPermission := userRepo.GetAllPermissionOfUser(user.UserId)
		resp["permissionUser"] = listPermission
	}

	resp["user"] = user

	return resp
}

func (s *UserServiceImpl) GetAll() map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	listuser := userRepo.GetAll()
	if listuser == nil {
		resp = map[string]interface{}{"status": false, "message": "Có lỗi khi lấy tất cả người dùng"}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Thành công"}
	}
	resp["alluser"] = listuser
	return resp
}

func (s *UserServiceImpl) GetById(id int) map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	user := userRepo.GetById(id)
	if user == nil {
		resp = map[string]interface{}{"status": false, "message": "Người dùng không tồn tại"}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Thành công"}
		listPermission := userRepo.GetAllPermissionOfUser(id)
		resp["permissionUser"] = listPermission
	}
	resp["user"] = user
	return resp
}

func (s *UserServiceImpl) GetDetailUser(id int) map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	user := userRepo.GetById(id)
	if user == nil {
		resp = map[string]interface{}{"status": false, "message": "Người dùng không tồn tại"}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Thành công"}
		// listPermissionOfUser := userRepo.GetAllPermissionOfUser(id)
		listPermission := userRepo.GetAllPermission()
		listRoleOfUser := userRepo.GetAllRoleOfUser(id)
		listRole := userRepo.GetAllRole()
		// resp["permissionUser"] = listPermissionOfUser
		resp["permission"] = listPermission
		resp["roleUser"] = listRoleOfUser
		resp["role"] = listRole
	}
	resp["user"] = user
	return resp
}

func (s *UserServiceImpl) UpdateUser(user *models.UpdateUserModel) map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	err := userRepo.UpdateUser(user)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": err}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Cập nhật người dùng thành công"}
	}
	return resp
}

func (s *UserServiceImpl) CreateUser(user *models.CreateUserModel) map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	err := userRepo.CreateUser(user)
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": err}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Tạo mới người dùng thành công"}
	}
	return resp
}

func (s *UserServiceImpl) DeleteUser(userId int) map[string]interface{} {
	context := data.NewContext()
	c := context.GormSession
	defer context.Close()
	userRepo := repositoryImplement.NewUserRepo(c)
	var resp = map[string]interface{}{"status": false, "message": "getall failed"}
	err := userRepo.DeleteUser(userId)
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": err}
		return resp
	} else {
		resp = map[string]interface{}{"status": true, "message": "Xóa người dùng thành công"}
	}
	return resp
}
