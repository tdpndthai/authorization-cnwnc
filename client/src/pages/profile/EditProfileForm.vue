<template>
  <card class="card" title="Edit Profile">
    <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-5">
            <fg-input
              type="text"
              label="Họ và tên"
              placeholder="Họ và tên"
              v-model="userData.username"
            >
            </fg-input>
          </div>
          <div class="col-md-5">
            <fg-input
              type="email"
              label="Email"
              placeholder="Email"
              v-model="userData.email"
            >
            </fg-input>
          </div>
        </div>
        <div class="text-center">
          <p-button type="info" round @click.native.prevent="updateProfile">
            Update Profile
          </p-button>
        </div>
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
</template>

<script>
export default {
  data() {
    return {
      userData: this.user,
    };
  },
  props: {
    user: {
      type: Object,
      required: true,
    },
  },
  methods: {
    updateProfile() {
      alert("Your data: " + JSON.stringify(this.userData));
    },
  },
};
</script>

<style>
</style>