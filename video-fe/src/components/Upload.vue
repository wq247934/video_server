<template>
  <v-container>
    <v-row align="center">
      <v-form ref="form">
        <v-text-field
            v-model="video.title"
            :counter="30"
            label="视频标题"
            required
        ></v-text-field>
        <v-text-field
            v-model="video.introduction"
            label="视频简介"
        ></v-text-field>
        <v-file-input accept="video/*" label="File input" show-size @change="onFileChange"></v-file-input>
        <v-switch v-model="video.public" class="ma-4" label="是否公开"></v-switch>
        <v-btn @click="submit">提交</v-btn>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: "Upload",
  data: () => ({
    video:{
      file:"",
      title:"",
      introduction:"",
      public:""
    }
  }),
  methods: {
    onFileChange(file) {
      this.video.file = file;
    },
    submit() {
      let formData = new FormData();
      formData.set("file",this.video.file)
      formData.set("title",this.video.title)
      formData.set("introduction",this.video.introduction)
      formData.set("public",this.video.public)
      console.log(formData)
      axios.post("/api/video/upload", formData).then((resp) => {
        console.log(resp)
      }).catch((err) => {
        console.log(err)
      })
    },
  }
}
</script>

<style scoped>

</style>
