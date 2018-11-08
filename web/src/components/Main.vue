<template>
  <div class="main">
    <div class="g-recaptcha" :data-sitekey="sitekey"></div>
    <button @click="verify">验证</button>

    <div class="console">
      {{ consoleMsg }}
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      sitekey: '6LcLcXkUAAAAAJniZZQramraMQ6rUjby0MA5d4dC',
      consoleMsg: ''
    }
  },
  methods: {
    verify () {
      let resp = window.grecaptcha.getResponse()
      axios.post('/api/verify',
        resp
      ).then(res => {
        this.consoleMsg = res.data
      })
      console.log(resp)
    }
  }
}
</script>

<style scoped>
  .main  {
    width: 240px;
    margin: 0 auto;
  }
  .console {
    height: 120px;
    border: solid gray;
  }
</style>
