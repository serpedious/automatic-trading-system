<template>
  <div class="home-container">
    <section class="container">
      <div>
        <app-logo/>
        <h1 class="title">
          Automatic Trading System
        </h1>
        <div class="links">
          <p
            class="button--green"
            target="_blank">
            <router-link to="/auth/signup">Signup</router-link>
          </p>
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
  background-color: #2d3a4b;
  overflow: hidden;
}
</style>

