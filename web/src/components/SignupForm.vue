<template>
  <ValidationObserver ref="observer" v-slot="{ passes }">
    <b-form @submit.prevent="passes(onSubmit);" @reset="resetForm">
      <BTextInputWithValidation
        rules="required|email"
        type="email"
        label="Email address:"
        name="Email"
        v-model="email"
        description="We'll never share your email with anyone else"
        placeholder="Enter email"
      />

      <BTextInputWithValidation
        rules="required"
        name="Password"
        vid="password"
        type="password"
        label="Password"
        v-model="password"
        description="We'll never share your password with anyone else"
        placeholder="Enter password"
      />

      <BTextInputWithValidation
        rules="required|confirmed:password"
        name="Password confirmation"
        type="password"
        label="Password confirmation"
        v-model="confirmation"
        description="We'll never share your password with anyone else"
        placeholder="Confirm password"
      />
      <b-button type="submit" variant="primary">Submit</b-button>
      <b-button type="reset" variant="danger">Reset</b-button>
    </b-form>
  </ValidationObserver>
</template>

<script>
import { ValidationObserver } from "vee-validate";
import BTextInputWithValidation from "./inputs/BTextInputWithValidation";


export default {
  name: "BootstrapForm",
  components: {
    ValidationObserver,
    BTextInputWithValidation
  },
  data: () => ({
    email: "",
    password: "",
    confirmation: "",
  }),
  watch: {
    subject(val) {
      console.log(val);
    }
  },
  methods: {
    onSubmit() {
      console.log("Submitted");
    },
    resetForm() {
      this.email = "";
      this.password = "";
      this.confirmation = "";
      requestAnimationFrame(() => {
        this.$refs.observer.reset();
      });
    }
  }
};
</script>
