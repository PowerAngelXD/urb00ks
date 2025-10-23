<template>
  <v-alert
    closable
    icon="fa-solid fa-check"
    title="提示"
    text="登录成功！"
    color="green"
    v-model="isShowLoginDone">
  </v-alert>

  <v-alert
    closable
    icon="fa-solid fa-triangle-exclamation"
    title="提示"
    text="登录失败！"
    color="red"
    v-model="isShowLoginFailed">
  </v-alert>

  <v-dialog v-model="loginDialog" width="600" color="blue lighten-2">
    <v-card color="blue lighten-2">
      <v-toolbar color="blue lighten-2">
        <v-container>
          <v-row align="center">
            <v-btn size="small" icon="fa-solid fa-close" @click="loginDialog = false"></v-btn>
            <v-label>请登录</v-label>
          </v-row>
        </v-container>
      </v-toolbar>
      <v-container>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-label>请输入您的账户昵称和密码以登录</v-label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-text-field label="账户名称" density="compact" variant="outlined" class="mt-3" v-model="userName"></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-text-field label="账户密码" density="compact" variant="outlined" class="mt-3" v-model="userPassword"
              :append-inner-icon="isShowPassword ? 'fa-solid fa-eye' : 'fa-solid fa-eye-slash'"
              :type="isShowPassword ? 'text' : 'password'"
              @click:append-inner="isShowPassword = !isShowPassword"></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="10"></v-col>
          <v-col cols="2">
            <v-btn color="white" variant="outlined" @click="login(userName, userPassword)">登录</v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>

  <v-dialog v-model="signupDialog">
    <v-card>

    </v-card>
  </v-dialog>

  <v-navigation-drawer v-model="userBar" temporary location="right" color="blue-accent-2">
    <br></br>
    <v-list-item v-if="isLogIn" :title="userInfo.name"></v-list-item>
    <v-list density="compact" nav v-else>
      <v-list-item prepend-icon="fa-solid fa-user-plus" title="没有帐户？请注册" value="login"></v-list-item>
      <v-list-item prepend-icon="fa-solid fa-user" title="已有账户？请登录" value="login"
        @click="loginDialog = true"></v-list-item>
    </v-list>
    <v-divider></v-divider>

    <v-list v-if="isLogIn" density="compact" nav>
      <v-list-item prepend-icon="fa-solid fa-star" title="我的收藏" value="myFavs"></v-list-item>
      <v-list-item prepend-icon="fa-solid fa-gear" title="设置" value="settings"></v-list-item>
    </v-list>
  </v-navigation-drawer>

  <v-app-bar :elevation="2" color="blue-lighten-1">
    <template v-slot:prepend>
      <v-btn icon="fa-solid fa-book" />
    </template>

    <v-btn v-if="isSubPage" icon="fa-solid fa-home" @click="goToHome" />

    <v-app-bar-title>urB00ks - 网上图书推荐</v-app-bar-title>
    <v-text-field label="搜索相关" density="compact" class="mt-5" variant="outlined"
      prepend-icon="fa-solid fa-magnifying-glass" height="5">
    </v-text-field>
    <v-spacer></v-spacer>
    <v-btn icon="fa-solid fa-circle-user" @click.stop="userBar = !userBar"></v-btn>
  </v-app-bar>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import axios from 'axios';
import router from '@/router';

let isSubPage = ref(false);
let userBar = ref(false);

let isLogIn = ref(false);
let isShowPassword = ref(false);
let loginDialog = ref(false);
let signupDialog = ref(false);
let userName = ref(null);
let userPassword = ref(null);
let userInfo = ref(null);
let isShowLoginDone = ref(false);
let isShowLoginFailed = ref(false);

const goToHome = () => {
  router.push(`/`);
}

const login = async (name, password) => {
  try {
    const response = await axios.get(`api/user?type=name&content=${name}`);
    if (response.status != 404) {
      if (password == response.data.data.password) {
        userInfo = response.data.data;
        loginDialog.value = false;
        isShowLoginDone.value = true;
        isLogIn.value = true;
        userBar.value = false;
      }
      else {
        loginDialog.value = false;
        isShowLoginFailed.value = true;
        userBar.value = false;
      }
    }
    else {
      loginDialog.value = false;
      isShowLoginFailed.value = true;
      userBar.value = false;
    }
  }
  catch (e) {
    console.log("occurred error:" + e);
  }
  finally {
    isUserOnLoading.value = false;
  }
}

watch(
  () => router.currentRoute.value,
  (newRoute) => {
    if (newRoute.path === "/") {
      isSubPage.value = false;
    }
    else {
      isSubPage.value = true;
    }
    console.log("is sub: " + isSubPage.value);
  },
  { immediate: true }
)
</script>