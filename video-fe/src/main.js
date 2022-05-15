import Vue from 'vue'//导入vue
import App from './App.vue'//导入app根组件
import vuetify from './plugins/vuetify'//导入vuetify
import store from './store'//导入store.js文件
import router from './router'  //引入vue-router
import VideoPlayer from 'vue-video-player'
import axios from "axios";
import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)
require('video.js/dist/video-js.css')
require('vue-video-player/src/custom-theme.css')
Vue.use(VideoPlayer)

Vue.config.productionTip = false
console.log(router)
new Vue({
  vuetify,
  store,//将store挂载到new出来的Vue实例中，这样在vue中的每个组件都可以访问到全局共享数据了
  router,
  render: h => h(App)
}).$mount('#app')

