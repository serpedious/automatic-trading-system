<template>
  <div class="container">
    <div>
    <h1>Automatic Trading System</h1>
    </div>
    <v-btn
      :to="'/auth/signin'"
      color="black"
      dark
      class="ma-10"
    >
    GET STARTED
    </v-btn>
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'
export default {
  components: {
    AppLogo
  },
}
</script>

<style scoped>
.container {
  min-height: 100vh;
  margin-top: 20%;
  justify-content: center;
  align-items: center;
  text-align: center
}
</style>
