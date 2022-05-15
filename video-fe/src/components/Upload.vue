<template>
<v-container>
  <v-app>
<v-form>
      <v-file-input accept="video/*" label="File input" show-size @change="onFileChange"></v-file-input>
      <v-btn @click="submit"></v-btn>
</v-form>
  </v-app>
</v-container>
</template>

<script>
import axios from "axios";
export default {
  name: "Upload",
  data:()=>({
    file:""
}),
  methods:{
    onFileChange(file){
      this.file = file;
    },
    submit(){
      let formData = new FormData();
      formData.append("file",this.file)
      axios.post("http://localhost:9000/video/2",formData).then((resp)=>{
        console.log(resp)
      }).catch((err)=>{
        console.log(err)
      })
    }
  }
}
</script>

<style scoped>

</style>
