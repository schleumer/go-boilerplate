import * as types from './mutation-types'

export default {
  [types.INC] (state) {
    state.count = state.count + 1
  }
}
