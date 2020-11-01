<template>
  <div>Quản lí nhập kho của Admin</div>
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
        .indexOf("Admin");
      if (res == -1) {
        this.$router.push({ name: "notfound" });
      }
    }
  },
}
</script>

<style>

</style>