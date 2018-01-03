import ActionType from './router.type'

export function router (router) {
  return function (state = {}, action) {
    switch (action.type) {
      case ActionType.PUSH: {
        router.push(action.route)

        return Object.assign({}, state, action.route)
      }

      default: {
        return state
      }
    }
  }
}
