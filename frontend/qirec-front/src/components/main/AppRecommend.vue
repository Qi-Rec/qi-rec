<template>
  <div class="rec-song">
    <div class="recommend">
      <form @submit.prevent="recommendSong" class="recommend-form">
        <div class="form-left-decoration"></div>
        <div class="form-right-decoration"></div>
        <div class="circle"></div>
        <div class="form-inner">
          <h3>Recommend a song</h3>
          <input type="text" v-model="playlistLink" placeholder="Your Link" required>
          <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
          <input type="submit" class="btn recommend" value="Recommend">
        </div>
      </form>
    </div>
    <RecommendedSong :recommended-song="recommendedSong"/>
    <button @click="getRecommendationHistory" class="btn-history">View Recommendation History</button>
  </div>
</template>

<script>
import RecommendedSong from "@/components/main/RecommendedSong.vue";

export default {
  name: 'Recommend',
  components: {RecommendedSong},
  data() {
    return {
      playlistLink: '',
      recommendedSong: null, // Хранение информации о песне
      errorMessage: '', // Сообщение об ошибке
    };
  },
  created() {
    this.loadState();
  },
  methods: {
    async recommendSong() {
      // Регулярное выражение для проверки ссылки на плейлист в Spotify
      const spotifyPlaylistRegex = /https:\/\/open\.spotify\.com\/playlist\/[a-zA-Z0-9]{22}\/*/;

      if (!spotifyPlaylistRegex.test(this.playlistLink)) {
        this.errorMessage = 'Spotify playlist link required';
        return;
      }

      this.errorMessage = ''; // Очистить сообщение об ошибке, если ссылка правильная

      try {
        const response = await this.$axios.post('/recommendation', {
          playlist_link: this.playlistLink
        });
        if (response.status >= 200 && response.status < 300) {
          this.recommendedSong = response.data;
          console.log('Recommended Song:', this.recommendedSong);
        }
      } catch (error) {
        if (error.response.status >= 400 && error.response.status < 500) {
          this.errorMessage = 'You must be authorized';
          console.error('User not authorized');
        } else {
          console.error('Error recommending song:', error);
        }
      }
    },
    async getRecommendationHistory() {
      try {
        const response = await this.$axios.get('/recommendation/history');

        if (response.status >= 200 && response.status < 300) {
          this.$emit('change-component', {component: 'AppRecomHistory', data: response.data.songs});
        }
      } catch (error) {
        console.error('Error fetching recommendation history:', error.response ? error.response.data : error.message);
      }
    },
    saveState() {
      if (!this.token) {
        return;
      }

      const state = {
        playlistLink: this.playlistLink,
        recommendedSong: this.recommendedSong,
        errorMessage: this.errorMessage
      };
      localStorage.setItem('recommendState', JSON.stringify(state));
    },
    loadState() {
      const savedState = localStorage.getItem('recommendState');
      if (savedState != null) {
        const state = JSON.parse(savedState);
        this.playlistLink = state.playlistLink;
        this.recommendedSong = state.recommendedSong;
        this.errorMessage = state.errorMessage;
      }
    }
  },
  watch: {
    playlistLink() {
      this.saveState();
    },
    recommendedSong() {
      this.saveState();
    },
    errorMessage() {
      this.saveState();
    }
  }
};
</script>

<style scoped>
@import '../../../public/css/styles.css';

.error-message {
  color: red;
  margin-top: 10px;
}

.btn-history {
  background-color: white;
  margin: 1.5rem 0;
  color: var(--sweet-purple);
  border: none;
  border-radius: 20px;
  padding: 10px 20px;
  cursor: pointer;
  width: 100%;
}

.btn-history:hover {
  background-color: lightgray;
}

.btn-history:active {
  background-color: lightgray;
  color: var(--sweet-purple);
  transform: scale(0.95);
}

</style>