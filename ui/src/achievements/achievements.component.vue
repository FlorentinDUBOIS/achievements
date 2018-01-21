<template>
  <div class="achievements-container">
    <div v-if="error" class="state">
        <md-empty-state
          class="md-accent"
          md-rounded
          md-icon="error_outline"
          md-label="Succès indisponibles"
          :md-size="360"
          md-description="Veuillez vérifier votre connexion à Internet">
        </md-empty-state>
    </div>

    <md-list v-if="!error" class="">
      <md-subheader>Succès</md-subheader>

      <md-list-item @click="go(`/achievements/${achievement.id}`)" v-bind:key="achievement.id" v-for="achievement in achievements">
        <md-avatar>
          <img :src="achievement.id | avatar" alt="avatar">
        </md-avatar>

        <div class="md-list-item-text">
          <span>{{ achievement.name }}</span>
          <span class="">{{ achievement.description }}</span>
        </div>
      </md-list-item>
    </md-list>
  </div>
</template>

<style lang="scss" scoped>
.achievements-container {
  width: 100%;
  height: 100%;

  & span {
    text-overflow: ellipsis;
  }

  & > .state {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    width: 100%;
    height: 100%;
  }
}
</style>

<script>
import {Avatar} from '../avatar/avatar.service'
import {Achievements} from './achievements.service'
import {RouterService} from '../router/router.service'

export default {
  name: 'SwAchievements',
  data () {
    return {
      error: false,
      achievements: []
    }
  },

  beforeMount () {
    Achievements
      .find()
      .then(achievements => { this.achievements = achievements })
      .catch(err => {
        this.error = true
        console.error(err)
      })
  },

  methods: {
    go (path) {
      RouterService.push({ path })
    }
  },

  filters: {
    avatar: function (id) {
      return new Avatar(id).generate()
    }
  }
}
</script>
