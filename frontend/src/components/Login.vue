<template>
  <div class="__login">
    <cool-form name="login" :submit="doLogin">
      <div class="row justify-content-center login-holder">
        <div class="col-xl-3 col-lg-4 col-md-6 col-sm-8 col-xs-12">
          <div class="card">
            <div class="card-header">
              Opa, pera lá, você precisa entrar!
            </div>
            <div class="card-block">
              <div v-if="$store.state.auth.signing" class="text-center">
                Conectando...
              </div>
              <div v-if="!$store.state.auth.signing">
                <bag type="alert" />

                <bag name="username">
                  <label>Login</label>
                  <input type="text" class="form-control" v-model="form.username" />
                </bag>

                <bag name="password">
                  <label>Senha</label>
                  <input type="password" class="form-control" v-model="form.password" />
                </bag>

                <div>
                  <button class="btn btn-primary btn-block">ENTRAR</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </cool-form>
  </div>
</template>

<script>
import Bag from './Bag'
import CoolForm from './Form'
import Test from './Test'

export default {
  components: { Bag, CoolForm, Test },
  name: 'login',
  data () {
    return {
      form: {
        username: 'admin',
        password: 'senha@123'
      }
    }
  },
  methods: {
    doLogin () {
      this.$store.dispatch('auth:do-login', this.form)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login-holder {
  margin-top: 100px
}
</style>
