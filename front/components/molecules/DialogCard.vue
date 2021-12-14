<template>
  <v-card>
    <v-card-title>Sell Order</v-card-title>
    <v-card-text>
      <v-container>
        <v-row>
          <v-col cols="8">
            <v-text-field label="Amount" v-model="returnData.amount" required></v-text-field>
          </v-col>
        </v-row>
      </v-container>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text color="success" @click="submit">
        Close
      </v-btn>
      <v-btn text color="success" @click="submit">
        Proceed
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    amount: ''
  },
  data() {
    return {
      returnData: {
        amount: this.amount
      }
    }
  },
  methods: {
    submit() {
      this.$emit('clickSubmit', this.returnData)
    },
  }
}
</script>

