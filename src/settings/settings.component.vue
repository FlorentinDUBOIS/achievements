<template>
  <div class="settings-container">
    <md-list>
      <md-subheader>Données</md-subheader>
      <md-list-item @click="removeData()">
        <md-icon>data_usage</md-icon>
        <span class="md-list-item-text">Supprimer les données locales</span>
      </md-list-item>
    </md-list>

    <md-snackbar md-position="center" md-duration.number="5000" :md-active.sync="local_data" md-persistent>
      <span>Les données locales ont été supprimées</span>
      <md-button class="md-accent" @click="local_data = false">Fermer</md-button>
    </md-snackbar>
  </div>
</template>

<style lang="scss" scoped>
.settings-container {
  width: 100%;
  height: 100%;
}
</style>


<script>
export default {
  name: 'SwSettings',
  data () {
    return {
      local_data: false
    }
  },

  methods: {
    removeData () {
      this.local_data = true
    }
  }
}
</script>
