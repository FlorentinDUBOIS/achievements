import { createStore, combineReducers } from 'redux'

import Router from './router/index'
import { router } from './router/router.reducer'

export const store = createStore(
  combineReducers({
    router: router(Router)
  })
)
