<template>
  <header class="header">
    <div class="header__content">
      <div class="logo-title" @click="goToRecommend">
        <img src="../assets/qi-recLogo.png" alt="Logo" class="logo" />
        <h1 class="title">Qi-Rec</h1>
      </div>
      <div class="auth-buttons">
        <button @click="goToRegister" class="btn register">Sign Up</button>
        <button @click="goToSignIn" class="btn login">Sign In</button>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'AppHeader',
  methods: {
    goToRecommend() {
      this.$emit('change-component', 'AppRecommend');
    },
    goToRegister() {
      this.$emit('change-component', 'AppRegister');
    },
    goToSignIn() {
      this.$emit('change-component', 'AppEnter');
    }
  }
};
</script>

<style scoped>
@import '../../public/css/styles.css';
</style>