<template>
  <v-card>
    <v-card-title>Order</v-card-title>
    <v-card-text>
      <v-container>
        <v-alert  
          outlined
          type="warning"
          border="left">
          <h4>Crypto: -> {{ value }}</h4>
        </v-alert>
        <v-alert  
          outlined
          type="warning"
          border="left">
          <h5>Side: -> {{ selected_side }}</h5>
        </v-alert>
        <v-text-field filled label="Amount" v-model="child_amount" required></v-text-field>
      </v-container>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text color="danger" @click="cancel">
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
  name: "diagCard",
  props: {
    crypto: '',
    side: '',
    amount: '',
    value: {
      type: String,
      require: true,
    },
    selected_side: {
      type: String,
      default: '',
      required: true
    }
  },
  data() {
    return {
      child_crypto: this.value,
      child_side: this.selected_side,
      child_amount: 0,
      returnData: {
        crypto: "",
        side: "",
        amount: ""
      }
    }
  },
  methods: {
    submit() {
      this.returnData.crypto = this.child_crypto
      this.returnData.side = this.child_side
      this.returnData.amount = this.child_amount
      this.$emit('clickSubmit', this.returnData)
      this.$router.go({path: this.$router.currentRoute.path, force: true})
    },
    cancel() {
      this.$emit('clickSubmit', "")
    }
  }
}
</script>

