<template>
  <nly-sidebar-menu
    :sidebarList="sidebarList"
    exact
    side-mini
    container-hover
    brand-to="/"
    :brand-img-src="logo"
    brand-img-circle
    brand-img-elevation
    brand-text="NLY Adminlte Vue"
  />
</template>

<script>
export default {
  name: "TheSidebar",
  props: {
    sidebar: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      logo: undefined
    };
  },
  created() {
    this.logo = require("../../assets/static/img/NLYLOGO.png");
  },
  computed: {
    sidebarList() {
      const sidebarVnode = [
        {
          _type: "nly-sidebar-nav",
          _key: 0,
          // flat: true,
          dataGroup: "one",
          compact: true,
          size: "sm",
          exact: true,
          _children: this.sidebar,
          sidebarNavClass: "mt-2"
        }
      ];
      return sidebarVnode;
    }
  }
};
</script>
