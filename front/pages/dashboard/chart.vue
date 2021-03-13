<template>
  <section class="container">
    <div>
      <h1>{{ msg }}</h1>
      <p>{{ deposit }}</p>
      <button @click="apiBalance">show balance</button>
      <button @click="apiBalanceHide">hide balance</button>
      <p>You don't have an account?
        <router-link to="/">go back to home</router-link>
      </p>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Chart',
  data () {
    return {
      msg: 'This is your balance',
      deposit: 'confirm your balance'
    }
  },
  methods: {
    apiBalance: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/balance')
      this.deposit = res.data
    },
    apiBalanceHide: async function () {
      this.deposit = 'confirm your balance'
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
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
