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
    <v-btn icon="fa-solid fa-circle-user" @click.stop="userBar = !userBar"></v-btn>
  </v-app-bar>

  <v-container>
    <v-skeleton-loader v-if="isBookOnLoading" type="card-avatar, article, actions" :max-width="300" class="mx-auto" />

    <v-row v-else>
      <v-col v-for="book in books" :key="book.id" cols="12" sm="6" md="4" lg="3">
        <v-card class="mx-auto" height="100%" hover>
          <v-img :src="book.url" height="200px" cover></v-img>

          <v-card-title>
            {{ book.title }}
          </v-card-title>

          <v-card-subtitle>
            {{ book.author }}
          </v-card-subtitle>

          <v-card-text>
            <div class="text-truncate">{{ book.views }}</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import axios from 'axios'

let userBar = ref(false);
let userName = ref(null);
let isUserOnLoading = ref(true);
let isBookOnLoading = ref(true);
let books = ref([])

const fetchUserName = async () => {
  try {
    const response = await axios.get("api/user/1");
    userName.value = response.data.data.name;
    isUserOnLoading.value = false;
    console.log("user name: " + userName)
  }
  catch (e) {
    console.log("occurred error:" + e);
  }
  finally {
    isUserOnLoading.value = false;
  }
}

const fetchAllBooks = async () => {
  try {
    const response = await axios.get("/api/book/list?num=all")
    books.value = response.data.data;

    for (let index = 0; index < books.length; index++) {
      console.log(books[i])
    }

    isBookOnLoading.value = false;
  }
  catch (e) {
    console.log("occurred error: " + e)
    isBookOnLoading.value = false;
  }
  finally {
    isBookOnLoading.value = false;
  }
}

onMounted(() => { fetchUserName(); })
onMounted(() => { fetchAllBooks(); })

</script>
