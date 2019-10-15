import { required, email, confirmed, min, max, alpha_num } from "vee-validate/dist/rules";
import { extend } from "vee-validate";

extend("required", {
  ...required,
  message: "This field is required"
});

extend("email", {
  ...email,
  message: "This field must be a valid email"
});

extend("confirmed", {
  ...confirmed,
  message: "This field confirmation does not match"
});

extend("min", {
  ...min,
  message: "{_field_} must be at least {length} characters"
});

extend("max", {
    ...max,
    message: "{_field_} must be at least {length} characters"
});

extend("alpha_num", {
    ...alpha_num,
    message: "{_field_} field may only contain alpha-num characters"
})