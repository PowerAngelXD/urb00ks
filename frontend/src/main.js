/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Styles
import 'unfonts.css'
import '@fortawesome/fontawesome-free/css/all.css'
import axios from 'axios'

// 请求拦截器
axios.interceptors.request.use(
    config => {
        const token = localStorage.getItem('auth_token');
        if (token) {
            config.headers['Authorization'] = 'Bearer ' + token;
        }
        return config;
    }, 
    error => {
        return Promise.reject(error);
    }
);

const app = createApp(App)

registerPlugins(app)

app.mount('#app')
