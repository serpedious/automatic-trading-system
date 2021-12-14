<template>
  <v-card>
    <v-card-title>Order</v-card-title>
    <v-card-text>
      <v-container>
        <v-row>
           <v-col
                cols="16"
                sm="12"
              >
                <v-select
                  :items="['BTC_JPY', 'XRP_JPY', 'ETH_JPY', 'XLM_JPY', 'MONA_JPY']"
                  label="Crypto"
                  v-model="returnData.crypto"
                  required
                ></v-select>
              </v-col>
               <v-col
                cols="12"
                sm="6"
              >
                <v-select
                  :items="['SELL', 'BUY']"
                  label="Side"
                  v-model="returnData.side"
                  required
                ></v-select>
              </v-col>
          <v-col cols="6">
            <v-text-field label="Amount" v-model="returnData.amount" required></v-text-field>
          </v-col>
        </v-row>
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
  props: {
    crypto: '',
    side: '',
    amount: ''
  },
  data() {
    return {
      returnData: {
        crypto: this.crypto,
        side: this.side,
        amount: this.amount
      }
    }
  },
  methods: {
    submit() {
      this.$emit('clickSubmit', this.returnData)
    },
    cancel() {
      this.$emit('clickSubmit', "")
    }
  }
}
</script>

