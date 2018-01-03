<template>
  <div class="home-container">
    <md-empty-state
      class=""
      md-rounded
      :md-size="size"
      :md-icon="icon"
      :md-label="label"
      :md-description="description">
        <md-button class="md-dense md-primary md-raised" v-if="!signin">
          Se connecter
        </md-button>
    </md-empty-state>
  </div>
</template>

<style lang="scss" scoped>
.home-container {
  width: 100%;
  height: 100%;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>

<script>
export default {
  name: 'SwHome',
  data () {
    return {
      size: 360,
      signin: false,
      icon: 'account_box',
      label: 'Bonjour !',
      description: 'Connexion via le qrcode de votre badge du startup week-end'
    }
  }
}
</script>
