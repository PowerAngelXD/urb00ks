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
                    <v-row class="mt-6">
                        <h3> 您可以打分：</h3>
                    </v-row>
                    <v-row class="mt-9">
                        <v-btn-toggle rounded="xl" border>
                            <v-btn>1</v-btn>
                            <v-btn>2</v-btn>
                            <v-btn>3</v-btn>
                            <v-btn>4</v-btn>
                            <v-btn>5</v-btn>
                        </v-btn-toggle>
                    </v-row>
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
import axios from 'axios';

const custom_rating = ref(0);
const route = useRoute();
const book_info = ref(null);

const fetchBookDetails = async () => {
    try {
        const response = await axios.get(`/api/book/${route.params.id}`);
        book_info.value = response.data.data;
    }
    catch (e) {

    }
}

onMounted(() => {
    fetchBookDetails();
})
</script>