<template>
  <v-alert closable icon="fa-solid fa-check" title="提示" text="登录成功！" color="green" v-model="isShowLoginDone">
  </v-alert>

  <v-alert closable icon="fa-solid fa-triangle-exclamation" title="提示" text="登录失败！" color="red"
    v-model="isShowLoginFailed">
    {{ details }}
  </v-alert>

  <v-alert closable icon="fa-solid fa-check" title="提示" text="注册成功！已为您自动登录" color="green" v-model="isShowRegisterDone">
  </v-alert>

  <v-alert closable icon="fa-solid fa-triangle-exclamation" title="提示" text="注册失败！" color="red"
    v-model="isShowRegisterFailed">
    {{ details }}
  </v-alert>

  <v-dialog v-model="registerDialog" width="600" color="blue lighten-2">
    <v-card color="blue lighten-2">
      <v-toolbar color="blue lighten-2">
        <v-container>
          <v-row align="center">
            <v-btn size="small" icon="fa-solid fa-close" @click="registerDialog = false"></v-btn>
            <v-label>新用户，请注册</v-label>
          </v-row>
        </v-container>
      </v-toolbar>
      <v-container>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-label>请输入填写新账户必要的信息来注册</v-label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-text-field label="账户名称" density="compact" variant="outlined" class="mt-3"
              v-model="userName"></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-text-field label="账户密码" density="compact" variant="outlined" class="mt-3"
              v-model="userPassword"></v-text-field>
          </v-col>
        </v-row>
        <!--
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="7">
            <v-label>生日</v-label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="1"></v-col>
          <v-col cols="4">
            <v-select clearable label="Year" variant="outlined" :items="yearItems" v-model="pickedYear"></v-select>
          </v-col>
          <v-col cols="3">
            <v-select clearable label="Month" variant="outlined" :items="monthItems" v-model="pickedMonth"></v-select>
          </v-col>
          <v-col cols="3">
            <v-select clearable label="Day" variant="outlined" :items="yearItems" v-model="pickedYear"></v-select>
          </v-col>
        </v-row>
          --->
        <v-row>
          <v-col cols="10"></v-col>
          <v-col cols="2">
            <v-btn color="white" variant="outlined" @click="register(userName, userPassword, '2025-10-10')">注册</v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>

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
            <v-text-field label="账户名称" density="compact" variant="outlined" class="mt-3"
              v-model="userName"></v-text-field>
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

  <v-navigation-drawer v-model="userBar" temporary location="right" color="blue-accent-2">
    <br></br>
    <v-list-item v-if="isLogIn" :title="userInfo.name"></v-list-item>
    <v-list density="compact" nav v-else>
      <v-list-item prepend-icon="fa-solid fa-user-plus" title="没有帐户？请注册" value="register"
        @click="registerDialog = true"></v-list-item>
      <v-list-item prepend-icon="fa-solid fa-user" title="已有账户？请登录" value="login"
        @click="loginDialog = true"></v-list-item>
    </v-list>
    <v-divider></v-divider>

    <v-list v-if="isLogIn" density="compact" nav>
      <v-list-item prepend-icon="fa-solid fa-star" title="我的收藏" value="myFavs"
        @click="router.push(`/my_favs/${userInfo.id}`)"></v-list-item>
      <v-list-item prepend-icon="fa-solid fa-gear" title="设置" value="settings"></v-list-item>
    </v-list>
  </v-navigation-drawer>

  <v-app-bar :elevation="2" color="blue-lighten-1">
    <template v-slot:prepend>
      <v-btn icon="fa-solid fa-book" />
    </template>

    <v-btn v-if="isSubPage" icon="fa-solid fa-home" @click="router.push(`/`)" />

    <v-app-bar-title> {{ title }} </v-app-bar-title>
    <v-text-field 
      label="搜索相关" 
      density="compact" 
      class="mt-5" 
      variant="outlined" 
      v-model="text"
      append-inner-icon="fa-solid fa-magnifying-glass" 
      height="5" 
      @click:append-inner="searchBook(text)">
    </v-text-field>
    <v-spacer></v-spacer>
    <v-btn v-if="isShowUserBtn" icon="fa-solid fa-circle-user" @click.stop="userBar = !userBar"></v-btn>
  </v-app-bar>
