<template>
  <section class="container">
    <div>
      <app-logo/>
      <h1 class="title">
        front
      </h1>
      <h2 class="subtitle">
        automatic-trading-front-test-page
      </h2>
      <h1>{{ msg }}</h1>
      <h2>Essential Links</h2>
      <button @click="apiPublic">public</button>
      <button @click="apiPrivate">private</button>
      <div class="links">
        <a
          href="https://nuxtjs.org/"
          target="_blank"
          class="button--green">Documentation</a>
        <a
          href="https://github.com/nuxt/nuxt.js"
          target="_blank"
          class="button--grey">GitHub</a>
        <p>You don't have an account?
          <router-link to="/auth/signup">create account now!!</router-link>
        </p>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
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
    apiPublic: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/public')
      this.msg = res.data
    },
    apiPrivate: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/private')
      this.msg = res.data
    }
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
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>

