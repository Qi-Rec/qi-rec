<template>
  <div id="app">
    <AppHeader :authorized="authorized" @change-component="changeComponent"/>
    <AppMiddle :currentComponent="currentComponent" @change-component="changeComponent"/>
    <AppFooter/>
  </div>
</template>

<script>
import AppHeader from './components/AppHeader.vue';
import AppMiddle from './components/AppMiddle.vue';
import AppFooter from './components/AppFooter.vue';

export default {
  name: 'App',
  components: {
    AppHeader,
    AppMiddle,
    AppFooter
  },
  data() {
    return {
      currentComponent: 'AppRecommend',
      authorized: localStorage.getItem('authorized') === 'true'
    };
  },
  created() {
    this.loadState();
  },
  methods: {
    changeComponent(component) {
      console.log("component changed")
      this.currentComponent = component;
      this.saveState();
    },
    saveState() {
      const state = {
        currentComponent: this.currentComponent
      };
      localStorage.setItem('appState', JSON.stringify(state));
    },
    loadState() {
      const savedState = localStorage.getItem('appState');
      if (savedState) {
        const state = JSON.parse(savedState);
        this.currentComponent = state.currentComponent;
      }
    }
  }
};
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
</style>