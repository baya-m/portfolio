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
        <b-card-header class="text-left bg-transparent">Input your ID & password</b-card-header>
        <ValidationObserver ref="observer" v-slot="{ passes }">
        <b-card-text class="text-left my-2">
          <div class="my-2">
            <label>Username</label>
            <input class="input form-control" type="text" placeholder="username" v-model="username" />
          </div>
          <SignUpForm/>
          <div class="my-2">
            <input
              class="form-control"
              name="password"
              type="password"
              placeholder="password"
              v-model="password"
            />
          </div>
          <div class="my-2">
            <label>Confirm Password</label>
            <input class="input form-control" type="password" placeholder="confirm" v-model="email" />
          </div>
        </b-card-text>
        <button class="btn btn-success btn-block">Signup</button>
        </ValidationObserver>
      </b-card>
    </b-card-group>
  </div>
</template>

<script>
import axios from "axios";
import { ValidationObserver } from "vee-validate";
import SignupForm from "../components/SignupForm";

export default {
  components: {
    ValidationObserver,
    BTextInputWithValidation,
  },
  data() {
    return {
      username: "",
      email: "",
      password: ""
    };
  },
  methods: {
    signup() {
      axios.post("/api/signup", {
        username: this.username,
        email: this.id,
        password: this.password
      });
    }
  }
};
</script>

<style>
</style>