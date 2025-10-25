<template>
    <meta charset="UTF-8" />
    <TopBar />

    <v-container>
        <v-skeleton-loader v-if="isFavBooksOnLoading" type="card-avatar, article, actions" :max-width="300"
            class="mx-auto" />

        <v-row v-else>
            <v-col v-for="book in favBooks" :key="book.id" cols="12" sm="6" md="4" lg="3">
                <v-card class="mx-auto" height="100%" color="light-blue-lighten-4" hover link
                    @click="goToBookDetails(book.id)">
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
    </v-container>
</template>

<script setup>
import { userInfo } from '@/script/user';
import { onMounted, ref } from 'vue';
import router from '@/router';
import axios from 'axios';

let isFavBooksOnLoading = ref(false)
let favBooks = ref(null)

const goToBookDetails = (id) => {
  const response = axios.patch(`/api/book/update?target=${id}&type=update_views&content=add_one`)
  router.push(`/book_info/${id}`)
}

const fetchFavBooks = async () => {
    try {
        const response = await axios.get(`/api/user?type=favs&content=${userInfo.value.id}&password=${userInfo.value.password}`)
        favBooks.value = response.data.data;
        
        isBookListInitialized.value = true;
    }
    catch (e) {
        console.log("occurred error: " + e)
    }
    finally {
        isFavBooksOnLoading.value = false;
    }
}

onMounted(async () => {
    fetchFavBooks();
})

</script>