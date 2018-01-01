<template>
  <div ref="container" class="container">
    <canvas ref="canvas"></canvas>

    <md-snackbar md-position="center" :md-duration="duration" :md-active.sync="snackbar" md-persistent>
      <span>La caméra est requise pour lire les qrcodes</span>
      <md-button class="md-accent" @click="start()">Réessayer</md-button>
    </md-snackbar>
  </div>
</template>

<style lang="scss" scoped>
.container {
  width: 100%;
  height: 100%;
}
</style>

<script>
import {Camera} from './camera.service'
import {Scheduler} from 'rxjs'

export default {
  name: 'Camera',
  data () {
    return {
      video: document.createElement('video'),
      snackbar: false,
      duration: 5000
    }
  },

  mounted () {
    // wait after the component to be fully rendered
    this.$nextTick(() => {
      const {video} = this
      const {canvas, container} = this.$refs

      const width = container.clientWidth
      const height = container.clientHeight

      video.setAttribute('width', width)
      video.setAttribute('height', height)

      canvas.setAttribute('width', width)
      canvas.setAttribute('height', height)
    })

    this.start()
  },

  beforeDestroy () {
    if (this.subscription) {
      this.subscription.unsubscribe()
    }
  },

  methods: {
    refresh (vm) {
      return function () {
        const {video} = vm
        const {canvas} = vm.$refs
        const context = canvas.getContext('2d')

        context.drawImage(video, 0, 0, canvas.getAttribute('width'), canvas.getAttribute('height'))

        this.schedule()
      }
    },

    start () {
      Camera
        .stream({video: true})
        .then(stream => {
          const {video} = this

          video.setAttribute('src', stream)
          video.addEventListener('canplay', event => {
            video.play()
          })

          this.subscription = Scheduler
            .animationFrame
            .schedule(this.refresh(this))
        })
        .catch(() => {
          this.snackbar = true
        })
    }
  }
}
</script>
