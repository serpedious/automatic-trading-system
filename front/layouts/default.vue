<template>
  <v-app id="inspire">
    <Header/>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import Header from '~/components/organisms/Header.vue'
export default {
  components: {
  Header
  }
}
</script>

<style>
</style>

