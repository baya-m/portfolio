<template>
<div>
  <b-card-group deck class="my-5">
    <b-card 
    bg-variant="light" 
    header="Create a new Account !!"
    align="center"
    style="max-width: 30rem;"
    class="mx-auto"
    >
    <b-card-header class="text-left bg-transparent">
      Input your ID & password
    </b-card-header>
    <b-card-text class="text-left my-2">
        <div>
        <input class="input form-control" type="text" placeholder="id" v-model="id" />
        </div>
        <div class="my-2">
        <input class="form-control" name="password" type="password" placeholder="password" v-model="password" />
        </div>
    </b-card-text>
            <button class="btn btn-success btn-block">Signup</button>
    </b-card>
  </b-card-group>

  </div>

</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      id: "",
      password: ""
    };
  },
  methods: {
    signup() {
      axios
        .post("/api/signup", {
          id: this.id,
          password: this.password
        })
    }
  }
};
</script>

<style>
</style>