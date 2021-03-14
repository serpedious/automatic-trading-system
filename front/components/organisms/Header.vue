<template>
  <div>
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            Dashboard
          </v-list-item-title>
          <v-list-item-subtitle>
          chose categories
        </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
        dense
        nav
      >
        <v-list-item
          v-for="category in categories"
          :key="category.title"
          :to="category.to"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ category.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ category.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      color="black"
      dark
    >
      <template v-slot:img="{ props }">
        <v-img
          v-bind="props"
          gradient="to top right, rgba(100,115,201,.7), rgba(25,32,72,.7)"
        ></v-img>
      </template>

      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Automatic Trading system</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-menu
        bottom
        left
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            color="blue"
            v-bind="attrs"
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="(item, i) in items"
            :key="i"
          >
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data: () => ({ 
    drawer: null,
    items: [
      { title: 'Account', icon: 'mdi-chart-line', to: '/dashboard/chart' },
      { title: 'Github', icon: 'mdi-currency-usd', to: '' },
    ],
    categories: [
      { title: 'Chart', icon: 'mdi-chart-line', to: '/dashboard/chart' },
      { title: 'Profit', icon: 'mdi-currency-usd', to: '/dashboard/profit' },
      { title: 'Compare', icon: 'mdi-arrow-left-right', to: '/dashboard/compare' },
      { title: 'Memo', icon: 'mdi-format-list-checks', to: '/dashboard/memo' },
      { title: 'Alert', icon: 'mdi-bell-outline', to: '/dashboard/alert' },
      { title: 'CSV', icon: 'mdi-file-delimited-outline', to: '/dashboard/csv' }
    ],
  }),
}
</script>

<style>
.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 300;
  font-size: 60px;
  color: rgb(41, 38, 38);
  letter-spacing: 1px;
}
</style>

