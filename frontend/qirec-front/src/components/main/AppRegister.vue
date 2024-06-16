<template>
  <div class="recommend">
    <form @submit.prevent="registerUser">
      <div class="form-left-decoration"></div>
      <div class="form-right-decoration"></div>
      <div class="circle"></div>
      <div class="form-inner">
        <h3>Register</h3>
        <div class="form-group">
          <input type="email" id="email" v-model="email" placeholder="Your Email" required>
        </div>
        <div class="form-group">
          <input type="password" id="password" v-model="password" placeholder="Your Password" required>
        </div>
        <button type="submit" class="btn signin">Sign up!</button>
      </div>
    </form>
    <div v-if="notification" class="notification">{{ notification }}</div>
  </div>
</template>

<script>
export default {
  name: 'AppRegister',
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
    async registerUser() {
      try {
        console.log("sended requerst")

        const response = await this.$axios.post('/signup', {
          email: this.email,
          password: this.password
        });

        console.log("recieved fresponce from server", response.data)

        if (response.status >= 200 && response.status < 300) {
          this.notification = "You've been registered";
          console.log("recieved fresponce from server", response.data);
          localStorage.setItem('authorized', "true");
          this.$emit('change-component', 'AppRecommend');
           // location.reload();
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
      localStorage.setItem('registerState', JSON.stringify(state));
    },
    loadState() {
      const savedState = localStorage.getItem('registerState');
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

.notification {
  color: green;
  margin-top: 10px;
}
</style>