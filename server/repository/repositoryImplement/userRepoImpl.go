package repositoryImplement

import (
	"admin-go/dataAccess"
	"admin-go/models"
	"admin-go/repository"
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.IUserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}

func (r *UserRepoImpl) Login(loginModel *models.LoginModel) *models.UserModel {

	row := r.DB.Debug().Table("users").Where("email=? and pass_word=?", loginModel.Email, loginModel.Password).Row()
	var userId int
	var username sql.NullString
	var password sql.NullString
	var email sql.NullString
	var isSuperAdmin bool
	err := row.Scan(&userId, &username, &password, &email, &isSuperAdmin)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		user := models.UserModel{
			UserId:       userId,
			UserName:     username.String,
			PassWord:     password.String,
			Email:        email.String,
			IsSuperAdmin: isSuperAdmin,
		}
		return &user
	}
}

func (r *UserRepoImpl) GetAll() []models.UserModel {
	rows, err := r.DB.Debug().Table("users").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data []models.UserModel
	for rows.Next() {
		var userId int
		var username sql.NullString
		var password sql.NullString
		var email sql.NullString
		var isSuperAdmin bool
		err := rows.Scan(&userId, &username, &password, &email, &isSuperAdmin)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			user := models.UserModel{
				UserId:       userId,
				UserName:     username.String,
				PassWord:     password.String,
				Email:        email.String,
				IsSuperAdmin: isSuperAdmin,
			}
			data = append(data, user)
		}
	}
	return data
}

func (r *UserRepoImpl) GetAllPermissionOfUser(userId int) []models.PermissionUserModel {
	rows, err := r.DB.Debug().
		Raw("Select permissions.permission_id,roles.role_id, roles.role_title,permissions.permission_title ,permissions.permission_constant_name ,role_permissions.role_permission_id from permissions INNER JOIN role_permissions ON permissions.permission_id = role_permissions.permission_id INNER JOIN	 roles ON role_permissions.role_id = roles.role_id INNER JOIN user_roles ON roles.role_id = user_roles.role_id INNER JOIN users ON user_roles.user_id = users.user_id Where users.user_id = ? group by permissions.permission_id,roles.role_id, roles.role_title,permissions.permission_title ,permissions.permission_constant_name ,role_permissions.role_permission_id ", userId).Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data []models.PermissionUserModel
	for rows.Next() {
		var permissionId int
		var roleId int
		var roleTitle sql.NullString
		var permissionTitle sql.NullString
		var permissionConstantName sql.NullString
		var rolePermissionId int
		err := rows.Scan(&permissionId, &roleId, &roleTitle, &permissionTitle, &permissionConstantName, &rolePermissionId)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			user := models.PermissionUserModel{
				PermissionId:           permissionId,
				RoleId:                 roleId,
				RoleTitle:              roleTitle.String,
				PermissionTitle:        permissionTitle.String,
				PermissionConstantName: permissionConstantName.String,
				RolePermissionId:       rolePermissionId,
			}
			data = append(data, user)
		}
	}
	return data
}

func (r *UserRepoImpl) GetById(id int) *models.UserModel {
	row := r.DB.Debug().Table("users").Where("user_id = ?", id).Row()
	var userId int
	var username sql.NullString
	var email sql.NullString
	var password sql.NullString
	var isSuperAdmin bool
	err := row.Scan(&userId, &username, &password, &email, &isSuperAdmin)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		user := models.UserModel{
			UserId:       userId,
			UserName:     username.String,
			PassWord:     password.String,
			Email:        email.String,
			IsSuperAdmin: isSuperAdmin,
		}
		return &user
	}
}

