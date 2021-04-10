<template>
<section class="container">
    <v-card class="mx-auto mt-5 pa-5" width="800px">
      <v-card-title class="pb-5">
        <h3>Memo</h3>
      </v-card-title>
  <div class="home">
    <v-text-field
     v-model="newTaskTitle"
     @click:append='addTask'
     @keyup.enter="addTask"
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
        v-for="task in tasks"
        :key="task.id"
      >
      <v-list-item
        @click="doneTask(task.id)"
        :class="{ 'blue-grey lighten-4' : task.done }"
      >
        <template v-slot:default>
          <v-list-item-action>
            <v-checkbox
              :input-value="task.done"
              color="primary"
            ></v-checkbox>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title
              :class="{ 'text-decoration-line-through' : task.done }"
            >
              {{ task.title }}
            </v-list-item-title>
          </v-list-item-content>
           <v-list-item-action>
          <v-btn
            @click.stop="deleteTask(task.id)"
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
      newTaskTitle: '',
      tasks: [
        // {
        //   id: 1,
        //   title: 'Wake up',
        //   done: false
        // },
        // {
        //   id: 2,
        //   title: 'Get bananas',
        //   done: false
        // },
        // {
        //   id: 3,
        //   title: 'Eat bananas',
        //   done: false
        // }
      ]
    }
  },
  methods: {
    addTask: async function() {
      let newTask = {
        id: Date.now(),
        title: this.newTaskTitle,
        done: false
      }
      this.tasks.push(newTask)
      let res = await axios.post(process.env.API_BASE_URL + "/creatememo", {
        content: this.newTaskTitle
      });
      this.newTaskTitle = ''
    },
    doneTask: async function(id) {
      let res = await axios.put(process.env.API_BASE_URL + "/donememo", {
        id: 13,
        done: true
      });
      let task = this.tasks.filter(task => task.id === id)[0]
      task.done = !task.done
    },
    deleteTask: async function(id) {
      let res = await axios.put(process.env.API_BASE_URL + "/deletememo", {
        id: 11
      });
      this.tasks = this.tasks.filter(task => task.id !== id)
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
