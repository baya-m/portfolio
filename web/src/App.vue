<template>
  <div>
    <Header></Header>
    <!-- <router-view name="header"></router-view> -->
    <router-view></router-view>
  </div>
</template>

<script>
import Header from "./views/Header.vue";
export default {
  components: {
    Header
  }
};
</script>

<style>
</style>