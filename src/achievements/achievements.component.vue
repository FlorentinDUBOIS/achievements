<template>
  <md-list class="achievements-container">
    <md-subheader>Succès</md-subheader>

    <md-list-item v-bind:key="achievement.id" v-for="achievement in achievements">
      <md-avatar>
        <img :src="achievement.name | avatar" alt="avatar">
      </md-avatar>

      <div class="md-list-item-text">
        <span>{{ achievement.name }}</span>
        <span>{{ achievement.description }}</span>
      </div>
    </md-list-item>
  </md-list>
</template>

<style lang="scss" scoped>
.achievements-container {
  & span {
    text-overflow: ellipsis;
  }
}
</style>

<script>
import {Avatar} from '../avatar/avatar.service'
import {Achievements} from './achievements.service'

const achievements = new Achievements()

export default {
  name: 'SwAchievements',
  data () {
    return {
      achievements: []
    }
  },

  beforeMounted () {
    this
      .populate()
      .then(achievement => { this.achievements = achievements })
      .catch(console.log)
  },

  methods: {
    async populate () {
      await achievements.connect()
      await achievements.insert({name: 'foo', description: 'bar'})

      return achievements.find()
    }
  },

  filters: {
    avatar: function (name) {
      return new Avatar(name).generate()
    }
  }
}
</script>
