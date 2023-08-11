<template>
  <div class="posts-container">
    <h1 class="posts-header">Post list</h1>

    <div class="post">
      <p class="post-title">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit
      </p>
      <p class="post-description">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
        dolore magna aliqua. Duis at tellus at urna condimentum. Ipsum dolor sit amet consectetur adipiscing elit.
        Vitae nunc sed velit dignissim.
      </p>
      <p class="post-date">
        22.07.2023
      </p>
    </div>

    <div class="post">
      <p class="post-title">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit
      </p>
      <p class="post-description">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
        dolore magna aliqua. Duis at tellus at urna condimentum. Ipsum dolor sit amet consectetur adipiscing elit.
        Vitae nunc sed velit dignissim.
      </p>
      <p class="post-date">
        22.07.2023
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PostList',
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
