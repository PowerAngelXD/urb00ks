<template>
  <meta charset="UTF-8">
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

  <v-container>
    <v-skeleton-loader v-if="isBookOnLoading" type="card-avatar, article, actions" :max-width="300" class="mx-auto" />

    <v-row v-else>
      <v-col v-for="book in books" :key="book.id" cols="12" sm="6" md="4" lg="3">
        <v-card class="mx-auto" height="100%" color="light-blue-lighten-4" hover link>
          <v-img :src="book.url" height="200px" cover></v-img>

          <v-card-title>
            {{ book.title }}
          </v-card-title>

          <v-card-subtitle>
            {{ book.author }}
          </v-card-subtitle>

          <v-card-text icon="fa-solid fa-eye">
            <div class="text-truncate">{{ book.views }}</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row v-if="isBookContinueLoading" justify="center" class="my-4">
      <v-progress-circular indeterminate color="blue-lighten-1" size="48" />
    </v-row>

    <v-row v-if="total_book !== null && existed_book_num >= total_book" justify="center" class="my-6">
      <v-alert type="info" variant="tonal" max-width="400">
        已加载所有书籍。
      </v-alert>
    </v-row>

    <v-fab-transition>
      <v-btn v-if="canShowArrowUpButton" color="blue-accent-2" icon="fa-solid fa-arrow-up" size="large" class="ma-4"
        style="position: fixed; bottom: 0; right: 0; z-index: 500;" @click="backUp" />
    </v-fab-transition>
  </v-container>
</template>

<script setup>
import { onMounted, ref, onBeforeUnmount } from 'vue';
import axios from 'axios'

let userBar = ref(false);
let userName = ref(null);
let isUserOnLoading = ref(true);
let isBookOnLoading = ref(true);
let isBookContinueLoading = ref(false)
let books = ref([])
let isBookListInitialized = ref(false)
let existed_book_num = ref(0)
let total_book = ref(0)
let canShowArrowUpButton = ref(false)

const showArrowUpButton = () => {
  if ((window.scrollY || document.documentElement.scrollTop) > 200) {
    canShowArrowUpButton.value = true;
  } else {
    canShowArrowUpButton.value = false;
  }
}

const backUp = () => {
  window.scrollTo({
    top: 0,
    left: 0,
    behavior: 'smooth'
  })
}

const initBookList = async () => {
  try {
    const response = await axios.get("/api/book/count")
    total_book.value = response.data.data
  }
  catch (e) {
    console.log("occurred error: " + e)
  }
  finally { }
}

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

const fetchBooks = async () => {
  try {
    console.log(isBookListInitialized.value)

    if (existed_book_num.value === 0) {
      isBookOnLoading.value = true;
    } else {
      isBookContinueLoading.value = true;
    }

    if (isBookListInitialized.value) {
      // 是否被初始化

      const response = await axios.get("/api/book/list?num=10&offset=" + existed_book_num.value)
      existed_book_num.value += 10;
      let array = ref([]);
      array.value = response.data.data;
      books.value.push(...array.value)

      isBookContinueLoading.value = false
    }
    else {
      const response = await axios.get("/api/book/list?num=10&offset=" + existed_book_num.value)
      existed_book_num.value += 10;
      books.value = response.data.data;
      isBookListInitialized.value = true;
    }

    console.log(existed_book_num.value)
  }
  catch (e) {
    console.log("occurred error: " + e)
  }
  finally {
    isBookOnLoading.value = false;
  }
}

const ScrollToBottom = async () => {
  console.log("e: " + existed_book_num.value + ", total: " + total_book.value)
  if (isBookContinueLoading.value || isBookOnLoading.value) {
    return;
  }

  if (window.scrollY > document.documentElement.offsetHeight - window.innerHeight - 100) {
    fetchBooks().then(() => {
      setTimeout(() => {
        isBookContinueLoading.value = false;
      }, 1000);
    });
  }
}

onMounted(() => {
  initBookList();
  fetchUserName();
  fetchBooks();

  window.addEventListener('scroll', ScrollToBottom);
  window.addEventListener('scroll', showArrowUpButton);
})

onBeforeUnmount(() => {
  window.removeEventListener('scroll', ScrollToBottom);
  window.removeEventListener('scroll', showArrowUpButton);
});
</script>
