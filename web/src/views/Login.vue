<template>
  <div>
    <b-card-group deck class="my-5">
      <b-card
        bg-variant="dark"
        header="Welcome to My Portfolio!"
        align="center"
        text-variant="white"
        header-border-variant="white"
        style="max-width: 30rem;"
        class="mx-auto"
      >
        <b-card-text>
          <input class="input form-control" type="text" placeholder="test@test.com" v-model="id" />
          <div class="my-2">
            <input class="form-control" type="password" placeholder="password" v-model="password" />
          </div>
        </b-card-text>
        <button class="btn btn-success btn-block">Login</button>
      </b-card>
    </b-card-group>
    <div class="text-center mt-n4">
      <router-link to="signup">Create New Account</router-link>
    </div>
  </div>

  <!-- 
  <form action>
    <h3 class="card-header bg-transparent font-weight-bold">Login</h3>
    <div class="card-body mx-auto col-sm-8 col-xs-12">
      <div class="form-label-group my-2">
        <label class="label">ID</label>
        <input class="input form-control" type="text" placeholder="test@test.com" v-model="id" />
      </div>
      <div class="form-label-group my-2">
        <label class="label">PASSWORD</label>
        <input class="form-control " type="password" placeholder="password" v-model="password" />
        <div class="text-right">
        </div>
      </div>
      <div class="card-footer bg-transparent text-center">
        <button class="btn btn-outline-success btn-lg">Login</button>
        <br/>
        <button class="btn btn-success btn-lg my-2">アカウント登録</button>
      </div>
      
    </div>
  </form>-->
</template>

<script>
export default {
  data() {
    return {
      id: "",
      password: ""
    };
  },
  methods: {
    login() {
      this.$store.dispatch("login", {
        id: this.id,
        password: this.password
      });
    }
  }
};
</script>

<style>
</style>