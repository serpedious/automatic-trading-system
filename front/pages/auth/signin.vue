<template>
  <section class="container">
    <div class="signin">
        <h2>Sign in</h2>
        <input type="text" placeholder="email" v-model="email">
        <input type="password" placeholder="Password" v-model="password">
        <button @click="apiSignin">Signin</button>
        <p>You don't have an account?
            <router-link to="/auth/signup">create account now!!</router-link>
        </p>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Signin',
  data: function () {
    return {
      token: ''
    }
  },
  methods: {
      apiSignin: async function() {
        let res = await axios.post(process.env.API_BASE_URL + "/signin", {
            email: this.email,
            password: this.password
        });
        this.token = res.data.token
        this.$router.push("/users");
      }
  }
}
</script>

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
.signin {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>