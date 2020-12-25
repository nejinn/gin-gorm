<template>
  <nly-navbar variant="white" header border size="sm" :dark="false">
    <nly-navbar-nav>
      <nly-nav-item v-nly-sidebar-collapse.sidebar-collapse>
        <nly-icon icon="nlyfont nly-icon-logo-windows" />
      </nly-nav-item>
      <nly-row align-v="center">
        <nly-col class="d-none d-md-block">
          <nly-breadcrumb
            :item="breadcrumbArray.breadcrumbItems"
            breadcrumb-class="align-items-center p-0 ml-2 mr-0 mt-0 mb-0 bg-transparent"
          />
        </nly-col>
      </nly-row>
    </nly-navbar-nav>
    <nly-navbar-nav class="ml-auto">
      <nly-nav-dropdown
        :popup="true"
        menu-class="border-0"
        size="lg"
        menu-direction="right"
      >
        <template slot="linkcontent">
          <nly-icon icon="nlyfont nly-icon-notification-fill" />
        </template>
        <template slot="menucontent">
          <nly-nav-item :nav-item="false" dropdown-item to="/"
            >我是nav-item="false"</nly-nav-item
          >
          <nly-nav-item :nav-item="false" dropdown-item to="nav"
            >dropdown-item</nly-nav-item
          >
        </template>
      </nly-nav-dropdown>

      <nly-nav-dropdown
        :popup="true"
        menu-class="border-0"
        size="lg"
        menu-direction="right"
      >
        <template slot="linkcontent">
          <nly-icon icon="nlyfont nly-icon-mail-fill" />
          <nly-badge bg-variant="fuchsia" badge-class="navbar-badge"
            >12</nly-badge
          >
        </template>
        <template slot="menucontent">
          <nly-nav-item :nav-item="false" dropdown-item to="/"
            >我是nav-item="false"</nly-nav-item
          >
          <nly-nav-item :nav-item="false" dropdown-item to="nav"
            >dropdown-item</nly-nav-item
          >
        </template>
      </nly-nav-dropdown>

      <nly-nav-dropdown
        :popup="true"
        menu-class="border-0"
        size="lg"
        menu-direction="right"
        class="user-menu"
        v-if="userInfo"
      >
        <template slot="linkcontent">
          <img
            :src="userInfo.user_pic"
            class="user-image img-circle elevation-2"
          />
          <span class="d-none d-md-inline">{{ userInfo.username }}</span>
        </template>
        <template slot="menucontent">
          <li class="user-header bg-primary" v-if="userInfo">
            <img
              :src="userInfo.user_pic"
              class="img-circle elevation-2"
              alt="User Image"
            />

            <p>
              {{ userInfo.username }} - Web Developer
              <small>Member since {{ userInfo.create_date }}</small>
            </p>
          </li>
          <li class="user-body">
            <div class="row">
              <div class="col-4 text-center">
                <a href="#">Followers</a>
              </div>
              <div class="col-4 text-center">
                <a href="#">Sales</a>
              </div>
              <div class="col-4 text-center">
                <a href="#">Friends</a>
              </div>
            </div>
            <!-- /.row -->
          </li>
          <li class="user-footer">
            <nly-button
              flat
              bg-gradient-variant="warning"
              @click="logout"
              button-class="btn-flat"
              >修改资料</nly-button
            >
            <nly-button
              flat
              button-class="float-right btn-flat"
              bg-gradient-variant="danger"
              @click="logout"
              >退出登录</nly-button
            >
          </li>
        </template>
      </nly-nav-dropdown>

      <nly-nav-item>
        <nly-icon icon="nlyfont nly-icon-setting-fill" />
      </nly-nav-item>
    </nly-navbar-nav>
  </nly-navbar>
</template>

<script>
export default {
  name: "TheHeader",
  props: {
    userInfo: {
      type: [Object],
      default: null
    }
  },
  methods: {
    logout() {
      this.$router.push({
        name: "Login"
      });
      this.$store.commit("clearLoginUserInfo");
    }
  },
  computed: {
    breadcrumbArray() {
      const routerArray = [];
      this.$route.matched.forEach(item => {
        if (item.name == this.$route.name) {
          routerArray.push({
            text: item.meta.title,
            active: true,
            to: item.path,
            linkClass: "text-orange"
          });
        } else {
          if (item.path == "") {
            if (this.$route.name !== "Home") {
              routerArray.push({
                text: item.meta.title,
                to: "/",
                linkClass: "text-info"
              });
            }
          } else {
            routerArray.push({
              text: item.meta.title,
              to: item.path,
              linkClass: "text-info"
            });
          }
        }
      });

      return {
        breadcrumbItems: routerArray,
        currentName: this.$route.matched.slice(-1)[0].meta.title
      };
    }
  }
};
</script>
