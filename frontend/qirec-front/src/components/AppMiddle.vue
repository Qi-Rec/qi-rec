<template>
  <div class="middle">
    <component :is="currentComponent" :custom-data="customData" @change-component="$emit('change-component', $event)"/>
  </div>
</template>

<script>
import AppRecommend from './main/AppRecommend.vue';
import AppRegister from './main/AppRegister.vue';
import AppEnter from './main/AppEnter.vue';
import AppRecomHistory from './main/AppRecomHistory.vue';

export default {
  name: 'AppMiddle',
  props: {
    currentComponent: {
      type: String,
      required: true
    },
    customData: {
      type: Object,
    }
  },
  components: {
    AppRecommend,
    AppRegister,
    AppEnter,
    AppRecomHistory
  }
};
</script>

<style scoped>
@import '../../public/css/styles.css';
</style>