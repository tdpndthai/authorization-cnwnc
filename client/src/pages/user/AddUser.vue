<template>
  <a-card title="Tạo mới người dùng" :bordered="false">
    <div>
      <label>Tên người dùng</label>
      <a-input placeholder="Nhập tên người dùng" v-model="data.username" />
    </div>
    <div>
      <label>Email người dùng</label>
      <a-input
        type="email"
        placeholder="Nhập email người dùng"
        v-model="data.email"
      />
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
      <a-row>
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(1)"
          />
          SupperAdmin
        </a-col>
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(2)"
          />
          Admin
        </a-col>
        <a-col :span="8">
          <input
            class="checkbox_cus pointer"
            type="checkbox"
            @click="clickCheckRole(3)"
          />
          User
        </a-col>
      </a-row>
    </div>
    <a-button type="primary" @click="add"> Thêm mới </a-button>
  </a-card>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      data: {
        username: "",
        password: "",
        email: "",
        roleId: 0,
      },
      isSuperAdmin: false,
    };
  },
  methods: {
    clickCheckRole(roleId) {
      this.data.roleId = roleId;
    },
    add() {
      //console.log(this.data);
      let json = JSON.stringify(this.data);
      axios
        .post("user/add", json)
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
    let userStoreage = JSON.parse(localStorage.getItem("user"));
    if (userStoreage == null) {
      this.$router.push({ name: "login" });
    }
    if (userStoreage.isSuperAdmin == true) {
      this.isSuperAdmin = true;
    } else {
      this.$router.push({ name: "notfound" });
    }
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