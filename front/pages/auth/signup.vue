<template>
  <section class="container">
    <v-card class="mx-auto mt-5 pa-5" width="400px" v-if="loading === 0">
      <v-card-title class="pb-10">
        <h2>Signup</h2>
      </v-card-title>
      <v-card-text>
        <form>
          <v-text-field
            v-model="email"
            :counter="20"
            label="E-mail"
            required
            :error-messages="emailErrors"
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
            @click="apiSignup"
            :disabled="password=='' || email=='' || checkbox === false"
          >
            create
          </v-btn>
          <p>You already have an account?
            <router-link to="/auth/signin">signin here</router-link>
          </p>
        </form>
      </v-card-text>
    </v-card>
    <v-card class="mx-auto mt-5 pa-16" width="400px" height="464" v-else>
    <v-card-actions class="justify-center pt-16 ma-16">
       <v-progress-circular
          class="pt-16"
          indeterminate
          color="primary"
          ></v-progress-circular>
 </v-card-actions>
  </v-card>
  </section>
</template>

<script>
import axios from 'axios'
import { validationMixin } from 'vuelidate'
import { required, maxLength } from 'vuelidate/lib/validators'
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
  name: 'Signup',
  data () {
    return {
      email: '',
      password: '',
      checkbox: false,
      loading: 0,
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
      apiSignup: async function() {
        this.loading = 1
        this.$v.$touch()
        let res = await axios.post(process.env.API_BASE_URL + "/signup", {
            email: this.email,
            password: this.password
        }).catch(error => {
          console.log(error)
          this.loading = 0
        });
        this.email = res.data.email
        this.$router.push("/dashboard")
        this.loading = 0
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
