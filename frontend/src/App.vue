<template>
  <HeaderMenu />
  <router-view />
</template>

<script>
import HeaderMenu from "@/components/HeaderMenu.vue";

export default {
  name: 'App',
  components: {
    HeaderMenu
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
