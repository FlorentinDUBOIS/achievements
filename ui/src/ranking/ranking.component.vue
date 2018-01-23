<template>
  <div class="ranking-container">
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
      <md-subheader>Rang</md-subheader>

      <md-list-item v-bind:key="rank.user_id" v-for="rank in ranks">
        <md-avatar>
          <img :src="rank.user_id | avatar" alt="avatar">
        </md-avatar>

        <div class="md-list-item-text">
          <span>{{ rank.user_id }}</span>
          <span>Nombre: {{ rank.count }}</span>
        </div>
      </md-list-item>
    </md-list>
  </div>
</template>

<style lang="scss" scoped>
.ranking-container {
  width: 100%;
  height: 100%;

  & .state {
    width: 100%;
    height: 100%;

    display: flex;
    flex-direction: column;

    justify-content: center;
    align-items: center;
  }
}
</style>

<script>
import {Avatar} from '../avatar/avatar.service'
import {RankingService} from './ranking.service'

export default {
  name: 'SwRanking',
  data () {
    return {
      error: true,
      ranks: []
    }
  },

  beforeMount () {
    RankingService
      .get()
      .then(ranks => {
        this.error = false
        this.ranks = ranks
      })
      .catch(() => { this.error = true })
  },

  mounted () {
    this.interval = setInterval(() => {
      RankingService
        .get()
        .then(ranks => {
          this.error = false
          this.ranks = ranks
        })
        .catch(() => { this.error = true })
    }, 5000)
  },

  beforeDestroy () {
    clearInterval(this.interval)
  },

  filters: {
    avatar: function (id) {
      return new Avatar(id || '').generate()
    }
  }
}
</script>
