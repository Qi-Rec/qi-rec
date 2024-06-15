<template>
  <div class="middle">
    <AppRecommend />
  </div>
</template>

<script>
import AppRecommend from './main/AppRecommend.vue';



export default {
  name: 'AppMiddle',
  data() {
    return {
      currentComponent: 'AppRecommend'
    };
  },
  components: {
    AppRecommend,
  },
  methods: {
    changeComponent(component) {
      this.currentComponent = component;
    }
  }
};
</script>

<style scoped>
@import '../../public/css/styles.css';
</style>