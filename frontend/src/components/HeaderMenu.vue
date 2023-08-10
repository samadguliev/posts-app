<template>
  <div>
    <ul>
      <li>
        <router-link to="/post-list">
          List
        </router-link>
      </li>
      <li>
        <router-link to="/create-post">
          Create
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HeaderMenu',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
