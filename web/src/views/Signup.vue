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
        <b-card-text class="my-2">
          <SignupForm />
        </b-card-text>
      </b-card>
    </b-card-group>
  </div>
</template>

<script>
import axios from "axios";
import SignupForm from "../components/SignupForm";

export default {
  components: {
    SignupForm
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