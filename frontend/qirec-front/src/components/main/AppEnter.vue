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
        <button type="submit" class="btn signin">Sign ip!</button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'AppEnter',
  data() {
    return {
      email: '',
      password: ''
    };
  },
  created() {
    this.loadState();
  },
  methods: {
    async enterUser() {

      try {
        const response = await axios.post('/signin', {
          email: this.email,
          password: this.password
        });

        if (response.status === 200) {
          console.log('Sign up successful:', response.data.message);

        }
      } catch (error) {
        console.error('Error during sign up:', error.response ? error.response.data : error.message);
      }
    },
    saveState() {
      const state = {
        email: this.email,
        password: this.password
      };
      localStorage.setItem('enterState', JSON.stringify(state));
    },
    loadState() {
      const savedState = localStorage.getItem('enterState');
      if (savedState) {
        const state = JSON.parse(savedState);
        this.email = state.email;
        this.password = state.password;
      }
    }
  },
  watch: {
    email() {
      this.saveState();
    },
    password() {
      this.saveState();
    }
  }
};
</script>

<style scoped>
@import '../../../public/css/styles.css';
</style>