<template>
  <div class="ml-5 mt-10">
    <v-btn color="error" @click.stop="dialog = true" class="ma-2">
      Sell
    </v-btn>
    <v-btn color="success" @click.stop="dialog = true" class="ma-2">
      Buy
    </v-btn>
    <v-dialog v-model="dialog" max-width="400px">
      <dialog-card
        v-on:clickSubmit="onSubmit"
        :name="amount"
      ></dialog-card>
    </v-dialog>
  </div>
</template>

<script>
import DialogCard from '../molecules/DialogCard.vue'
import axios from 'axios'

export default {
  components: {
    DialogCard
  },
  data() {
    return {
      dialog: false,
      crypto: "",
      amount: "",
      side: ""
    }
  },
  methods: {
    onSubmit(params) {
      if (params === "") {
          this.dialog = false
          return
      }
      this.dialog = false
      this.crypto = params.crypto
      this.side = params.side
      this.amount = parseFloat(params.amount)
      this.sendOrder();
    },
    sendOrder: async function() {
      await axios.post(process.env.API_BASE_URL + "/sendorder", {
        product_code: this.crypto,
        child_order_type: "MARKET",
        size: this.amount,
        side: this.side,
        minute_to_expire: 1,
        time_in_force: "GTC"
      });
    },
  }
}
</script>

<style>

</style>
