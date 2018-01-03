import { store } from '../app.store'
import ActionType from './router.type'

export class RouterService {
  static push (route) {
    store.dispatch({
      type: ActionType.PUSH,
      route
    })
  }

  static route () {
    const { router } = store.getState()

    return router
  }
}
