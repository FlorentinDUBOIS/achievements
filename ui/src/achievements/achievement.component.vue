<template>
  <div class="achievement-container">
    <md-card class="mt">
      <md-card-header>
        <md-card-header-text>
          <div class="md-title">{{ achievement.name }}</div>
          <div class="md-subtitle">{{ achievement.description }}</div>
        </md-card-header-text>

        <md-card-media md-medium>
          <img v-bind:src="id | avatar" alt="People">
        </md-card-media>
      </md-card-header>
    </md-card>

    <md-card>
      <md-card-header>
        <md-card-header-text>
          <div class="md-title">QrCode du Succès</div>
        </md-card-header-text>
      </md-card-header>

      <md-card-media md-medium id="qrcode" >
        <granite-qrcode-generator
          v-bind:data="id"
          mode="alphanumeric"
          modulesize="10"
          auto>
        </granite-qrcode-generator>
      </md-card-media>
    </md-card>
  </div>
</template>

<style lang="scss" scoped>
.achievement-container {
  & span {
    text-overflow: ellipsis;
  }

  & .mt {
    margin-top: 8px;
  }

  & #qrcode {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>

<script>
import {Avatar} from '../avatar/avatar.service'
import {Achievements} from './achievements.service'

export default {
  name: 'SwAchievement',
  props: ['id'],
  data () {
    return {
      achievement: {}
    }
  },

  beforeMount () {
    Achievements
      .get(this.id)
      .then(achievement => { this.achievement = achievement })
      .catch(console.error)
  },

  filters: {
    avatar: function (id) {
      return new Avatar(id || '').generate()
    }
  }
}
</script>
