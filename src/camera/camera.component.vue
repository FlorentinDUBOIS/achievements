<template>
  <div ref="container" class="container">
    <canvas ref="canvas"></canvas>
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

export default {
  name: 'Camera',
  data () {
    return {
      video: document.createElement('video')
    }
  },

  mounted () {
    Camera
      .stream({video: true})
      .then(stream => {
        const {video} = this

        video.setAttribute('src', stream)
        video.addEventListener('canplay', event => {
          video.play()
        })

        requestAnimationFrame(this.refresh)
      })
      .catch(console.error)
  },

  methods: {
    refresh () {
      const {video} = this
      const {canvas, container} = this.$refs
      const context = canvas.getContext('2d')

      const width = container.clientWidth
      const height = container.clientHeight

      canvas.setAttribute('width', width)
      canvas.setAttribute('height', height)

      context.drawImage(video, 0, 0, width, height)

      requestAnimationFrame(this.refresh)
    }
  }
}
</script>
