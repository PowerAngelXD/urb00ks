<template>
  <meta charset="UTF-8">
  <TopBar/>

  <v-container>
    <v-skeleton-loader v-if="isBookOnLoading" color="white" type="card-avatar, article, actions" :max-width="300" class="mx-auto" />

    <v-row v-else>
      <v-col v-for="book in books" :key="book.id" cols="12" sm="6" md="4" lg="3">
        <v-card class="mx-auto" height="100%" color="light-blue-lighten-4" hover link @click="goToBookDetails(book.id)">
          <v-img :src="book.url" height="200px" cover></v-img>

          <v-card-title>
            {{ book.title }}
          </v-card-title>

          <v-card-subtitle>
            {{ book.author }}
          </v-card-subtitle>

          <v-card-text>
            👁 {{ book.views }}
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row v-if="isBookContinueLoading" justify="center" class="my-4">
      <v-progress-circular indeterminate color="blue-lighten-1" size="48" />
    </v-row>

    <v-row v-if="(total_book !== null && existed_book_num >= total_book) && !isBookOnLoading" justify="center" class="my-6">
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
import TopBar from '@/components/TopBar.vue';
import { onMounted, ref, onBeforeUnmount } from 'vue';
import axios from 'axios';
import router from '@/router';

let isUserOnLoading = ref(true);
let isBookOnLoading = ref(true);
let isBookContinueLoading = ref(false);
let books = ref([]);
let isBookListInitialized = ref(false);
let existed_book_num = ref(0);
let total_book = ref(0);
let canShowArrowUpButton = ref(false);

const goToBookDetails = (id) => {
  const response = axios.patch(`/api/book/update?target=${id}&type=update_views&content=add_one`)
  router.push(`/book_info/${id}`)
}

const showArrowUpButton = () => {
  if ((window.scrollY || document.documentElement.scrollTop) > 500) {
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
  finally {}
}

const fetchBooks = async () => {
  try {
    console.log(isBookListInitialized.value + ":" + existed_book_num.value)

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
      const response = await axios.get("/api/book/list?num=10&offset=0")
      existed_book_num.value += 10;
      books.value = response.data.data;
      console.log(books.value)
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
  fetchBooks();

  window.addEventListener('scroll', ScrollToBottom);
  window.addEventListener('scroll', showArrowUpButton);
})

onBeforeUnmount(() => {
  window.removeEventListener('scroll', ScrollToBottom);
  window.removeEventListener('scroll', showArrowUpButton);
});
</script>
