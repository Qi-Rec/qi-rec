<template>
  <div id="app">
    <AppHeader @change-component="changeComponent" />
    <AppMiddle ref="middle" />
    <AppFooter />
  </div>
</template>

<script>
import AppHeader from './components/AppHeader.vue';
import AppMiddle from './components/AppMiddle.vue';
import AppFooter from './components/AppFooter.vue';
import axios from "axios"

export default {
  name: 'App',
  components: {
    AppHeader,
    AppMiddle,
    AppFooter
  },

  methods: {
    changeComponent(component) {
      this.$refs.middle.changeComponent(component);
    }
  },
  beforeCreate() {
    this.$root.$on("onChangePage", (page) => this.page = page);

    this.$root.$on("onRecommendSong", (link) => {
      axios.post("/api/1/recommend", {
        link
      }).then(response => {
        this.$root.$emit("onRecommendSongSuccess", response.data);
      }).catch(error => {
        this.$root.$emit("onRecommendSongError", error.response.data);
      });
    });

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