<template>
  <div class="posts-container">
    <h1 class="posts-header">Post list</h1>

    <div class="post" v-for="(post, index) in posts" :key="index">
      <p class="post-title">
        {{ post.title }}
      </p>
      <p class="post-description">
        {{ post.description }}
      </p>
      <p class="post-date">
        {{ post.created_at }}
      </p>
    </div>

  </div>
</template>

<script>
import axios from 'axios'
const API_URL = "http://localhost:3000"

export default {
  name: 'PostList',
  data() {
    return {
      posts: []
    }
  },
  mounted() {
    this.getPostList()
  },
  methods: {
    getPostList() {
      axios
          .get(`${API_URL}/posts`)
          .then(response => this.posts = response.data)
          .catch(error => console.log(error))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.posts-container {
  margin-top: 50px;
}

.posts-header {
  text-align: center;
}

.post {
  max-width: 65%;
  background-color: grey;
  color: white;
  padding: 30px;
  border-radius: 10px;
  margin: 0 auto 30px;
}

.post-title {
  font-size: 18px;
}

</style>
