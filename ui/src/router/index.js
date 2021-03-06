import Router from 'vue-router'

import About from '../about/about.component'
import Camera from '../camera/camera.component'
import Home from '../home/home.component'
import Settings from '../settings/settings.component'
import Achievements from '../achievements/achievements.component'
import Achievement from '../achievements/achievement.component'
import Messenger from '../messenger/messenger.component'
import Staff from '../staff/staff.component'
import Ranking from '../ranking/ranking.component'

export default new Router({
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/about',
      component: About
    },
    {
      path: '/camera',
      component: Camera
    },
    {
      path: '/settings',
      component: Settings
    },
    {
      path: '/achievements',
      component: Achievements
    },
    {
      path: '/achievements/:id',
      component: Achievement,
      props: true
    },
    {
      path: '/messenger',
      component: Messenger
    },
    {
      path: '/ranking',
      component: Ranking
    },
    {
      path: '/staff',
      component: Staff
    }
  ]
})
