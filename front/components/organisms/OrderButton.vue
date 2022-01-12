<template>
  <div class="ml-5 mt-10">
    <v-btn color="error" v-on:click="setSellSide" @click.stop="dialog = true" class="ma-2">
      Sell
    </v-btn>
    <v-btn color="success" v-on:click="setBuySide" @click.stop="dialog = true" class="ma-2">
      Buy
    </v-btn>
    <v-dialog v-model="dialog" max-width="400px">
      <dialog-card
        v-on:clickSubmit="onSubmit"
        :name="amount"
        :value="message"
        :selected_side="selected_side"
      ></dialog-card>
    </v-dialog>
  </div>
</template>

<script>
import DialogCard from '../molecules/DialogCard.vue'
import axios from 'axios'

export default {
  props: {
    message: {
      type: String,
      require: true,
    }
  },
  components: {
    DialogCard
  },
  data() {
    return {
      dialog: false,
      crypto: "",
      amount: "",
      side: "",
      selected_side: ""
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
    setBuySide: function () {
      this.selected_side = 'BUY'
    },
    setSellSide: function () {
      this.selected_side = 'SELL'
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
