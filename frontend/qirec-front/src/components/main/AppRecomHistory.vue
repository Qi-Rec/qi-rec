<template>
  <div class="recom-history">
    <h3>Recommendation History</h3>
    <div v-for="song in recentSongs" :key="song.id" class="song-frame">
      <iframe
        style="border-radius:20px"
        :src="'https://open.spotify.com/embed/track/' + song.id + '?utm_source=generator'"
        width="100%"
        height="300"
        frameborder="0"
        allowfullscreen=""
        allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
        loading="lazy">
      </iframe>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AppRecomHistory',
  props: {
    recentSongs: {
      type: Array,
      required: true
    }
  }
};
</script>

<style scoped>
.recom-history {
  padding: 20px;
}

.song-frame {
  margin-bottom: 20px;
}
</style>