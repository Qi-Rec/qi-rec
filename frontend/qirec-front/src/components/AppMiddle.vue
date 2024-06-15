<template>
  <div class="middle">
    <component :is="currentComponent" />
  </div>
</template>

<script>
import AppRecommend from './main/AppRecommend.vue';
import AppRegister from './main/AppRegister.vue';
import AppEnter from './main/AppEnter.vue';

export default {
  name: 'AppMiddle',
  props: {
    currentComponent: {
      type: String,
      required: true
    }
  },
  components: {
    AppRecommend,
    AppRegister,
    AppEnter
  }
};
</script>

<style scoped>
@import '../../public/css/styles.css';
</style>