func (r *UserRepoImpl) GetAllRoleOfUser(id int) []models.RoleUserModel {
	//var roles dataAccess.Roles
	//row := r.DB.Debug().Preload("UserRoles").Find(&roles)
	// Raw("select UserRoles.RoleId,Roles.RoleTitle,UserRoles.UserRoleId,Permission.PermissionId,Permission.PermissionConstantName,Permission.PermissionTitle from UserRoles inner join Roles on UserRoles.RoleId = Roles.RoleId  inner join RolePermissions on RolePermissions.RoleId = Roles.RoleId inner join Permission on Permission.PermissionId = RolePermissions.PermissionId	 where UserRoles.UserId = ? ", id).
	// Rows()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// var roleId int
	// var roleTitle sql.NullString
	// var userRoleId int
	// var permissionId int
	// var permissionTitle sql.NullString
	// var permissionConstantName sql.NullString
	// var permission []models.PermissionModel
	// row.Scan(roles)

	var data []models.RoleUserModel
	//var dataper models.PermissionModel
	// for rows.Next() {
	// 	err1 := rows.Scan(&roleId, &roleTitle, &userRoleId, &permissionId, &permissionTitle, &permissionConstantName)
	// 	if err1 != nil {
	// 		fmt.Println(err1)
	// 		return nil
	// 	} else {
	// 		permission := models.PermissionModel{
	// 			PermissionId:           permissionId,
	// 			PermissionTitle:        permissionTitle.String,
	// 			PermissionConstantName: permissionConstantName.String,
	// 		}
	// 		//dataper = append(dataper, permission)
	// 		role := models.RoleUserModel{
	// 			RoleId:     roleId,
	// 			RoleTitle:  roleTitle.String,
	// 			UserRoleId: userRoleId,
	// 			Permission: permission,
	// 		}
	// 		data = append(data, role)
	// 		//dataper = make([]models.PermissionModel, 0)
	// 	}
	// }

	//get all roles
	roles, err := r.DB.Debug().Raw("select user_role_id,user_roles.role_id,role_title from user_roles,users,roles where roles.role_id = user_roles.role_id and users.user_id = user_roles.user_id and users.user_id = ?", id).Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var userRoleId int
	var roleId int
	var roleTitle sql.NullString
	for roles.Next() {
		err1 := roles.Scan(&userRoleId, &roleId, &roleTitle)
		if err1 != nil {
			fmt.Println(err1)
			return nil
		} else {
			role := models.RoleUserModel{
				UserRoleId: userRoleId,
				RoleId:     roleId,
				RoleTitle:  roleTitle.String,
			}
			data = append(data, role)
		}
	}
	//get all permission of roles
	for index, role := range data {

		permissions, err3 := r.DB.Debug().Raw("Select role_permission_id,roles.role_id,permissions.permission_id,permission_title,permission_constant_name from roles,role_permissions,permissions where role_permissions.permission_id = permissions.permission_id and role_permissions.role_id = roles.role_id and roles.role_id = ?", role.RoleId).Rows()
		if err3 != nil {
			fmt.Println(err3)
			return nil
		}
		var rolePermissionId int
		var roleId int
		var permissionId int
		var permissionTitle sql.NullString
		var permissionConstantName sql.NullString
		var dataPermission []models.PermissionUserModel
		for permissions.Next() {
			err4 := permissions.Scan(&rolePermissionId, &roleId, &permissionId, &permissionTitle, &permissionConstantName)
			if err4 != nil {
				fmt.Println(err4)
				return nil
			} else {
				per := models.PermissionUserModel{
					RolePermissionId:       rolePermissionId,
					RoleId:                 roleId,
					PermissionId:           permissionId,
					PermissionTitle:        permissionTitle.String,
					PermissionConstantName: permissionConstantName.String,
				}
				dataPermission = append(dataPermission, per)
			}
		}
		data[index].Permissions = dataPermission
	}

	return data
}

func (r *UserRepoImpl) GetAllRole() []models.Role {
	rows, err := r.DB.Debug().
		Table("roles").
		Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var roleId int
	var roleTitle sql.NullString
	var data []models.Role
	for rows.Next() {
		err1 := rows.Scan(&roleId, &roleTitle)
		if err1 != nil {
			fmt.Println(err1)
			return nil
		} else {
			role := models.Role{
				RoleId:    roleId,
				RoleTitle: roleTitle.String,
			}
			data = append(data, role)
		}
	}
	return data
}

func (r *UserRepoImpl) GetAllPermission() []models.PermissionModel {
	rows, err := r.DB.Debug().
		Table("permissions").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data []models.PermissionModel
	for rows.Next() {
		var permissionId int
		var permissionTitle sql.NullString
		var permissionConstantName sql.NullString
		err := rows.Scan(&permissionId, &permissionTitle, &permissionConstantName)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			permission := models.PermissionModel{
				PermissionId:           permissionId,
				PermissionTitle:        permissionTitle.String,
				PermissionConstantName: permissionConstantName.String,
			}
			data = append(data, permission)
		}
	}
	return data
}

