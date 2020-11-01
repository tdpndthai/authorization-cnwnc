<template>
  <div>
    <md-table md-card>
      <md-table-toolbar>
        <h1 class="md-title">Danh sách người dùng</h1>
        <router-link :to="{ name: 'adduser' }">
          <a-button v-if="isAddUser" type="primary">
            Thêm mới người dùng
          </a-button>
        </router-link>
      </md-table-toolbar>

      <md-table-row>
        <md-table-head md-numeric>ID</md-table-head>
        <md-table-head>Họ tên</md-table-head>
        <md-table-head>Email</md-table-head>
        <md-table-head>Vai trò</md-table-head>
        <md-table-head>Hành động</md-table-head>
      </md-table-row>

      <md-table-row v-for="(item, index) in listuser" :key="index">
        <md-table-cell md-numeric>{{ item.userId }}</md-table-cell>
        <md-table-cell>{{ item.username }}</md-table-cell>
        <md-table-cell>{{ item.email }}</md-table-cell>
        <md-table-cell>{{
          item.isSuperAdmin ? "Quản trị" : "Người dùng"
        }}</md-table-cell>
        <md-table-cell>
          <md-button
            v-if="isDetail"
            class="md-raised"
            @click="showDetail(item.userId)"
            >Chi tiết</md-button
          >
          <router-link
            :to="{ name: 'updateuser', params: { id: item.userId } }"
          >
            <md-button v-if="isEdit" class="md-raised md-primary"
              >Sửa</md-button
            >
          </router-link>

          <md-button
            @click="showConfirmDelete(item.userId)"
            v-if="isDelete"
            class="md-raised md-accent"
            >Xóa</md-button
          >
        </md-table-cell>
      </md-table-row>
    </md-table>
    <a-modal v-model="visible" title="Chi tiết người dùng" @ok="handleOk">
      <p>Tên người dùng: {{ dataDetail.user.username }}</p>
      <p>Email: {{ dataDetail.user.email }}</p>
      <p>
        Vai trò: {{ dataDetail.user.isSuperAdmin ? "Quản trị" : "Người dùng" }}
      </p>
      <div>
        <p>
          Quyền của
          {{ dataDetail.user.isSuperAdmin ? "quản trị" : "người dùng" }}
        </p>
        <div>
          <ul v-for="(item, index) in dataDetail.permissionUser" :key="index">
            <li v-if="item.roleTitle == 'SuperAdmin'">
              <a-tag color="blue"> {{ item.roleTitle }} </a-tag>
            </li>
            <li v-else-if="item.roleTitle == 'Admin'">
              <a-tag color="cyan"> {{ item.roleTitle }} </a-tag>
            </li>
            <li v-else>
              <a-tag color="green"> {{ item.roleTitle }} </a-tag>
            </li>
            <li>Quyền đang có: {{ item.permissionTitle }}</li>
          </ul>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      listuser: [],
      user: {},
      permission: [],
      isAddUser: false,
      isDetail: false,
      isEdit: false,
      isDelete: false,
      dataDetail: {
        user: {},
        permissionUser: [],
      },
      visible: false,
    };
  },
  methods: {
    getAllUser() {
      axios
        .get("user/getall")
        .then((response) => {
          // console.log(response.data);
          if (response.data.status) {
            this.listuser = response.data.alluser;
            this.$notification.success({
              message: response.data.message,
              description:
                "Lấy ra được " + response.data.alluser.length + " bản ghi",
            });
          } else {
            this.$notification.error({
              message: "Thất bại",
              description: response.data.message,
            });
          }
        })
        .catch((e) =>
          this.$notification.warning({
            message: "Thất bại",
            description: "Lỗi hệ thống " + e,
          })
        );
    },
    showDetail(id) {
      axios
        .get("user/getById/" + id)
        .then((response) => {
          //console.log(response);
          this.dataDetail.user = response.data.user;
          this.dataDetail.permissionUser = response.data.permissionUser;
          this.visible = true;
        })
        .catch((e) =>
          this.$notification.warning({
            message: "Thất bại",
            description: "Lỗi hệ thống " + e,
          })
        );
    },
    handleOk() {
      //console.log(e);
      this.visible = false;
      this.dataDetail.user = {};
      this.dataDetail.permissionUser = [];
    },
    showConfirmDelete(userId) {
      let _this = this;
      this.$confirm({
        title: "Bạn có muốn xóa người dùng này ngay bây giờ?",
        content: "Nhấn OK để xóa",
        okText: "Ok",
        okType: "danger",
        cancelText: "Không",
        onOk() {
          axios
            .delete("user/delete/" + userId)
            .then((res) => {
              if (res.data.status) {
                _this.$notification.success({
                  message: "Thành công",
                  description: res.data.message,
                });
                _this.getAllUser();
              } else {
                _this.$notification.error({
                  message: "Thất bại",
                  description: res.data.message,
                });
              }
            })
            .catch((e) => {
              _this.$notification.warning({
                message: "Thất bại",
                description: "Lỗi hệ thống " + e,
              });
            });
        },
        onCancel() {},
      });
    },
  },
  created() {
    this.getAllUser();
    let userStoreage = JSON.parse(localStorage.getItem("user"));
    if (userStoreage.isSuperAdmin == true) {
      this.isAddUser = true;
    }
    let permissionStoreage = JSON.parse(localStorage.getItem("permission"));
    if (permissionStoreage == null) {
      this.$router.push({ name: "login" });
    }
    permissionStoreage.map((e) => {
      switch (e.permissionConstantName) {
        case "Add":
        case "Detail":
          this.isDetail = true;
          break;
        case "Edit":
          this.isEdit = true;
          break;
        case "Delete":
          this.isDelete = true;
          break;
        default:
          break;
      }
    });
    this.user = userStoreage;
  },
  components: {},
};
</script>

<style>
</style>