<template>
  <header class="header">
    <div class="header__content">
      <div class="logo-title">
        <img src="../assets/qi-recLogo.png" alt="Logo" class="logo" />
        <h1 class="title">Qi-Rec</h1>
      </div>
      <div class="auth-buttons">
        <button @click="$emit('change-component', 'Register')" class="btn register">Sign up</button>
        <button class="btn login">Sign in</button>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'AppHeader', // Изменено имя компонента
};
</script>

<style scoped>
@import '../../public/css/styles.css';
</style>