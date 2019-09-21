<template>
  <div>
    <h1>ログイン</h1>
    <lable for="email">Email:</lable>
    <input id="email" placeholder="email" type="text" v-model="email" />
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
      email: "",
      password: ""
    };
  }
};
</script>

<style>
</style>