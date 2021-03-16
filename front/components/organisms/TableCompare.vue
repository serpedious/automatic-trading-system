<template>
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
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>