<template>
  <card class="card-user">
    <div slot="image">
      <img src="@/assets/img/background.jpg" alt="..." />
    </div>
    <div>
      <div class="author">
        <img
          class="avatar border-white"
          src="@/assets/img/faces/face-2.jpg"
          alt="..."
        />
        <h4 class="title">
          {{ userData.username }}
          <br />
          <a href="#">
            <small>@{{ userData.email }}</small>
          </a>
        </h4>
      </div>
      <p class="description text-center">
        "I like the way you work it
        <br />
        No diggity <br />
        I wanna bag it up"
      </p>
    </div>
    <hr />
    <div class="text-center">
      <div class="row"></div>
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
};
</script>

<style>
</style>