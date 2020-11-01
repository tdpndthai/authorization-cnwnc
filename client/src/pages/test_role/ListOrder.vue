<template>
  <div>Quản lí đơn hàng của quyền User</div>
</template>

<script>
export default {
  created() {
    let permissionStoreage = JSON.parse(localStorage.getItem("permission"));
    if (permissionStoreage == null) {
      this.$router.push({ name: "login" });
    } else {
      let res = permissionStoreage
        .map((e) => {
          return e.roleTitle;
        })
        .indexOf("User");
      if (res == -1) {
        this.$router.push({ name: "notfound" });
      }
    }
  },
};
</script>

<style>
</style>