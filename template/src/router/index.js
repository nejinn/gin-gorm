import Vue from "vue";
import VueRouter from "vue-router";
const Login = () => import("../views/Login/Login.vue");
const TheContainer = () => import("../views/Layout/TheContainer.vue");
const Home = () => import("../views/Home/Home.vue");

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "Login",
    component: Login
  },
  {
    path: "/",
    name: "Index",
    component: TheContainer,
    redirect: "/dashboard",
    meta: {
      login: true,
      title: "首页"
    },
    children: [
      {
        path: "dashboard",
        name: "Home",
        component: Home,
        meta: {
          title: "仪表盘",
          login: true
        }
      }
    ]
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
