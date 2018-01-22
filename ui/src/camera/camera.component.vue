<template>
  <div ref="container" class="camera-container">
    <div v-if="!allowed" class="state">
      <md-empty-state
        class="md-accent"
        md-rounded
        md-icon="error_outline"
        md-label="Caméra inaccessible"
        :md-size="360"
        md-description="Veuillez activer la caméra afin de scanner les qrcodes">
      </md-empty-state>
    </div>

    <div v-if="allowed">
      <granite-qrcode-scanner ref="scanner"></granite-qrcode-scanner>
    </div>

    <md-snackbar md-position="center" :md-duration="5000" :md-active.sync="snackbar" md-persistent>
      <span>La caméra est requise pour lire les qrcodes</span>
      <md-button class="md-accent">Réessayer</md-button>
    </md-snackbar>
  </div>
</template>

<style lang="scss" scoped>
.camera-container {
  width: 100%;
  height: 100%;

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
// import {Camera} from './camera.service'

export default {
  name: 'SwCamera',
  data () {
    return {
      allowed: true,
      snackbar: false
    }
  },

  mounted () {
    // wait after the component to be fully rendered
    this.$nextTick(() => {
      const {scanner, container} = this.$refs

      const width = container.clientWidth
      const height = container.clientHeight

      scanner.setAttribute('active', 'true')
      scanner.setAttribute('continuous', 'true')
      scanner.setAttribute('width', width)
      scanner.setAttribute('height', height)
      scanner.addEventListener('qrcode-decoded', event => {
        console.log(event.detail)
      })
    })
  }
}
</script>