</template>

<script setup>
import { ref, watch } from 'vue';
import { userInfo, isLogIn, searchResults } from '@/script/user'
import axios from 'axios';
import router from '@/router';

let title = ref("urB00ks - 网上图书推荐");

let isSubPage = ref(false);
let userBar = ref(false);
let isShowUserBtn = ref(false);

let isShowPassword = ref(false);
let loginDialog = ref(false);
let userName = ref(null);
let userPassword = ref(null);
let isShowLoginDone = ref(false);
let isShowLoginFailed = ref(false);

let registerDialog = ref(false);
let isShowRegisterFailed = ref(false);
let isShowRegisterDone = ref(false);

let details = ref(null)
let text = ref("")
const searchBook = async (target) => {
  try {
    const response = await axios.get(`api/book/search/${target}`)

    searchResults.value = response.data.data;

    router.push("/search")
  }
  catch(e) {

  }
  finally{
  }
}

const login = async (name, password) => {
  if (name == null || password == null) {
    details.value = `输入内容不能为空！`
    loginDialog.value = false;
    isShowLoginFailed.value = true;
    userBar.value = false;

    return;
  }

  try {
    const response = await axios.get(`api/user?type=name&content=${name}&password=${password}`);
    userInfo.value = response.data.data;
    loginDialog.value = false;
    isShowLoginDone.value = true;
    isLogIn.value = true;
  }
  catch (e) {
    console.log(`error: ${e.response.data.msg}`)
    details.value = `拒绝登录：${e.response.data.msg}`
    loginDialog.value = false;
    isShowLoginFailed.value = true;
    console.log(e)
  }
  finally {
    userBar.value = false;
  }
}

const register = async (name, password, birthday) => {
  if (name == null || password == null) {
    details.value = `输入内容不能为空！`
    registerDialog.value = false;
    isShowRegisterFailed.value = true;
    userBar.value = false;

    return;
  }

  try {
    const response = await axios.post(`api/user/register?name=${name}&birth=${birthday}&pswd=${password}`);
    if (response.status == 200 && response.data.status == 200) {
      registerDialog.value = false;

      const user_response = await axios.get(`api/user?type=name&content=${name}&password=${password}`);
      userInfo.value = user_response.data.data;
      isLogIn.value = true;
      registerDialog.value = false;
      isShowRegisterDone.value = true;
    }
    else {
      details.value = `注册失败：${response.data.msg}`;
      registerDialog.value = false;
      isShowRegisterFailed.value = true;
    }
  }
  catch (e) {
    console.log("occurred error:" + e);
  }
  finally {
    userBar.value = false;
  }
}

watch(
  () => router.currentRoute.value,
  (newRoute) => {
    if (newRoute.path === "/") {
      isSubPage.value = false;
      isShowUserBtn.value = true;
    }
    else if (newRoute.matched.some(record => record.path === "/book_info/:id")) {
      title.value = "urB00ks - 图书详情"
      isSubPage.value = true;
      isShowUserBtn.value = false;
    }
    else if (newRoute.matched.some(record => record.path === "/my_favs/:id")) {
      title.value = "urB00ks - 我的收藏"
      isSubPage.value = true;
      isShowUserBtn.value = false;
    }
    else if (newRoute.matched.some(record => record.path === "/search")) {
      title.value = "urB00ks - 搜索结果"
      isSubPage.value = true;
      isShowUserBtn.value = false;
    }
    else {
      isSubPage.value = true;
    }
    console.log("is sub: " + isSubPage.value);
  },
  { immediate: true }
)

</script>