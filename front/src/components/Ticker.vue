<template>
  <div class='ticker'>
    <h1>{{ msg }}</h1>
    <p>{{ ticker }}</p>
    <button @click="apiTicker">show ticker</button>
    <button @click="apiTickerHide">hide ticker</button>
    <p>You don't have an account?
      <router-link to="/">go back to home</router-link>
    </p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Ticker',
  data () {
    return {
      msg: 'This is ticker',
      ticker: 'confirm ticker infomation'
    }
  },
  methods: {
    apiTicker: async function () {
      let res = await axios.get('http://localhost:8000/ticker')
      this.ticker = res.data
    },
    apiTickerHide: async function () {
      this.ticker = 'confirm ticker infomation'
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
