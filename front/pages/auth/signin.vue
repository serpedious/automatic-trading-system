<template>
  <div class="signin-container">
    <section class="container">
      <div class="signin">
          <h2>Sign in</h2>
          <input type="text" placeholder="Email" v-model="email">
          <input type="password" placeholder="Password" v-model="password">
          <button @click="apiSignin">Signin</button>
          <p>You don't have an account?
            <router-link to="/auth/signup">create account now!!</router-link>
          </p>
          <router-link to="/">Back to home</router-link>
          <!-- <div class="tips">
            <span style="margin-right:20px;">username: admin</span>
            <span> password: any</span>
          </div> -->
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Signin',
  data: function () {
    return {
      token: '',
      email: '',
      password: ''
    }
  },
  methods: {
      apiSignin: async function() {
        let res = await axios.post(process.env.API_BASE_URL + "/signin", {
            email: this.email,
            password: this.password
        });
        this.token = res.data.token
        this.$router.push("/dashboard");
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
.signin-container {
  min-height: 100%;
  width: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
}
.tips {
  font-size: 14px;
  color: #eee;
  margin-bottom: 10px
} 
</style>