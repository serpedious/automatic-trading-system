<template>
<section class="container">
    <v-card class="mx-auto ma-3 pa-5" width="1150px">
      <v-card-title class="pb-5">
        <h3>Memo</h3>
      </v-card-title>
  <div class="home">
    <v-text-field
     v-model="newMemoContent"
     @click:append='addMemo'
     @keyup.enter="addMemo"
      class="pa-3"
      outlined
      label="Add Memo"
      append-icon="mdi-plus-circle"
      hide-details
      clearable
    ></v-text-field>

    <v-list
      class="pt-0"
      flat
    >
      <div
        v-for="memo in allMemos"
        :key="memo.id"
      >
      <v-list-item
        @click="doneMemo(memo.id)"
        :class="{ 'blue-grey lighten-4' : memo.done }"
      >
        <template v-slot:default>
          <v-list-item-action>
            <v-checkbox
              :input-value="memo.done"
              color="red"
            ></v-checkbox>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title
              :class="{ 'text-decoration-line-through' : memo.done }"
            >
              {{ memo.content }}
            </v-list-item-title>
          </v-list-item-content>
           <v-list-item-action>
          <v-btn
            @click.stop="deleteMemo(memo.id)"
            icon
          >
            <v-icon color="grey">mdi-delete-circle</v-icon>
          </v-btn>
          </v-list-item-action>

        </template>

       </v-list-item>
       <v-divider></v-divider>
      </div>
    </v-list>
  </div>
    </v-card>
</section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Memo',
  data() {
    return {
      newMemoContent: '',
      allMemos: []
    }
  },
  mounted() {
    this.getAllMemos();
  },
  methods: {
    getAllMemos: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/getallmemos')
      this.allMemos = res.data
    },
    addMemo: async function() {
      let newMemo = {
        id: Date.now(),
        content: this.newMemoContent,
        done: false,
        delete: false
      }
      if (this.newMemoContent.length <= 0) {
        console.log('please fill in something');
      } else {
        this.allMemos.push(newMemo)
        let res = await axios.post(process.env.API_BASE_URL + "/creatememo", {
          content: this.newMemoContent
        });
        this.newMemoContent = ''
        this.getAllMemos();
     }
    },
    doneMemo: async function(id) {
      let res = await axios.put(process.env.API_BASE_URL + "/donememo", {
        id: id
      });
      let memo = this.allMemos.filter(memo => memo.id === id)[0]
      memo.done = !memo.done
    },
    deleteMemo: async function(id) {
      let res = await axios.put(process.env.API_BASE_URL + "/deletememo", {
        id: id
      });
      this.allMemos = this.allMemos.filter(memo => memo.id !== id)
    }
  }
}
</script>

<style>
.container {
  min-height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
