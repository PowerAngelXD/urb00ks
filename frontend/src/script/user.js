import { ref } from 'vue'
import axios from 'axios'

export let userInfo = ref(null);
export let isLogIn = ref(false);
export let searchResults = ref([])

export const debugInit = async () => {
    const response = await axios.get("api/user?type=name&content=fzsgball&password=123456");
    userInfo.value = response.data.data;
    isLogIn.value = true;
}