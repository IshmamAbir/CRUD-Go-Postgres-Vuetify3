<template>
  <v-app class="grey lighten-4">
    <Navbar />
    <!-- <v-content>
      <br><br> -->
      <router-view/>
    <!-- </v-content> -->
  </v-app>
</template>

<script>
import Navbar from './components/Navbar.vue';
import HomeView from './views/HomeView.vue';


export default {
  name: "app",
  components: { Navbar, HomeView }
};
</script>

<style>
@import url(https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css);
</style>