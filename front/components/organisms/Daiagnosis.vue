<template>
  <div>
        <v-card class=fluid style="margin: 5px; padding: 30px; width: 100%">
          <v-toolbar extended>
             <v-card-title class="pt-16">
               <h3>Diagnostic Robot</h3>
             </v-card-title>
              <v-btn
                v-show="!hidden"
                @click="sendStart"
                class="mx-2 mt-14"
                fab
                dark
                large
                color="teal"
               >
                <v-icon dark>
                  mdi-android
                </v-icon>
            </v-btn>
          </v-toolbar>
          <v-list three-line>
      <template v-for="(item, index) in messages">
        <v-subheader
          v-if="item.header"
          :key="item.header"
          v-text="item.header"
        ></v-subheader>

        <v-divider
          v-else-if="item.divider"
          :key="index"
          :inset="item.inset"
        ></v-divider>

        <v-list-item
          v-else
          :key="item.title"
        >
          <v-list-item-avatar>
            <v-img :src="item.avatar"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title v-html="item.title"></v-list-item-title>
            <v-list-item-subtitle v-html="item.subtitle"></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
    <div>
      <v-text-field 
        v-show="hidden" 
        filled label="text" 
        v-model="input"
        required
        :counter="10"
        :error-messages="inputErrors"
        @input="$v.input.$touch()"
        @blur="$v.input.$touch()"  
      >
      </v-text-field>
      <v-btn
      v-show="hidden"
      @click="sendMessage"
      outlined
      :disabled="input==''"
      color="teal"
    >
      <v-icon left>
        mdi-send
      </v-icon>
      Send
    </v-btn>
    </div>
     </v-card>
  </div>
</template>

<script>
import io from 'socket.io-client';
import { validationMixin } from 'vuelidate'
import { maxLength } from 'vuelidate/lib/validators'

export default {
  mixins: [validationMixin],
  name: 'room',
  validations: {
      input: { maxLength: maxLength(10) },
    },
  data() {
    return {
      hidden: false,
      messages: [
        { header: 'chat for diagnosis' },
      ],
      message: "",
      turn: 0,
      input: "",
      socket : io(process.env.PY_API_BASE_URL)
    }
  },
   computed: {
      inputErrors () {
        const errors = []
        if (!this.$v.input.$dirty) return errors
        !this.$v.input.maxLength && errors.push('Input must be at most 10 characters long')
        // !this.$v.input.required && errors.push('Input is required.')
        return errors
      },
    },
  methods: {
    sendMessage() {
      this.$v.$touch()
      this.socket.send(this.input, this.turn);
      this.turn += 1
      this.input = ''
      if (this.turn === 6) {
        this.turn = 0
        this.hidden = !this.hidden
        this.messages = [{ header: 'chat for diagnosis' }]
      }
    },
    sendStart() {
      this.hidden = !this.hidden
      this.socket.emit('SEND_MESSAGE', { data: "hoge" });
    }
  },
  mounted() {
    this.socket.on('message', (data) => {
        this.messages = [...this.messages, {divider: true, inset: true}]
        this.messages = [...this.messages, data];
    });
  }
}
</script>

