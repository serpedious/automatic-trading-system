<template>
  <section class="container">
    <v-card class="mx-auto mt-5 pt-5" width="400px">
      <v-card-title class="pb-10">
        <h2>Edit Password</h2>
      </v-card-title>
      <v-card-text>
        <form>
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
            :disabled="password=='' || checkbox === false"
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
</template>

<script>
  import { validationMixin } from 'vuelidate'
  import { required, maxLength } from 'vuelidate/lib/validators'
  import axios from 'axios'

  export default {
    mixins: [validationMixin],

    validations: {
      password: { required, maxLength: maxLength(10) },
      checkbox: {
        checked (val) {
          return val
        },
      },
    },

    data: () => ({
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
    },

    methods: {
      submit () {
        this.$v.$touch()
        this.apiEdit();
      },
      clear () {
        this.$v.$reset()
        this.password = ''
        this.checkbox = false
      },
      apiEdit: async function() {
        console.log(this.password)
        let res = await axios.post(process.env.API_BASE_URL + "/editpass", {
          password: this.password
        });
        if (res.status === 200) {
          location.href='/dashboard/home';
        }
      },
    },
  }
</script>

<style scoped>
.container {
  min-height: 90vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>