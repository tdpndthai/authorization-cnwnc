
const routes = [
  {
    path: "/login",
    name: "login",
    component: () => import('@/pages/login/Login.vue')
  },
  {
    path: "/",
    component: () => import('@/layout/dashboard/DashboardLayout.vue'),
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: () => import('@/pages/Dashboard.vue')
      },
      {
        path: "/profile",
        name: "profile",
        component: () => import('@/pages/profile/UserProfile.vue')
      },
      {
        path: "/list-user",
        name: "listuser",
        component: () => import('@/pages/user/ListUser.vue'),
      },
      {
        path: "/update-user/:id",
        name: "updateuser",
        component: () => import('@/pages/user/EditUser.vue')
      },
      {
        path: "/add-user",
        name: "adduser",
        component: () => import('@/pages/user/AddUser.vue')
      },
      {
        path: "/list-product",
        name: "listproduct",
        component: () => import('@/pages/test_role/ListProduct.vue')
      },
      {
        path: "/list-order",
        name: "listorder",
        component: () => import('@/pages/test_role/ListOrder.vue')
      },
      {
        path: "/list-receipt",
        name: "listreceipt",
        component: () => import('@/pages/test_role/ListGoodsReceipt.vue')
      },
      {
        path: "/list-issue",
        name: "listissue",
        component: () => import('@/pages/test_role/ListGoodsIssue.vue')
      },
    ]
  },
  { path: "*", name: "notfound", component: () => import('@/pages/NotFoundPage.vue') }
];


export default routes;
