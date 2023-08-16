<template>
  <div class="create-post-container">
    <h1 class="create-post-header">
      Create post
    </h1>

    <div class="create-post-form-container">
      <input
          class="create-post-input"
          type="text"
          v-model="title"
          placeholder="Title"
      >

      <input
          class="create-post-input"
          type="text"
          v-model="description"
          placeholder="Description"
      >

      <textarea
          class="create-post-text"
          name="text"
          cols="50"
          rows="10"
          v-model="text"
      />
    </div>

    <button type="submit" class="create-post-button" @click="save()">
      Save
    </button>
  </div>
</template>

<script>
import axios from "axios";
const API_URL = "http://localhost:3000"

export default {
  name: 'CreatePost',
  data() {
    return {
      title: "",
      description: "",
      text: "",
    }
  },
  methods: {
    save() {
      axios
          .post(`${API_URL}/posts`, {
            title: this.title,
            description: this.description,
            text: this.text,
          })
          .catch(error => console.log(error))

      this.title = ""
      this.description = ""
      this.text = ""
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.create-post-header {
  text-align: center;
}

.create-post-form-container {
  max-width: 60%;
  margin: 0 auto;
}

.create-post-input {
  display: block;
  min-width: 300px;
  margin: 0 auto 30px;
  padding: 10px 20px;
  border-radius: 5px;
  border-width: 1px;
  font-size: 16px;
}

.create-post-text {
  margin: 0 auto;
  display: block;
}

.create-post-button {
  display: block;
  margin: 50px auto 0;
  font-size: 16px;
  padding: 10px;
  width: 150px;
  background-color: #00ce0d;
  color: #fff;
  border: 0;

}
</style>
