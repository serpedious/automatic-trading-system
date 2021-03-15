<template>
 <v-simple-table dense>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Side
          </th>
          <th class="text-left">
            Amount
          </th>
          <th class="text-left">
            Date
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in desserts"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.side }}</td>
          <td>{{ item.amount }}</td>
          <td>{{ item.date}}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
export default {
  name: 'Profit',
  data () {
      return {
        desserts: [
          {
            name: 'Frozen Yogurt',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Ice cream sandwich',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Eclair',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Cupcake',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Gingerbread',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Jelly bean',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Lollipop',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Honeycomb',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'Donut',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
          },
          {
            name: 'KitKat',
            side: 'SELL',
            amount: 100,
            date: '2020/08/23',
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