func (r *UserRepoImpl) UpdateUser(model *models.UpdateUserModel) error {
	err := r.DB.Debug().Raw("update users Set user_name = ?,pass_word = ?,email =? where user_id = ?", model.UserName, model.PassWord, model.Email, model.UserId).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	//lặp qua DeletePermission xóa các permission từng check nhưng đã bỏ
	for _, value := range model.DeletePermission {
		err := r.DB.Debug().Table("role_permissions").Model(&models.RolePermissions{}).Where("role_permission_id = ?", value.RolePermissionId).Delete(&models.RolePermissions{}).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	//lặp qua các deleteRole xóa các role đã từng check nhưng đã bỏ.trong đó xóa cả các permissions của role đã xóa
	for _, roles := range model.DeleteRole {
		//xóa hết các permission con bên trong roles
		for _, permission := range roles.Permissions {
			err := r.DB.Debug().Table("role_permissions").Model(&dataAccess.RolePermission{}).Where("role_permission_id = ?", permission.RolePermissionId).Delete(&models.RolePermissions{}).Error
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		//tiếp tục xóa các role
		err := r.DB.Debug().Table("user_roles").Model(&models.UserRoles{}).Where("user_role_id = ?", roles.UserRoleId).Delete(&models.UserRoles{}).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	//cập nhật role mới
	for _, newRole := range model.RoleUser {
		//role này đã từng tạo rồi
		if newRole.UserRoleId != 0 { //khác 0 tức là từng tạo rồi
			//kiểm tra tiếp đến permission xem có cái nào mới hay không
			for _, permision := range newRole.Permissions {
				if permision.RoleId == 0 { //roleId bằng 0 thì cho phép tạo mới
					// rolespermission := models.RolePermissions{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}
					err := r.DB.Debug().Table("role_permissions").Create(&dataAccess.RolePermission{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}).Error
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
			}
		} else {
			//nếu role này chưa được tạo
			//kiểm tra tiếp đến các permission của role,tất cả đều là mới hết
			for _, permision := range newRole.Permissions {
				//nếu có roleId thì ko làm gì cả,nếu ko có thì tạo mới permission insert 2 bảng table roleId và permissionId
				if permision.RoleId == 0 {
					err := r.DB.Debug().Table("role_permissions").Create(&dataAccess.RolePermission{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}).Error
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
			}
			//tạo mới role
			err := r.DB.Debug().Table("user_roles").Create(&models.UserRoles{UserId: model.UserId, RoleId: newRole.RoleId}).Error
			if err != nil {
				fmt.Println(err)
				return err
			}

		}
	}
	return err
}

func (r *UserRepoImpl) Update(model *models.UpdateUserModel) error {
	//cập nhật role mới
	for _, newRole := range model.RoleUser {
		//role này đã từng tạo rồi
		if newRole.UserRoleId != 0 { //khác 0 tức là từng tạo rồi
			//kiểm tra tiếp đến permission xem có cái nào mới hay không
			// for _, permision := range newRole.Permissions {
			// 	if permision.RoleId == 0 { //roleId bằng 0 thì cho phép tạo mới
			// 		// rolespermission := models.RolePermissions{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}
			// 		err := r.DB.Debug().Table("role_permissions").Create(&dataAccess.RolePermission{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}).Error
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			return err
			// 		}
			// 	}
			// }
		} else {
			//nếu role này chưa được tạo
			//kiểm tra tiếp đến các permission của role,tất cả đều là mới hết
			// for _, permision := range newRole.Permissions {
			// 	//nếu có roleId thì ko làm gì cả,nếu ko có thì tạo mới permission insert 2 bảng table roleId và permissionId
			// 	if permision.RoleId == 0 {
			// 		err := r.DB.Debug().Table("role_permissions").Create(&dataAccess.RolePermission{RoleId: newRole.RoleId, PermissionId: permision.PermissionId}).Error
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			return err
			// 		}
			// 	}
			// }
			//tạo mới role
			err := r.DB.Debug().Table("user_roles").Create(&models.UserRoles{UserId: model.UserId, RoleId: newRole.RoleId}).Error
			if err != nil {
				fmt.Println(err)
				return err
			}

		}
	}
	return nil
}

func (r *UserRepoImpl) CreateUser(model *models.CreateUserModel) error {
	//tạo mới người dùng
	user := dataAccess.User{UserName: model.UserName, Email: model.Email, PassWord: model.PassWord}
	err := r.DB.Debug().Table("users").Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	//tạo mới role,insert vào role_users id người dùng vừa tạo và roleId truyền xuống
	err = r.DB.Debug().Table("user_roles").Create(&models.UserRoles{UserId: user.UserId, RoleId: model.RoleId}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (r *UserRepoImpl) DeleteUser(userId int) error {
	err := r.DB.Debug().Table("users").Where("user_id = ?", userId).Delete(&dataAccess.User{}).Error
	return err
}
