<template>
    <meta charset="UTF-8" />
    <TopBar />

    <v-container>
        <v-skeleton-loader v-if="isSearchBookOnLoading && !isEmpty" type="card-avatar, article, actions" :max-width="300"
            class="mx-auto" />

        <v-row v-else-if="!isEmpty">
            <v-col v-for="book in results" :key="book.id" cols="12" sm="6" md="4" lg="3">
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
        <v-row v-if="isSearchBookContinueLoading && !isEmpty" justify="center" class="my-4">
            <v-progress-circular indeterminate color="blue-lighten-1" size="48" />
        </v-row>
        <v-row v-if="isEmpty">
            <v-label>没有搜索结果</v-label>
        </v-row>

        <v-fab-transition>
            <v-btn v-if="canShowArrowUpButton" color="blue-accent-2" icon="fa-solid fa-arrow-up" size="large"
                class="ma-4" style="position: fixed; bottom: 0; right: 0; z-index: 500;" @click="backUp" />
        </v-fab-transition>
    </v-container>
</template>

<script setup>
import { searchResults } from '@/script/user';
import { onMounted, ref, onBeforeUnmount } from 'vue';

let isSearchBookOnLoading = ref(false);
let isSearchBookContinueLoading = ref(false);
let isSearchBookInitialized = ref(false)
let canShowArrowUpButton = ref(false);
let isEmpty = ref(false);

let existedResult = ref();
let cursor = ref(0);
let results = ref([]);
let total = ref(0)

const addElements = () => {
    try {
        if (existedResult.value === 0) {
            isSearchBookOnLoading.value = true;
        } else {
            isSearchBookContinueLoading.value = true;
        }

        if (isSearchBookInitialized.value) {
            // 是否被初始化
            let array = searchResults.value.slice(cursor, 10)
            cursor.value += 10;
            results.value.push(...array)

            isSearchBookContinueLoading.value = false
        }
        else {
            let init_array = searchResults.value.slice(cursor, 10);
            cursor.value += 10;
            results.value = init_array;
            isSearchBookInitialized.value = true;
        }
    }
    catch (e) {
        console.log("occurred error: " + e)
    }
    finally {
        isSearchBookOnLoading.value = false;
    }
}

const ScrollToBottom = async () => {
    if (isSearchBookContinueLoading.value || isSearchBookOnLoading.value) {
        return;
    }

    if (window.scrollY > document.documentElement.offsetHeight - window.innerHeight - 100) {
        addElements().then(() => {
            setTimeout(() => {
                isSearchBookContinueLoading.value = false;
            }, 1000);
        });
    }
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

onMounted(() => {
    if (searchResults == []) {
        isEmpty.value = true;
    }
    else {
        total.value = searchResults.value.length
        addElements();

        window.addEventListener('scroll', ScrollToBottom);
        window.addEventListener('scroll', showArrowUpButton);
    }
})

onBeforeUnmount(() => {
    if (searchResults != []) {
        window.removeEventListener('scroll', ScrollToBottom);
        window.removeEventListener('scroll', showArrowUpButton);
    }
});

</script>
