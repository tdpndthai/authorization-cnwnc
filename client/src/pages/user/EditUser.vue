<template>
  <a-card title="Cập nhật người dùng" :bordered="false">
    <div>
      <label>Tên người dùng</label>
      <a-input placeholder="Nhập tên người dùng" v-model="data.username" />
    </div>
    <div>
      <label>Email người dùng</label>
      <a-input placeholder="Nhập email người dùng" v-model="data.email" />
    </div>
    <div>
      <label>Mật khẩu người dùng</label>
      <a-input
        type="password"
        placeholder="Nhập password người dùng"
        v-model="data.password"
      />
    </div>
    <div>
      <p style="margin-bottom: 0px"><label>Vai trò</label></p>
      <!-- <a-checkbox-group :options="checkboxOptions"></a-checkbox-group> -->
      <a-row>
        <!-- <a-col :span="8">
          <div v-for="(role, idxRole) in listRole" :key="idxRole">
            <input
              class="checkbox_cus pointer"
              @click="onChangeCheckRole(role, idxRole)"
              :checked="setCheckRole(role.roleId)"
              :value="role.roleId"
              type="checkbox"
            />
            {{ role.roleTitle }}
            <div v-if="showSuperAdmin || showAdmin || showUser">
              <p style="margin-bottom: 0px">
                <label style="padding-left: 30px">Quyền: </label>
              </p>
              <div
                style="padding-left: 30px"
                v-for="(per, idxPermission) in listPermission"
                :key="idxPermission"
              >
                <input
                  class="checkbox_cus pointer"
                  @click="
                    onChangeCheckPermission(per, role, idxRole, idxPermission)
                  "
                  :checked="setCheckPermission(per, role)"
                  :value="per.permissionId"
                  type="checkbox"
                />{{ per.permissionTitle }}
              </div>
            </div>
          </div>
        </a-col> -->
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(1)"
            :checked="showSuperAdmin"
          />
          SupperAdmin
          <div v-if="showSuperAdmin">
            <label>Quyền: </label>
            <div v-for="(per, idxPer) in listPermission" :key="idxPer">
              <p style="margin-bottom: 0px">
                <input
                  :value="per.permissionId"
                  type="checkbox"
                  class="checkbox_cus pointer"
                  @click="clickCheckPer(per, 1)"
                  :checked="checkPermissionSuperAdmin(per, 1)"
                />{{ per.permissionTitle }}
              </p>
            </div>
          </div>
        </a-col>
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(2)"
            :checked="showAdmin"
          />
          Admin
          <div v-if="showAdmin">
            <label>Quyền: </label>
            <div v-for="(per, idxPer) in listPermission" :key="idxPer">
              <p style="margin-bottom: 0px">
                <input
                  :value="per.permissionId"
                  type="checkbox"
                  class="checkbox_cus pointer"
                  @click="clickCheckPer(per, 2)"
                  :checked="checkPermissionAdmin(per, 2)"
                />{{ per.permissionTitle }}
              </p>
            </div>
          </div>
        </a-col>
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(3)"
            :checked="showUser"
          />
          User
          <div v-if="showUser">
            <label>Quyền: </label>
            <div v-for="(per, idxPer) in listPermission" :key="idxPer">
              <p style="margin-bottom: 0px">
                <input
                  :value="per.permissionId"
                  type="checkbox"
                  class="checkbox_cus pointer"
                  :checked="checkPermissionUser(per, 3)"
                  @click="clickCheckPer(per, 3)"
                />{{ per.permissionTitle }}
              </p>
            </div>
          </div>
        </a-col>
      </a-row>
    </div>
    <a-button type="primary" @click="update"> Cập nhật </a-button>
  </a-card>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      data: {},
      permiss: "Edit",
      listRole: [],
      roleDisplay: [],
      listPermission: [],
      disabled: false,
      showSuperAdmin: false,
      showAdmin: false,
      showUser: false,
    };
  },
  methods: {
    setCheckRole(id) {
      if (this.data.roleUser) {
        let valueCheck = this.data.roleUser.find((item) => item.roleId === id);
        if (!valueCheck) {
          return false;
        } else {
          switch (id) {
            case 1:
              this.showSuperAdmin = true;
              break;
            case 2:
              this.showAdmin = true;
              break;
            case 3:
              this.showUser = true;
              break;
          }
          return true;
        }
      }
    },
    setCheckPermission(per, role) {
      if (this.data.roleUser && per && role) {
        let checkRole = this.data.roleUser.find(
          (item) => item.roleId === role.roleId
        );
        //nếu có vai trò trùng với vai trò đang kiểm tra
        if (checkRole) {
          let checkPermission = this.data.roleUser.find((item) =>
            item.permissions.map((e) => {
              e.permissionId === per.permissionId;
            })
          );
          //nếu có quyền của vai trò đang kiểm tra
          if (checkPermission) {
            return true;
          } else {
            return false;
          }
        }
      }
    },
    checkDisabled(role) {
      let valueCheck = this.data.roleUser.find(
        (item) => item.roleId === role.roleId
      );
      if (!valueCheck) {
        return true;
      } else {
        return false;
      }
    },
    onChangeCheckRole(role, idxRole) {
      //console.log(`Vai trò ${idxRole + 1} được chọn`);
      let res = this.data.roleUser
        .map((item) => {
          return item.roleId;
        })
        .indexOf(role.roleId);
      //ko thấy
      if (res > -1) {
        console.log(`Vai trò ${idxRole + 1} bị xóa`);
        this.data.roleUser.splice(res, 1);
        this.checkDisabled({ ...role });
      } else {
        console.log(`Vai trò ${idxRole + 1} được chọn`);
        this.data.roleUser = [...this.data.roleUser, { ...role }];
        console.log("this.data.roleUser", this.data.roleUser);
        this.checkDisabled({ ...role });
      }
    },
    onChangeCheckPermission(per, role, idxRole, idxPermission) {
      console.log(
        `Vai trò thứ ${idxRole + 1} ,quyền thứ ${idxPermission + 1} được chọn`
      );
    },

    checkPerSupperAdmin(per, roleId) {
      if (this.data.roleUser) {
        //có role này hay không?
        let checkRole = this.data.roleUser.find(
          (item) => item.roleId === roleId
        );
        if (checkRole) {
          this.data.roleUser.map((role) => {
            let checkPer = role.permissions
              .map((e) => {
                return e.permissionId;
              })
              .indexOf(per.permissionId);
            if (checkPer == -1 && role.roleId === roleId) {
              role.permissions = [...role.permissions, { ...per }];
            } else if (role.roleId === roleId) {
              role.permissions.splice(checkPer, 1);
            }
          });
        } else {
          let data = {
            roleId: 1,
            roleTitle: "SuperAdmin",
            permissions: [{ ...per }],
          };
          this.data.roleUser.push(data);
        }
      }
    },
    checkPerAdmin(per, roleId) {
      if (this.data.roleUser) {
        //có role này hay không?
        let checkRole = this.data.roleUser.find(
          (item) => item.roleId === roleId
        );
        if (checkRole) {
          this.data.roleUser.map((role) => {
            let checkPer = role.permissions
              .map((e) => {
                return e.permissionId;
              })
              .indexOf(per.permissionId);
            if (checkPer == -1 && role.roleId === roleId) {
              role.permissions = [...role.permissions, { ...per }];
            } else if (role.roleId === roleId) {
              role.permissions.splice(checkPer, 1);
            }
          });
        } else {
          let data = {
            roleId: 2,
            roleTitle: "Admin",
            permissions: [{ ...per }],
          };
          this.data.roleUser.push(data);
        }
      }
    },

    clickCheckPer(per, roleId) {
      if (this.data.roleUser) {
        //có role này hay không?
        let checkRole = this.data.roleUser.find(
          (item) => item.roleId === roleId
        );
        if (checkRole) {
          this.data.roleUser.map((role) => {
            let checkPer = role.permissions
              .map((e) => {
                return e.permissionId;
              })
              .indexOf(per.permissionId);
            //quyền phải thuôc về người dùng này
            if (checkPer == -1 && role.roleId === roleId) {
              let checkRoleDeleted = this.data.deletePermission
                .map((e) => {
                  return e.roleId;
                })
                .indexOf(roleId);
              //> -1 tức là đã tồn tại và index của nó đang là checkRoleDeleted
              if (checkRoleDeleted > -1) {
                let perIdExist = this.data.deletePermission
                  .map((item) => {
                    return item.permissionId;
                  })
                  .indexOf(per.permissionId);
                if (perIdExist > -1) {
                  //nếu có quyền của vai trò này đã bị xóa mất thì khôi phục lại quyền này
                  role.permissions = [
                    ...role.permissions,
                    { ...this.data.deletePermission[perIdExist] },
                  ];
                  this.data.deletePermission.splice(perIdExist, 1);
                } else {
                  role.permissions = [...role.permissions, { ...per }];
                }
              } else {
                //nếu nó ko tồn tại trong trong deletePermission,thì thêm mới per này
                role.permissions = [...role.permissions, { ...per }];
              }
              //role.permissions = [...role.permissions, { ...per }];
            } else if (role.roleId === roleId) {
              //chỉ cho vào deletePermission nếu nó có rolePermissionId
              if (
                role.permissions[checkPer].rolePermissionId &&
                role.permissions[checkPer].rolePermissionId != ""
              ) {
                this.data.deletePermission = [
                  ...this.data.deletePermission,
                  { ...role.permissions[checkPer] },
                ];
              }
              role.permissions.splice(checkPer, 1);
            }
          });
        } else {
          switch (roleId) {
            case 1:
              var dataSA = {
                roleId: roleId,
                roleTitle: "SuperAdmin",
                permissions: [{ ...per }],
              };
              this.data.roleUser.push(dataSA);
              break;
            case 2:
              var dataA = {
                roleId: roleId,
                roleTitle: "Admin",
                permissions: [{ ...per }],
              };
              this.data.roleUser.push(dataA);
              break;
            case 3:
              var dataU = {
                roleId: 3,
                roleTitle: "User",
                permissions: [{ ...per }],
              };
              this.data.roleUser.push(dataU);
              break;
            default:
              null;
          }
        }
      }
    },

    clickCheckRole(roleId) {
      switch (roleId) {
        case 1:
          this.showSuperAdmin = !this.showSuperAdmin;
          var indexRoleSA = this.data.roleUser
            .map((e) => {
              return e.roleId;
            })
            .indexOf(roleId);
          //trả về index của role trong mảng
          if (indexRoleSA > -1) {
            //xóa role trong mảng roleUser
            this.data.deleteRole = [
              ...this.data.deleteRole,
              { ...this.data.roleUser[indexRoleSA] },
            ];
            this.data.roleUser.splice(indexRoleSA, 1);
          } else {
            //nếu role đang ko có trong mảng roleUser
            let indexRoleInDeleteRole = this.data.deleteRole
              .map((e) => {
                return e.roleId;
              })
              .indexOf(roleId);
            //nếu role đang tồn tại trong delete tức là gốc đã được check=> hồi phục lại cái role này
            if (indexRoleInDeleteRole > -1) {
              this.data.roleUser = [
                ...this.data.roleUser,
                { ...this.data.deleteRole[indexRoleInDeleteRole] },
              ];
              this.data.deleteRole.splice(indexRoleInDeleteRole, 1);
            }
          }
          break;
        case 2:
          this.showAdmin = !this.showAdmin;
          var indexRoleA = this.data.roleUser
            .map((e) => {
              return e.roleId;
            })
            .indexOf(roleId);
          //trả về index của role trong mảng
          if (indexRoleA > -1) {
            //xóa role trong mảng roleUser
            this.data.deleteRole = [
              ...this.data.deleteRole,
              { ...this.data.roleUser[indexRoleA] },
            ];
            this.data.roleUser.splice(indexRoleA, 1);
          } else {
            //nếu role đang ko có trong mảng roleUser
            let indexRoleInDeleteRole = this.data.deleteRole
              .map((e) => {
                return e.roleId;
              })
              .indexOf(roleId);
            //nếu role đang tồn tại trong delete tức là gốc đã được check=> hồi phục lại cái role này
            if (indexRoleInDeleteRole > -1) {
              this.data.roleUser = [
                ...this.data.roleUser,
                { ...this.data.deleteRole[indexRoleInDeleteRole] },
              ];
              this.data.deleteRole.splice(indexRoleInDeleteRole, 1);
            }
          }
          break;
        case 3:
          this.showUser = !this.showUser;
          var indexRoleU = this.data.roleUser
            .map((e) => {
              return e.roleId;
            })
            .indexOf(roleId);
          //trả về index của role trong mảng
          if (indexRoleU > -1) {
            //xóa role trong mảng roleUser
            this.data.deleteRole = [
              ...this.data.deleteRole,
              { ...this.data.roleUser[indexRoleU] },
            ];
            this.data.roleUser.splice(indexRoleU, 1);
          } else {
            //nếu role đang ko có trong mảng roleUser
            let indexRoleInDeleteRole = this.data.deleteRole
              .map((e) => {
                return e.roleId;
              })
              .indexOf(roleId);
            //nếu role đang tồn tại trong delete tức là gốc đã được check=> hồi phục lại cái role này
            if (indexRoleInDeleteRole > -1) {
              this.data.roleUser = [
                ...this.data.roleUser,
                { ...this.data.deleteRole[indexRoleInDeleteRole] },
              ];
              this.data.deleteRole.splice(indexRoleInDeleteRole, 1);
            }
          }
          break;
        default:
          null;
      }
    },

    checkPermissionUser(per, roleId) {
      if (this.data.roleUser) {
        let checkRole = this.data.roleUser
          .map((item) => {
            return item.roleId;
          })
          .indexOf(roleId);
        if (checkRole > -1) {
          if (this.data.roleUser[checkRole].permissions != null) {
            let checkPer = this.data.roleUser[checkRole].permissions
              .map((permis) => {
                return permis.permissionId;
              })
              .indexOf(per.permissionId);
            if (checkPer > -1) {
              return true;
            } else {
              return false;
            }
          } else {
            this.data.roleUser[checkRole].permissions = [];
          }
        }
      }
    },
    checkPermissionSuperAdmin(per, roleId) {
      if (this.data.roleUser) {
        let checkRole = this.data.roleUser
          .map((item) => {
            return item.roleId;
          })
          .indexOf(roleId);
        if (checkRole > -1) {
          if (this.data.roleUser[checkRole].permissions != null) {
            let checkPer = this.data.roleUser[checkRole].permissions
              .map((permis) => {
                return permis.permissionId;
              })
              .indexOf(per.permissionId);
            if (checkPer > -1) {
              return true;
            } else {
              return false;
            }
          } else {
            this.data.roleUser[checkRole].permissions = [];
          }
        }
      }
    },
    checkPermissionAdmin(per, roleId) {
      if (this.data.roleUser) {
        let checkRole = this.data.roleUser
          .map((item) => {
            return item.roleId;
          })
          .indexOf(roleId);
        if (checkRole > -1) {
          if (this.data.roleUser[checkRole].permissions != null) {
            let checkPer = this.data.roleUser[checkRole].permissions
              .map((permis) => {
                return permis.permissionId;
              })
              .indexOf(per.permissionId);
            if (checkPer > -1) {
              return true;
            } else {
              return false;
            }
          } else {
            this.data.roleUser[checkRole].permissions = [];
          }
        }
      }
    },
    update() {
      //console.log(JSON.stringify(this.data));
      var json = JSON.stringify(this.data);
      axios
        .put("user/update", json)
        .then((res) => {
          if (res.data.status) {
            this.$router.push({ name: "listuser" });
            this.$notification.success({
              message: "Thành công",
              description: res.data.message,
            });
          } else {
            this.$notification.warning({
              message: "Thất bại",
              description: res.data.message,
            });
          }
        })
        .catch((e) => {
          this.$notification.warning({
            message: "Thất bại",
            description: "Lỗi hệ thống " + e,
          });
        });
    },
  },
  created() {
    let permissionStoreage = JSON.parse(localStorage.getItem("permission"));
    if (permissionStoreage == null) {
      this.$router.push({ name: "login" });
    }
    let res = permissionStoreage
      .map((e) => {
        return e.permissionConstantName;
      })
      .indexOf(this.permiss);
    if (res == -1) {
      this.$router.push({ name: "notfound" });
    }
    let userStoreage = JSON.parse(localStorage.getItem("user"));
    if (userStoreage.isSuperAdmin == true) {
      this.isSuperAdmin = true;
    } else {
      this.$router.push({ name: "notfound" });
    }
    let id = parseInt(this.$route.params.id);
    axios
      .get("user/getDetail/" + id)
      .then((res) => {
        this.data = res.data.user;
        this.data.deletePermission = [];
        this.data.deleteRole = [];
        this.data.roleUser = res.data.roleUser;
        res.data.roleUser.map((e) => {
          switch (e.roleId) {
            case 1:
              this.showSuperAdmin = true;
              break;
            case 2:
              this.showAdmin = true;
              break;
            case 3:
              this.showUser = true;
              break;
            default:
              null;
          }
        });
        this.listRole = res.data.role;
        this.listPermission = res.data.permission;
      })
      .catch((e) => {
        console.log(e);
        this.$notification.warning({
          message: "Thất bại",
          description: "Lỗi hệ thống " + e,
        });
      });
  },
};
</script>

<style scoped>
.pointer {
  cursor: pointer;
}

.checkbox_cus {
  width: 20px;
  height: 20px;
  margin-right: 10px;
}
</style>