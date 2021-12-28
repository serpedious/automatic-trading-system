<template>
  <section>
    <v-card class=fluid style="margin: 5px; padding: 30px; width: 100%">
      <v-card-title class="pb-3">
        <h3>Alert</h3>
      </v-card-title>
    <v-tabs
      centered
      fixed-tabs
      v-model="tab"
      background-color="indigo"
      dark
      slider-color="purple"
    >
      <v-tab
        v-for="tab in tabs"
        :key="tab"
        @click="selected = tab;"
        :class="['tab-btn', { active: selected === tab }]"
      >
        {{ tab }}
      </v-tab>
    </v-tabs>
    <component :is="selected" class="tab"></component>
  </v-card>
</section>
</template>

<script>
import HandleCredit from '~/components/organisms/HandleCredit.vue'
import TradeNotification from '~/components/organisms/TradeNotification'
  export default {
    name: 'Csv',
    components: {
      HandleCredit,
      TradeNotification
    },
    data () {
      return {
        tabs: ["HandleCredit", "TradeNotification"],
        selected: "HandleCredit"
      }
    },
  }
</script>

<style scoped>
.container {
  min-height: 90vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
