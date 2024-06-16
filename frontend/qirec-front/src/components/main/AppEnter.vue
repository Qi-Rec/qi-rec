<template>
  <div class="recommend">
    <form @submit.prevent="enterUser">
      <div class="form-left-decoration"></div>
      <div class="form-right-decoration"></div>
      <div class="circle"></div>
      <div class="form-inner">
        <h3>Enter</h3>
        <div class="form-group">
          <input type="email" id="email" v-model="email" placeholder="Your Email" required>
        </div>
        <div class="form-group">
          <input type="password" id="password" v-model="password" placeholder="Your Password" required>
        </div>
        <button type="submit" class="btn signin">Sign in!</button>
      </div>
    </form>
    <div v-if="notification" class="notification">{{ notification }}</div>
  </div>
</template>

<script>

export default {
  name: 'AppEnter',
  data() {
    return {
      email: '',
      password: '',
      notification: ''
    };
  },
  created() {
    this.loadState();
  },
  methods: {
    goToRecommend() {
      this.$emit('change-component', 'AppRecommend');
    },
    async enterUser() {
      try {
        const response = await this.$axios.post('/signin', {
          email: this.email,
          password: this.password
        });

        if (response.status >= 200 && response.status < 300) {
          this.notification = "Sign in successful";
          localStorage.setItem('authorized', "true");
          this.goToRecommend();
          location.reload();
        }
        if (response.status >= 400 && response.status < 500) {
          this.notification = "Sign in failed";
        }
      } catch (error) {
        console.error('Error during sign up:', error.response ? error.response.data : error.message);
        this.notification = 'Sign in failed';
      }
    },
    loadState() {
      const savedState = localStorage.getItem('enterState');
      if (savedState) {
        const state = JSON.parse(savedState);
        this.email = state.email;
        this.password = state.password;
      }
    }
  }
};
</script>

<style scoped>

.notification {
  color: darkviolet;
  text-align: center;
  margin-top: 10px;
}

@import '../../../public/css/styles.css';
</style>