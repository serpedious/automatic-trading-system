<template>
  <section>
  <v-card class="mx-auto pa-5 ma-3 ml-4" width="570px">
    <v-card-title class="pb-5">
      <h3>CSV</h3>
    </v-card-title>
    <v-tabs
      centered
      fixed-tabs
      v-model="tab"
      background-color="primary"
      dark
      slider-color="purple"
    >
      <v-tab
        v-for="item in items"
        :key="item.tab"
      >
        {{ item.tab }}
      </v-tab>
    </v-tabs>
    

    <v-tabs-items v-model="tab">
      <v-tab-item
        v-for="item in items"
        :key="item.tab"
      >
        <v-card flat>
          <v-btn
          color="purple"
          dark
          class="ma-5"
          @click="downloadCsv"
        >
        DownLoad
        </v-btn>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
  </section>
</template>

<script>
import axios from 'axios'
  export default {
    name: 'Csv',
    data () {
      return {
        tab: null,
        items: [
          { tab: 'history', content: 'Tab 1 Content' },
          { tab: 'profit', content: 'Tab 2 Content' },
        ],
      }
    },
    methods: {
    downloadCsv: async function () {
      // let res = await axios.get(process.env.API_BASE_URL + '/getcsvfile')
    axios({
      url: 'http://localhost:8000/getcsvfile',
      method: 'GET',
      responseType: 'blob',
    }).then((response) => {
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', 'history.csv');
      document.body.appendChild(link);
      link.click();
    });
      
    },
    }
  }
</script>

<style scoped>
.container {
  min-height: 70vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>