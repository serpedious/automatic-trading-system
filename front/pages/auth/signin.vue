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
            :error-messages="emailErrors"
            required
            :counter="20"
            @input="$v.email.$touch()"
            @blur="$v.email.$touch()"
          ></v-text-field>
          <v-text-field
            v-model="password"
            :counter="10"
            :error-messages="passwordErrors"
            label="password"
            required
            @input="$v.password.$touch()"
            @blur="$v.password.$touch()"  
          ></v-text-field>
          <v-checkbox
            v-model="checkbox"
            :error-messages="checkboxErrors"
            label="Do you agree?"
            required
            @change="$v.checkbox.$touch()"
            @blur="$v.checkbox.$touch()"
          ></v-checkbox>
          <v-btn
            class="ma-5"
            @click="apiSignin"
            :disabled="password=='' || email=='' || checkbox === false"
          >
            submit
          </v-btn>
          <p>You don't have an account?
            <router-link to="/auth/signup">create account now</router-link>
          </p>
          <v-btn
            class="ma-1"
            @click="guestSignin"
            color="primary"
          >
            guest
          </v-btn>
        </form>
      </v-card-text>
    </v-card>
  </section>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, maxLength } from 'vuelidate/lib/validators'
import axios from 'axios'
axios.defaults.withCredentials = true;
export default {
   mixins: [validationMixin],
    validations: {
      email: { required, maxLength: maxLength(20) },
      password: { required, maxLength: maxLength(10) },
      checkbox: {
        checked (val) {
          return val
        },
      },
    },
  name: 'Signin',
  data: function () {
    return {
      token: '',
      email: '',
      password: '',
      checkbox: false,
    }
  },
   computed: {
      checkboxErrors () {
        const errors = []
        if (!this.$v.checkbox.$dirty) return errors
        !this.$v.checkbox.checked && errors.push('You must agree to continue!')
        return errors
      },
      emailErrors () {
        const errors = []
        if (!this.$v.password.$dirty) return errors
        !this.$v.email.maxLength && errors.push('Email must be at most 20 characters long')
        !this.$v.email.required && errors.push('Email is required.')
        return errors
      },
      passwordErrors () {
        const errors = []
        if (!this.$v.password.$dirty) return errors
        !this.$v.password.maxLength && errors.push('Password must be at most 10 characters long')
        !this.$v.password.required && errors.push('Password is required.')
        return errors
      },
    },
  methods: {
      apiSignin: async function() {
        this.$v.$touch()
        let res = await axios.post(process.env.API_BASE_URL + "/signin", {
            email: this.email,
            password: this.password,
        }, { withCredentials: true });
        this.token = res.data.token
        this.$router.push("/dashboard");
      },
      guestSignin: async function() {
        let res = await axios.post(process.env.API_BASE_URL + "/signin", {
            email: "test@test.com",
            password: "hogehoge",
        }, { withCredentials: true });
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
</style>
