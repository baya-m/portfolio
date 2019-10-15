import Vue from "vue";
import Router from "vue-router";

const Home = () => import("./views/Home.vue");
const Header = () => import("./components/Header.vue");
const User = () => import("./views/User.vue");
const Login = () => import("./views/Login.vue");
const Logout = () => import("./views/Logout.vue");
const Signup = () => import("./views/Signup.vue");
const Main = () => import("./views/Main.vue");
const HeaderMain = () => import("./components/HeaderMain.vue");
const SidebarMain = () => import("./components/SidebarMain.vue");


Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      components: {
        default: Home,
        header: Header
      },
      meta: {
        isPublic: true
      }
    },
    {
      path: "/login",
      components: {
        default: Login,
        header: Header,
      },
      meta: {
        isPublic: true
      }
    },
    {
      path: "/logout",
      components: {
        default: Logout,
        header: Header
      }
    },
    {
      path: "/signup",
      components: {
        default: Signup,
        header: Header
      },
      meta: {
        isPublic: true
      }
    },
    {
      path: "/user",
      components: {
        default: User,
        header: Header
      }
    },
    {
      path: "/main",
      components: {
        default: Main,
        header: HeaderMain,
        sidebar: SidebarMain
      },
      meta: {
        isPublic: true
      }
    }
  ]
});



