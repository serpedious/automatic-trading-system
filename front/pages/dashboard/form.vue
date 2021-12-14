<template>
  <div>
    <v-btn color="success" @click.stop="dialog = true">
      OPEN
    </v-btn>
    <v-dialog v-model="dialog" max-width="400px">
      <dialog-card
        v-on:clickSubmit="onSubmit"
        :name="amount"
      ></dialog-card>
    </v-dialog>
    <v-text-field v-model="amount" label="Amount" disabled></v-text-field>
  </div>
</template>

<script>
import DialogCard from '../../components/molecules/DialogCard.vue'
import axios from 'axios'

export default {
  components: {
    DialogCard
  },
  data() {
    return {
      dialog: false,
      amount: ""
    }
  },
  methods: {
    onSubmit(params) {
      this.dialog =false
      this.amount = parseFloat(params.amount)
      console.log(this.amount)
      this.sendOrder();
    },
    sendOrder: async function() {
      await axios.post(process.env.API_BASE_URL + "/sendorder", {
        product_code: "XRP_JPY",
        child_order_type: "MARKET",
        size: this.amount,
        side: "BUY",
        minute_to_expire: 1,
        time_in_force: "GTC"
      });
    },
  }
}
</script>

<style>

</style>


	ProductCode:     "XRP_JPY",
	// 	ChildOrderType:  "MARKET",
	// 	Side:            "BUY",
	// 	Size:            size,
	// 	MinuteToExpires: 1,
	// 	TimeInForce:     "GTC",

    	ID                     int     `json:"id"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	ProductCode            string  `json:"product_code"`
	ChildOrderType         string  `json:"child_order_type"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	MinuteToExpires        int     `json:"minute_to_expire"`
	TimeInForce            string  `json:"time_in_force"`
	Status                 string  `json:"status"`
	ErrorMessage           string  `json:"error_message"`
	AveragePrice           float64 `json:"average_price"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
	Count                  int     `json:"count"`
	Before                 int     `json:"before"`
	After                  int     `json:"after"`