<template>
  <div class="__form">
    <form @submit.prevent="preSubmit">
      <slot/>
    </form>
  </div>
</template>

<script>
export default {
  name: 'form',
  props: {
    name: {
      type: String,
      default: null
    },
    submit: {
      type: Function,
      default: () => null
    }
  },
  methods: {
    preSubmit () {
      this.submit()
    }
  },
  data () {
    return {}
  },
  mounted () {
  }
}

</script>
