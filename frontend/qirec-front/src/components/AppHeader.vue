<template>
  <header class="header">
    <div class="header__content">
      <div class="logo-title" @click="goToRecommend">
        <img src="../assets/qi-recLogo.png" alt="Logo" class="logo"/>
        <h1 class="title">Qi-Rec</h1>
      </div>
      <div class="auth-buttons">
        <template v-if="authorized">
          <button @click="logout" class="btn logout">Logout</button>
        </template>
        <template v-else>
          <button @click="goToRegister" class="btn register">Sign Up</button>
          <button @click="goToSignIn" class="btn login">Sign In</button>
        </template>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'AppHeader',
  props: ['authorized'],
  methods: {
    goToRecommend() {
      this.$emit('change-component', 'AppRecommend');
    },
    goToRegister() {
      this.$emit('change-component', 'AppRegister');
    },
    goToSignIn() {
      this.$emit('change-component', 'AppEnter');
    },
    async logout() {
      console.log(123);
      await this.$axios.post('/logout');
      localStorage.setItem('authorized', "false");
      this.goToRecommend();
      location.reload();
    }
  }
};
</script>

<style scoped>
@import '../../public/css/styles.css';

.logout {
  background-color: var(--sweet-purple);
  color: black;
  border: none;
  border-radius: 5px;
}
</style>