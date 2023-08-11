<template>
  <div class="header-menu-container">
    <ul class="header-menu-list">
      <li class="header-menu-element">
        <router-link to="/post-list">
          List
        </router-link>
      </li>
      <li class="header-menu-element">
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
.header-menu-container {
  background-color: grey;
  width: 100%;
}

.header-menu-list {
  list-style: none;
  display: flex;
  justify-content: flex-start;
  padding: 20px 40px;
  font-size: 20px;
  margin-top: 0;
}

.header-menu-element {
  margin-right: 20px;
}

.header-menu-element a {
  color: #fff;
  text-decoration: none;
}

</style>
