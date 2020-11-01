<template>
  <div class="wrapper">
    <side-bar>
      <template slot="links">
        <sidebar-link to="/dashboard" name="Trang chủ" icon="ti-panel" />
        <sidebar-link to="/profile" name="Hồ sơ người dùng" icon="ti-user" />
        <!-- <sidebar-link
          to="/role-permission"
          name="Quản lí quyền"
          icon="ti-view-list-alt"
        /> -->
        <sidebar-link
          v-if="isSuperAdmin"
          to="/list-user"
          name="Danh sách người dùng"
          icon="ti-themify-logo"
        />
        <sidebar-link
          v-if="isQLSP || isUser"
          to="/list-product"
          name="Quản lí sản phẩm"
          icon="ti-themify-logo"
        />
        <sidebar-link
          v-if="isQLDH || isUser"
          to="/list-order"
          name="Quản lí đơn hàng"
          icon="ti-themify-logo"
        />
        <sidebar-link
          v-if="isQLXK || isAdmin"
          to="/list-receipt"
          name="Quản lí nhập kho"
          icon="ti-themify-logo"
        />
        <sidebar-link
          v-if="isQLNK || isAdmin"
          to="/list-issue"
          name="Quản lí xuất kho"
          icon="ti-themify-logo"
        />
      </template>
      <mobile-menu>
        <li class="nav-item">
          <a class="nav-link">
            <i class="ti-panel"></i>
            <p>Stats</p>
          </a>
        </li>
        <drop-down
          class="nav-item"
          title="5 Notifications"
          title-classes="nav-link"
          icon="ti-bell"
        >
          <a class="dropdown-item">Notification 1</a>
          <a class="dropdown-item">Notification 2</a>
          <!-- <a class="dropdown-item">Notification 3</a>
          <a class="dropdown-item">Notification 4</a> -->
          <a class="dropdown-item">Another notification</a>
        </drop-down>
        <li class="nav-item">
          <a class="nav-link">
            <i class="ti-settings"></i>
            <p>Settings</p>
          </a>
        </li>
        <li class="divider"></li>
      </mobile-menu>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>

      <dashboard-content @click.native="toggleSidebar"> </dashboard-content>

      <content-footer></content-footer>
    </div>
  </div>
</template>
<style lang="scss">
</style>
<script>
import TopNavbar from "./TopNavbar.vue";
import ContentFooter from "../content/ContentFooter.vue";
import DashboardContent from "../content/Content.vue";
import MobileMenu from "./MobileMenu";
export default {
  components: {
    TopNavbar,
    ContentFooter,
    DashboardContent,
    MobileMenu,
  },
  data() {
    return {
      isSuperAdmin: false,
      isAdmin: false,
      isUser: false,
      isQLSP: false,
      isQLDH: false,
      isQLXK: false,
      isQLNK: false,
    };
  },
  methods: {
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    },
  },
  created() {
    let userStoreage = JSON.parse(localStorage.getItem("user"));

    if (userStoreage == null) {
      this.$router.push({ name: "login" });
    } else if (userStoreage.isSuperAdmin == true) {
      this.isSuperAdmin = true;
    }
    let permissionStoreage = JSON.parse(localStorage.getItem("permission"));
    permissionStoreage.map((e) => {
      switch (e.permissionConstantName) {
        case "QLSP":
          this.isUser = true;
          this.isQLSP = true;
          break;
        case "QLDH":
          this.isUser = true;
          this.isQLSP = true;
          break;
        case "QLXK":
          this.isAdmin = true;
          this.isQLSP = true;
          break;
        case "QLNK":
          this.isAdmin = true;
          this.isQLSP = true;
          break;
        default:
          null;
      }
    });
  },
};
</script>
