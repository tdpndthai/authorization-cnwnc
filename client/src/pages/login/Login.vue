<template>
  <div class="main">
    <p class="sign" align="center">Đăng nhập hệ thống</p>
    <div class="form1">
      <input v-model="user.email" class="un" type="text" />
      <input v-model="user.password" class="pass" type="password" />
      <button @click="loginHandler" class="submit" align="center">
        Đăng nhập
      </button>
      <p class="forgot" align="center"><a href="#">Quên mật khẩu?</a></p>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      user: {
        email: "nguyendanhthai1605@gmail.com",
        password: "1605",
      },
    };
  },
  methods: {
    loginHandler() {
      axios
        .post("user/login", this.user)
        .then((response) => {
          //console.log(response.data);
          let permission = [];
          if (response.data.status) {
            if (
              response.data.permissionUser &&
              response.data.permissionUser.length > 0
            ) {
              permission = response.data.permissionUser;
            }
            // response.data.permissionUser &&
            //   response.data.permissionUser.length > 0 &&
            //   response.data.permissionUser.map((e) => {
            //     let res = permission.indexOf(e.permissionConstantName);
            //     if (res == -1) {
            //       permission = [...permission, e];
            //     }
            //   });
            localStorage.setItem("user", JSON.stringify(response.data.user));
            localStorage.setItem("permission", JSON.stringify(permission));
            this.$notification.success({
              message: "Thành công",
              description: response.data.message + response.data.user.username,
            });
            this.$router.push({ name: "dashboard" });
          } else {
            this.$notification.warning({
              message: "Thất bại",
              description: "Lỗi " + response.data.message,
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
  },
};
</script>

<style scoped>
body {
  background-color: #f3ebf6;
  font-family: "Ubuntu", sans-serif;
}
.main {
  background-color: #ffffff;
  width: 400px;
  height: 400px;
  margin: 7em auto;
  border-radius: 1.5em;
  box-shadow: 0px 11px 35px 2px rgba(0, 0, 0, 0.14);
}

.sign {
  padding-top: 40px;
  color: #8c55aa;
  font-family: "Ubuntu", sans-serif;
  font-weight: bold;
  font-size: 23px;
}

.un {
  width: 76%;
  color: rgb(38, 50, 56);
  font-weight: 700;
  font-size: 14px;
  letter-spacing: 1px;
  background: rgba(136, 126, 126, 0.04);
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  outline: none;
  box-sizing: border-box;
  border: 2px solid rgba(0, 0, 0, 0.02);
  margin-bottom: 50px;
  margin-left: 46px;
  text-align: center;
  margin-bottom: 27px;
  font-family: "Ubuntu", sans-serif;
}

form.form1 {
  padding-top: 40px;
}

.pass {
  width: 76%;
  color: rgb(38, 50, 56);
  font-weight: 700;
  font-size: 14px;
  letter-spacing: 1px;
  background: rgba(136, 126, 126, 0.04);
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  outline: none;
  box-sizing: border-box;
  border: 2px solid rgba(0, 0, 0, 0.02);
  margin-bottom: 50px;
  margin-left: 46px;
  text-align: center;
  margin-bottom: 27px;
  font-family: "Ubuntu", sans-serif;
}

.un:focus,
.pass:focus {
  border: 2px solid rgba(0, 0, 0, 0.18) !important;
}

.submit {
  cursor: pointer;
  border-radius: 5em;
  color: #fff;
  background: linear-gradient(to right, #9c27b0, #4093fb);
  border: 0;
  padding-left: 40px;
  padding-right: 40px;
  padding-bottom: 10px;
  padding-top: 10px;
  font-family: "Ubuntu", sans-serif;
  margin-left: 35%;
  font-size: 13px;
  box-shadow: 0 0 20px 1px rgba(0, 0, 0, 0.04);
}

.forgot {
  text-shadow: 0px 0px 3px rgba(117, 117, 117, 0.12);
  color: #e1bee7;
  padding-top: 15px;
}

a {
  text-shadow: 0px 0px 3px rgba(117, 117, 117, 0.12);
  color: #e1bee7;
  text-decoration: none;
}

@media (max-width: 600px) {
  .main {
    border-radius: 0px;
  }
}
</style>