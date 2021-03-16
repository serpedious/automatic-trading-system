<template>
  <div class="home-container">
    <section class="container">
      <div>
        <app-logo/>
        <h1>
          Automatic Trading System
        </h1>
        <div class="text-center">
        <v-btn
          :to="'/auth/signin'"
          color="black"
          dark
          class="ma-10"
        >
        GET STARTED
        </v-btn>
      </div>
      </div>
    </section>
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'

export default {
  name: 'HelloWorld',
  components: {
    AppLogo
  },
  data () {
    return {
      msg: 'it is lambda batch test3'
    }
  },
  methods: {

  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.home-container {
  min-height: 100%;
  width: 100%;
  overflow: hidden;
}
</style>

