<template>
  <div class="messenger-container">
    <div class="flex">
      <div class="state">
        <md-empty-state
          md-icon="chat_bubble_outline"
          md-label="Aucun message">
        </md-empty-state>
      </div>
    </div>

    <md-toolbar class="md-primary" md-elevation="0">
      <div class="md-toolbar-row">
        <md-autocomplete
          v-model="message"
          :md-options="options"
          md-layout="box">
          <label>Message</label>
        </md-autocomplete>
      </div>
    </md-toolbar>
  </div>
</template>

<style lang="scss" scoped>
.messenger-container {
  display: flex;
  flex-direction: column;

  width: 100%;
  height: 100%;

  & > .flex {
    flex: 1;

    max-height: 100%;
    overflow: auto;

    & > .state {
      width: 100%;
      height: 100%;

      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }
  }
}
</style>

<script>
export default {
  name: 'SwMessenger',
  data () {
    return {
      size: 360,
      message: '',
      options: [],
      messages: []
    }
  }
}
</script>
