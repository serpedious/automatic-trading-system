<template>
  <section class="container">
    <v-card class="mx-auto mt-5 pa-5" width="400px">
      <v-card-title class="pb-10">
        <h2>Signin</h2>
      </v-card-title>
      <v-card-text>
        <form>
          <v-text-field
            v-model="email"
            label="E-mail"
            required
            @input="$v.email.$touch()"
            @blur="$v.email.$touch()"
          ></v-text-field>
          <v-text-field
            v-model="password"
            :counter="10"
            label="password"
            required
            @input="$v.password.$touch()"
            @blur="$v.password.$touch()"
          ></v-text-field>
          <v-btn
            class="mt-5"
            @click="apiSignin"
          >
            submit
          </v-btn>
          <p>You don't have an account?
            <router-link to="/auth/signup">create account now</router-link>
          </p>
        </form>
      </v-card-text>
    </v-card>
  </section>
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
p {
  padding-top: 30px;
}
</style>