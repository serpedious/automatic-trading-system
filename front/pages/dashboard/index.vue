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
          <a
            href="https://github.com/serpedious/automatic-trading-system"
            target="_blank"
            class="button--green">GitHub</a>
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

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 300;
  font-size: 60px;
  color: #eee;
  letter-spacing: 1px;
}

.links {
  padding-top: 15px;
  color: #42b983;
}

.home-container {
  min-height: 100%;
  width: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
}
</style>

