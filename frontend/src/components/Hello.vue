<template>
  <div class="hello">
    <router-link :to="{'name': 'Dashboard'}" class="btn btn-primary">Dashboard</router-link>
  </div>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {}
  }
}
</script>

<style scoped>
.hello {
  width: 100%;
  margin-top: 300px;
  text-align: center;
}
</style>
