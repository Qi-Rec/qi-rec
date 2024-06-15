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
  methods: {
    async recommendSong() {
      // // Регулярное выражение для проверки ссылки на плейлист в Spotify
      // const spotifyPlaylistRegex = /^https:\/\/open\.spotify\.com\/playlist\/[a-zA-Z0-9]+\?si=$[a-zA-Z0-9]+/;
      //
      // if (!spotifyPlaylistRegex.test(this.playlistLink)) {
      //   this.errorMessage = 'Spotify playlist link required';
      //   return;
      // }

      this.errorMessage = ''; // Очистить сообщение об ошибке, если ссылка правильная

      try {
        const response = await this.$axios.post('/recommendation', {
          playlist_link: this.playlistLink
        });
        this.recommendedSong = response.data;
        console.log(response.data)
        console.log('Recommended Song:', this.recommendedSong);
      } catch (error) {
        console.error('Error recommending song:', error);
      }
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

</style>