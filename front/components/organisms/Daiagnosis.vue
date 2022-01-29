<template>
  <div>
     <v-container fluid>
    <v-row>
      <v-col
        cols="12"
        sm="12"
        md="6"
        offset-md="3"
      >
        <v-card>
          <v-toolbar extended>
            <v-app-bar-nav-icon></v-app-bar-nav-icon>
            <h1>Diagnosis Robot</h1>
              <v-btn v-show="!hidden" text color="success" @click="sendStart">
                Start
              </v-btn>
          </v-toolbar>
      <tr
        v-for="(msg, index) in messages" :key="index"
      >
        <td>{{ msg }}</td>
      </tr>
    <div>
      <v-text-field v-show="hidden" filled label="Amount" v-model="ddd" required></v-text-field>
      <v-btn v-show="hidden" text color="success" @click="sendMessage">
        Proceed
      </v-btn>
    </div>
     </v-card>
      </v-col>
    </v-row>
  </v-container>
  </div>
</template>

<script>
import io from 'socket.io-client';


export default {
  name: 'room',
  data() {
    return {
      hidden: false,
      messages: [],
      message: "",
      turn: 0,
      ddd: "",
      socket : io(process.env.PY_API_BASE_URL)
    }
  },
  methods: {
    sendMessage() {
      console.log(this.ddd)
      this.socket.send(this.ddd, this.turn);
      console.log(this.turn)
      console.log("turn")
      this.turn += 1
      this.ddd = ''
      if (this.turn === 6) {
        this.turn = 0
        this.hidden = !this.hidden
        this.messages = []
      }
    },
    sendStart() {
      this.hidden = !this.hidden
      this.socket.emit('SEND_MESSAGE', { data: "hoge" });
    }
  },
  mounted() {
    this.socket.on('message', (data) => {
        this.messages = [...this.messages, data];
        console.log('Upload component', data);
        console.log('Upload component', this.message);
        console.log(this.messages.length)
    });
  }
}
</script>

