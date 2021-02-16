<template>
  <div class="signup">
    <h2>Sign up</h2>
    <p>{{ email }}</p>
    <input type="text" placeholder="Email" v-model="email" required>
    <input type="password" placeholder="Password" v-model="password" required>
    <button @click="apiSignup">Register</button>
    <p>Do you have an account?
      <router-link to="/signin">sign in now!!</router-link>
    </p>
    <p>Do you want to confirm your balance?
      <router-link to="/balance">Let's try to see your balance</router-link>
    </p>
    <p>Do you want to confirm ticker?
      <router-link to="/ticker">Let's try to see ticker infomation</router-link>
    </p>
  </div>
</template>

<script>
import axios from 'axios'
import router from '../router'
export default {
  name: 'Signup',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
      apiSignup: async function() {
        let res = await axios.post("https://api.serpedious.link/signup", {
            email: this.email,
            password: this.password
        });
        this.email = res.data.email
        router.push("/memo");
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
.signup {
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
