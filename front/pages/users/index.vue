<template>
  <section class="container">
    <div>
      <h1>usersのindexページです。</h1>
      <p>count={{count}}</p>
      <p>back to home?
        <router-link to="/">home</router-link>
      </p>
      <button @click="addCount">カウントアップ</button>
    </div>
  </section>
</template>

<script>
export default {
  computed: {
    count () { return this.$store.state.counter.count }
  },
  methods: {
    addCount (e) {
      this.$store.commit('counter/add')
    }
  }
}
</script>