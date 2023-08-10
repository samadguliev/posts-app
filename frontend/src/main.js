import {createApp} from 'vue'
import {createRouter, createWebHistory} from 'vue-router'
import App from './App.vue'

import PostList from "@/components/PostList.vue";
import CreatePost from "@/components/CreatePost.vue";

const router = createRouter({
    routes: [
        {
            path: '/post-list',
            component: PostList
        },
        {
            path: '/create-post',
            component: CreatePost
        }
    ],
    history: createWebHistory()
})

const app = createApp(App)
app.use(router)
app.mount('#app')
