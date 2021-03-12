<template>
  <div class="signup-container">
    <section class="container">
      <div class="signup">
        <h2>Sign up</h2>
        <p>{{ email }}</p>
        <input type="text" placeholder="Email" v-model="email" required>
        <input type="password" placeholder="Password" v-model="password" required>
        <button @click="apiSignup">Register</button>
        <p>Do you have an account?
          <router-link to="/auth/signin">sign in now!!</router-link>
        </p>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios'
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
        let res = await axios.post(process.env.API_BASE_URL + "/signup", {
            email: this.email,
            password: this.password
        });
        this.email = res.data.email
        this.$router.push("/auth/signin");
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
  color: #eee;
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
.signup-container {
  min-height: 100%;
  width: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
}
</style>
