<template>
    <meta charset="UTF-8" />
    <TopBar />

    <v-container v-if="book_info">
        <v-row cols="4">
            <v-col>
                <img :width="400" :src="book_info.url"></img>
            </v-col>
            <v-col>
                <v-container>
                    <v-row>
                        <h1> {{ book_info.title }} </h1>
                    </v-row>
                    <v-row>
                        <h2> {{ book_info.author }} </h2>
                    </v-row>
                    <v-row>
                        <h3> views: {{ book_info.views }} </h3>
                    </v-row>
                    <v-row>
                        <h3> rating: {{ book_info.rating }}</h3>
                    </v-row>
                    <v-row class="mt-7">
                        <v-btn v-if="isFaved && isInFavs() && isLogIn" prepend-icon="fa-solid fa-star" @click="isFaved = true">收藏</v-btn>
                        <v-alert v-else-if="isFaved && isInFavs()" type="success" variant="tonal" max-width="400"> 已收藏 </v-alert>
                    </v-row>
                    <div v-if="!isJudged">
                        <v-row class="mt-6">
                            <h3> 您可以打分：</h3>
                        </v-row>
                        <v-row class="mt-9">
                            <v-btn-toggle v-model="custom_rating" rounded="xl" border>
                                <v-btn v-for="score in 5" @click="judgeBook(score)"> {{ score }} </v-btn>
                            </v-btn-toggle>
                        </v-row>
                    </div>
                    <div v-else>
                        <v-row class="mt-6">
                            <v-alert type="success" variant="tonal" max-width="400">
                                您已成功评分
                            </v-alert>
                        </v-row>
                    </div>
                </v-container>
            </v-col>
        </v-row>
    </v-container>
    <v-container v-else>
        <v-col></v-col>
        <v-col>
            <v-progress-circular indeterminate color="blue-lighten-1" size="48" />
        </v-col>
        <v-col>
            <h3>Loading .....</h3>
        </v-col>
        <v-col></v-col>
    </v-container>
</template>

<script setup>
import TopBar from '@/components/TopBar.vue';
import { onMounted, ref, onBeforeUnmount } from 'vue';
import { useRoute } from 'vue-router';
import { isLogIn, userInfo } from '@/script/user'
import axios from 'axios';

let isJudged = ref(false);
let isFaved = ref(false);
const custom_rating = ref(null);
const route = useRoute();
const book_info = ref(null);

const isInFavs = () => {
    return userInfo.value.favs.indexOf() != -1;
}

const fetchBookDetails = async () => {
    try {
        console.log("fetch bool: " + isJudged.value)
        const response = await axios.get(`/api/book/${route.params.id}`);
        book_info.value = response.data.data;
        console.log(response)
    }
    catch (e) {

    }
}

const judgeBook = async (score) => {
    try {
        const response = await axios.patch(`/api/book/update?target=${route.params.id}&type=update_rating&content=${score}`)
        console.log(response)
        isJudged.value = true;
    }
    catch (e) {

    }
}

onMounted(() => {
    console.log(isJudged.value)
    fetchBookDetails();
})
</script>