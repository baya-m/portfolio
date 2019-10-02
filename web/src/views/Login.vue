<template>
  <div>
    <h1>ログイン</h1>
    <lable for="id">Id:</lable>
    <input id="id" placeholder="id" type="text" v-model="id" />
    <br />
    <label for="password">Password:</label>
    <input id="password" placeholder="password" type="text" v-model="password" />
    <button @click="login">ログイン</button>
    <hr />
    <router-link to="/">もどる</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: "",
      password: "",
    };
  },
    methods: {
    login() {
      this.$store.dispatch('login', {
        id : this.id,
        password: this.password,
      })
    },
  }
};

</script>

<style>
</style>