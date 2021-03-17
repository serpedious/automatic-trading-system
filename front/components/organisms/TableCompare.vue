<template>
<section class="container">
 <v-card class="mx-auto mt-2 pa-4" width="800px">
    <v-card-title class="pb-5">
      <h3>Compare</h3>
    </v-card-title>
 <v-simple-table dense>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            bitFlyer
          </th>
          <th class="text-left">
            Coincheck
          </th>
          <th class="text-left">
            DMM Bitcoin
          </th>
          <th class="text-left">
            SBI VC Trade
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in desserts"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.amount }}</td>
          <td>{{ item.amount}}</td>
          <td>{{ item.amount}}</td>
          <td>{{ item.amount}}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
 </v-card>
</section>
</template>

<script>
export default {
  name: 'History',
  data () {
      return {
        desserts: [
          {
            name: 'Frozen Yogurt',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
          {
            name: 'Ice cream sandwich',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
          {
            name: 'Eclair',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
          {
            name: 'Cupcake',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
          {
            name: 'Gingerbread',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
          {
            name: 'KitKat',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
            amount: 100,
          },
        ],
      }
    },
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