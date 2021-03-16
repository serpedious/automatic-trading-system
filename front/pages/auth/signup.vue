<template>
   <div class="signup-container">
    <section class="container">
      <v-card class="mx-auto mt-5 pa-5" width="400px">
        <v-card-title class="pb-10">
          <h2>Signup</h2>
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
              class="ma-3"
              @click="apiSignup"
            >
              create
            </v-btn>
            <v-btn
              class="ma-3"
              :to="'/dashboard'"
            >
              dashboard
            </v-btn>
            <p>You already have an account?
              <router-link to="/auth/signin">signin here</router-link>
            </p>
          </form>
        </v-card-text>
      </v-card>
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
.signup-container {
  min-height: 100%;
  width: 100%;
  overflow: hidden;
}
p {
  padding-top: 30px;
}
</style>
