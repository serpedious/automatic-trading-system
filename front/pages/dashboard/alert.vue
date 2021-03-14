<template>
  <section class="container">
    <div>
      <!-- <ul v-for="(item, index) of items" v-bind:key="index">
        <li>No.{{ index + 1 }} {{ item.name }}</li>
      </ul> -->
      <ul v-for="balance of balancelists" v-bind:key="balance.amount">
        <li>{{ balance.currency_code }}</li>
        <li>{{ balance.amount }}</li>
      </ul>
      <button @click="apiBalance">show balance</button>
      <button @click="apiBalanceHide">hide balance</button>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Chart',
  data () {
    return {
    //   items: [
    //  { name: 'いちご' },
    //  { name: 'りんご' },
    //  { name: 'みかん' },
    //  { name: 'ぶどう' }
    // ],
      balancelists: []
    }
  },
  methods: {
    apiBalance: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/balance')
      this.balancelists = res.data
    },
    apiBalanceHide: async function () {
      this.balancelists = ''
    }
  }
}
</script>

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
