<template>
    <footer>
        <div v-if="!core.footer" class="footer text-center mb-4 p-2">
            <a href="https://github.com/statping/statping" target="_blank">
                Statping {{core.version}} made with <font-awesome-icon icon="heart"/>
            </a> |
            <router-link :to="core.logged_in ? '/dashboard' : '/login'">Dashboard</router-link>
        </div>
        <div v-else class="footer text-center mb-4 p-2" v-html="core.footer"></div>
    </footer>
</template>

<script>
  import Dashboard from "../../pages/Dashboard";

  export default {
  name: 'Footer',
  components: {
    Dashboard
  },
  props: {
    version: String,
    logged_in: Boolean
  },
  computed: {
      core() {
        return this.$store.getters.core
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
