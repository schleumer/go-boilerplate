<template>
  <div class="__app">
    <div v-if="!$store.state.system.booted">
      Tamo carregando as parada, guenta aí.
    </div>
    <div v-if="$store.state.system.booted">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
  created () {
    this.$store.dispatch('system:boot')
  }
}

</script>

<style></style>
