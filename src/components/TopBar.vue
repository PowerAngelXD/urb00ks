<template>
    <v-navigation-drawer v-model="userBar" temporary location="right" color="blue-accent-2">
    <br></br>
    <v-list-item :title="userName"></v-list-item>

    <v-divider></v-divider>

    <v-list density="compact" nav>
      <v-list-item prepend-icon="fa-solid fa-star" title="我的收藏" value="myFavs"></v-list-item>
      <v-list-item prepend-icon="fa-solid fa-gear" title="设置" value="settings"></v-list-item>
    </v-list>
  </v-navigation-drawer>

  <v-app-bar :elevation="2" color="blue-lighten-1">
    <template v-slot:prepend>
      <v-btn icon="fa-solid fa-book" />
    </template>

    <v-app-bar-title>urB00ks - 网上图书推荐</v-app-bar-title>
    <v-text-field label="搜索相关" density="compact" class="mt-5" variant="outlined"
      prepend-icon="fa-solid fa-magnifying-glass" height="5">
    </v-text-field>
    <v-spacer></v-spacer>
    <v-btn icon="fa-solid fa-circle-user" @click.stop="userBar = !userBar"></v-btn>
  </v-app-bar>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

let userName = ref(null);
let userBar = ref(false);

const fetchUserName = async () => {
  try {
    const response = await axios.get("api/user/1");
    userName.value = response.data.data.name;
    isUserOnLoading.value = false;
    console.log("user name: " + userName.value)
  }
  catch (e) {
    console.log("occurred error:" + e);
  }
  finally {
    isUserOnLoading.value = false;
  }
}

onMounted(() => { fetchUserName(); })
</script>