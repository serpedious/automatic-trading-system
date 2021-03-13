<template>
  <v-app id="inspire" style="background-color: #ECEFF1;">
    <div v-if="this.$route.path.match(/dashboard\/*/)">
      <Header/>
    </div>

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

