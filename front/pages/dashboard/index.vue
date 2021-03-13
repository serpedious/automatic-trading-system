<template>
  <div class="home-container">
    <section class="container">
      <div>
        <app-logo/>
        <h1>This is dashboard index</h1>
      </div>
    </section>
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'

export default {
  name: 'Dashboard',
  components: {
    AppLogo
  },
  data () {
    return {
      msg: 'it is lambda batch test3'
    }
  },
  methods: {

  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.home-container {
  min-height: 100%;
  width: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
}
</style>

