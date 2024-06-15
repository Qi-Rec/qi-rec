<template>
  <div v-if="recommendedSong" class="song-info">
    <iframe style="border-radius:20px"
            :src="'https://open.spotify.com/embed/track/' + recommendedSong.id + '?utm_source=generator'" width="100%"
            height="232" frameBorder="0" allowfullscreen=""
            allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
  </div>
</template>
<script>
export default {
  name: 'RecommendedSong',
  props: {
    recommendedSong: {}
  }
}
</script>
<style scoped>
@import '../../../public/css/styles.css';

</style>