<template>
  <div class="page-container">
    <md-app md-mode="reveal">
      <md-app-toolbar class="md-primary">
        <div class="md-toolbar-row">
          <div class="md-toolbar-section-start">
            <md-button class="md-icon-button" @click="toggle()">
              <md-icon>menu</md-icon>
            </md-button>

            <span class="md-title">SW Succès</span>
          </div>

          <div class="md-toolbar-section-end">
            <md-button class="md-icon-button">
              <md-icon>power_settings_new</md-icon>
            </md-button>
          </div>
        </div>
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="drawer">
        <md-list>
          <md-subheader>Paramètres</md-subheader>
          <md-list-item :to="item.to" v-bind:key="item.icon" v-for="item in settingsItems">
            <md-icon>{{ item.icon }}</md-icon>
            <span class="md-list-item-text">{{ item.label }}</span>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <div class="content">
          <router-view></router-view>
        </div>

        <md-bottom-bar class="md-accent" md-sync-route>
          <md-bottom-bar-item @click="go(item.to)" :md-label="item.label" :md-icon="item.icon" v-bind:key="item.icon" v-for="item in bottomItems"></md-bottom-bar-item>
        </md-bottom-bar>
      </md-app-content>
    </md-app>
  </div>
</template>

<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css?family=Roboto:300,400,500,700,400italic|Material+Icons");

.page-container {
  & > .md-app {
    height: 100vh;

    & .md-app-content {
      padding: 0;

      display: flex;
      flex-direction: column;

      max-height: calc(100vh - 56px);
      overflow: auto;

      & .content {
        max-height: calc(100vh - 56px);
        overflow: auto;

        position: relative;
        flex: 1;
      }
    }
  }

  & .md-bottom-bar {
    & .md-bottom-bar-item {
      max-width: none;
    }
  }
}
</style>

<script>
export default {
  name: 'App',
  data () {
    return {
      drawer: false,
      settingsItems: [{
        icon: 'settings',
        label: 'Paramètres',
        to: '/settings'
      }, {
        icon: 'code',
        label: 'A propos',
        to: '/about'
      }],
      bottomItems: [{
        icon: 'home',
        label: 'Home',
        to: '/'
      }, {
        icon: 'star',
        label: 'Succès',
        to: '/achievements'
      }, {
        icon: 'camera',
        label: 'Caméra',
        to: '/camera'
      }]
    }
  },

  methods: {
    toggle () {
      this.drawer = !this.drawer
    },

    go (path) {
      this.$router.push({ path })
    }
  }
}
</script>
