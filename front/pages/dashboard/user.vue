<template>
  <div class="user-container">
    <section class="container">
      <v-card class="mx-auto mt-5 pt-5" width="400px">
        <v-card-title class="pb-10">
          <h2>Edit Profile</h2>
        </v-card-title>
        <v-card-text>
          <form>
            <v-text-field
              v-model="email"
              :error-messages="emailErrors"
              label="E-mail"
              required
              @input="$v.email.$touch()"
              @blur="$v.email.$touch()"
            ></v-text-field>
            <v-text-field
              v-model="password"
              :error-messages="passwordErrors"
              :counter="10"
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
              @click="submit"
              class="ma-4"
            >
              submit
            </v-btn>
            <v-btn 
              @click="clear"
              class="ma-4"
            >
              clear
            </v-btn>
          </form>
        </v-card-text>
      </v-card>
    </section>
  </div>
</template>

<script>
export default {
  name: 'User',
}
</script>

<style scoped>
</style>

<script>
  import { validationMixin } from 'vuelidate'
  import { required, maxLength, email } from 'vuelidate/lib/validators'

  export default {
    mixins: [validationMixin],

    validations: {
      email: { required, email },
      password: { required, maxLength: maxLength(10) },
      checkbox: {
        checked (val) {
          return val
        },
      },
    },

    data: () => ({
      email: '',
      password: '',
      checkbox: false,
    }),

    computed: {
      checkboxErrors () {
        const errors = []
        if (!this.$v.checkbox.$dirty) return errors
        !this.$v.checkbox.checked && errors.push('You must agree to continue!')
        return errors
      },
      passwordErrors () {
        const errors = []
        if (!this.$v.password.$dirty) return errors
        !this.$v.password.maxLength && errors.push('Password must be at most 10 characters long')
        !this.$v.password.required && errors.push('Password is required.')
        return errors
      },
      emailErrors () {
        const errors = []
        if (!this.$v.email.$dirty) return errors
        !this.$v.email.email && errors.push('Must be valid e-mail')
        !this.$v.email.required && errors.push('E-mail is required')
        return errors
      },
    },

    methods: {
      submit () {
        this.$v.$touch()
      },
      clear () {
        this.$v.$reset()
        this.email = ''
        this.password = ''
        this.checkbox = false
      },
    },
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
.user-container {
  min-height: 100%;
  width: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  background-color: #2d3a4b;
}
</style>