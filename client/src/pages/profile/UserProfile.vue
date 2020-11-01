<template>
  <div class="row">
    <div class="col-xl-8 col-lg-7 col-md-6">
      <edit-profile-form :user="user"> </edit-profile-form>
    </div>
    <div class="col-xl-4 col-lg-5 col-md-6">
      <user-card :user="user"> </user-card>
    </div>
  </div>
</template>

<script>
import EditProfileForm from "./EditProfileForm";
import UserCard from "./UserCard.vue";
export default {
  data() {
    return {
      user: {
        userId: Number,
        username: String,
        password: String,
        email: String,
        isSuperAdmin: Boolean,
      },
    };
  },
  created() {
    let user = JSON.parse(localStorage.getItem("user"));
    if (user != null) {
      this.user = user;
    }
  },
  components: {
    EditProfileForm,
    UserCard,
  },
};
</script>

<style>
</style>