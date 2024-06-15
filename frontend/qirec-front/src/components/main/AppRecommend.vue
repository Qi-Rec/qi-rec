<template>
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

    <div v-if="recommendedSong" class="song-info">
      <h4>Recommended Song</h4>
      <p><strong>Name:</strong> {{ recommendedSong.name }}</p>
      <p><strong>Artist:</strong> {{ recommendedSong.artist }}</p>
      <img :src="recommendedSong.cover_link" alt="Song Cover" class="song-cover" />
      <p><a :href="recommendedSong.song_link" target="_blank" class="btn song-link">Listen on Spotify</a></p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Recommend',
  data() {
    return {
      playlistLink: '',
      recommendedSong: null, // Хранение информации о песне
      errorMessage: '', // Сообщение об ошибке
    };
  },
  methods: {
    async recommendSong() {
      // Регулярное выражение для проверки ссылки на плейлист в Spotify
      const spotifyPlaylistRegex = /^https:\/\/open\.spotify\.com\/playlist\/[a-zA-Z0-9]+\?si=$[a-zA-Z0-9]+/;

      if (!spotifyPlaylistRegex.test(this.playlistLink)) {
        this.errorMessage = 'Spotify playlist link required';
        return;
      }

      this.errorMessage = ''; // Очистить сообщение об ошибке, если ссылка правильная

      try {
        const response = await this.$axios.post('/recommendation', {
          playlist_link: this.playlistLink
        });
        this.recommendedSong = response.data.song;
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

.song-info {
  margin-top: 20px;
  padding: 20px;
  background: #fff;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.song-info h4 {
  margin-top: 0;
}

.song-cover {
  max-width: 100%;
  border-radius: 5px;
}

.btn.song-link {
  background-color: var(--sweet-purple);
  color: white;
  text-decoration: none;
  padding: 10px 20px;
  border-radius: 5px;
  display: inline-block;
  margin-top: 10px;
}

.btn.song-link:hover {
  background-color: var(--sweet-purple)
}
</style>