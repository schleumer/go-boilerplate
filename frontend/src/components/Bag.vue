<template>
  <div class="__bag form-group">
    <slot />
    <div v-if="type === 'message'">
      <div v-for="message in messages" class="text-danger">
        {{ message.message }}
      </div>
    </div>
    <div v-if="type === 'alert'">
      <div v-for="message in messages" class="alert alert-danger">
        {{ message.message }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'bag',
  props: {
    name: {
      type: String,
      default: null
    },
    type: {
      type: String,
      default: 'message'
    }
  },
  data () {
    return {
      formName: null
    }
  },
  mounted () {
    const parentForm = this.$nextParent('form')

    if (parentForm) {
      this.formName = parentForm.name
    }
  },
  computed: {
    messages () {
      return this.$store.state.chaos.messages.filter(_ => _.form === this.formName && _.field === this.name)
    }
  }
}

</script>
