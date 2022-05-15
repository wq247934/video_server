<template>
  <v-container>
    <router-link to="/upload">上传视频</router-link>
    <v-app>
<!--      <v-overlay :absolute="false" :value="videoModal">-->
<!--        <v-icon-->
<!--            @click="videoModal=false"-->
<!--            style="font-size:30px;position: absolute;top:-15px;right:-15px;z-index:2;cursor:pointer;"-->
<!--        >mdi-close-circle-outline</v-icon>-->
        <div class="input_video">
          <video-player
              id = "video"
              class="video-player vjs-custom-skin"
              ref="videoPlayer"
              :playsinline="true"
              :options="playerOptions"
              @play="onPlayerPlay($event)"
              @ready="playerIsReady"
          ></video-player>
        </div>
<!--      </v-overlay>-->
    </v-app>
  </v-container>
</template>

<script>
import 'videojs-hotkeys'
  export default {
    name: 'HelloWorld',
    data: () => ({
      videoModal:true,
      playerOptions: {
        playbackRates: [0.5, 1.0, 1.5, 2.0],//倍速控制
        autoplay: false,//是否自动播放
        muted: false,//是否静音播放
        loop: false,//是否循环播放
        preload: "auto",
        language: "zh-CN",
        aspectRatio: "16:9",//比例
        fluid: true,
        sources: [
          {
            type: "video/mp4",
            src:
                "http://localhost:9000/video/1", //url地址
          },
        ],
        poster: "", //你的封面地址
        notSupportedMessage: "此视频暂无法播放，请稍后再试",
        controlBar: {
          timeDivider: true,
          durationDisplay: true,//剩余时间是否显示
          remainingTimeDisplay: true,//剩余时间是否显示，有一个即可
          fullscreenToggle: true,//全屏按钮
        },
      },
    }),
    methods:{
      // listen event
      onPlayerPlay(player) {
        var video = document.getElementById('video');
        console.log('player play!', video.playbackRate)
      },
      playerIsReady(player) {
        console.log('example 2 ready!', player)
        player.hotkeys({
          volumeStep: 0.1,
          seekStep: 5,
          enableModifiersForNumbers: false,
          fullscreenKey: function(event, player) {
            // override fullscreen to trigger when pressing the F key or Ctrl+Enter
            return ((event.which === 70) || (event.ctrlKey && event.which === 13));
          }
        })
      }
    }
  }
</script>
