<template>
  <ValidationObserver ref="observer" v-slot="{ passes }">
    <b-form @submit.prevent="passes(onSubmit);" @reset="resetForm">
      <BTextInputWithValidation
        class="text-left"
        rules="required|alpha_num|min:4|max:15"
        type="username"
        label="Username:"
        name="Username"
        v-model="username"
        placeholder="Enter username"
      />
      <BTextInputWithValidation
        class="text-left"
        rules="required|email"
        type="email"
        label="Email address:"
        name="Email"
        v-model="email"
        placeholder="Enter email"
      />
      <BTextInputWithValidation
        class="text-left"
        rules="required|alpha_num|min:4|max:20"
        name="Password"
        vid="password"
        type="password"
        label="Password:"
        v-model="password"
        placeholder="Enter password"
      />
      <BTextInputWithValidation
        class="text-left"
        rules="required|confirmed:password"
        name="Password confirmation"
        type="password"
        label="Password confirmation:"
        v-model="confirmation"
        placeholder="Confirm password"
      />
      <BCheckboxesWithValidation v-model="sex" rules="required">
        <b-form-checkbox value="Man">Man</b-form-checkbox>
        <b-form-checkbox value="Woman">Woman</b-form-checkbox>
      </BCheckboxesWithValidation>

      <b-button type="submit" variant="primary">Submit</b-button>
      <b-button type="reset" variant="danger">Reset</b-button>
    </b-form>
  </ValidationObserver>
</template>

<script>
import axios from "axios";
import { ValidationObserver } from "vee-validate";
import BTextInputWithValidation from "./inputs/BTextInputWithValidation";
import BCheckboxesWithValidation from "./inputs/BCheckboxesWithValidation";

export default {
  name: "BootstrapForm",
  components: {
    ValidationObserver,
    BTextInputWithValidation,
    BCheckboxesWithValidation
  },
  data: () => ({
    email: "",
    password: "",
    confirmation: "",
    sex: []
  }),
  watch: {
    subject(val) {
      console.log(val);
    }
  },
  methods: {
    onSubmit() {
      axios.post("/api/signup", {
        username: this.username,
        email: this.id,
        password: this.password,
        sex: this.sex
      });
    },
    resetForm() {
      this.email = "";
      this.password = "";
      this.confirmation = "";
      this.choices = [];
      requestAnimationFrame(() => {
        this.$refs.observer.reset();
      });
    }
  }
};
</script>
