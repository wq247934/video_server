<template>
  <v-app>
    <v-snackbar
        v-model="promptMessage"
        top
        :timeout="3000"
        style="position: fixed; top: 20px"
    >
      {{ snackbarInfo }}

      <template v-slot:action="{ attrs }">
        <v-btn icon v-bind="attrs" @click="promptMessage = false">
          <v-icon>mdi-close-circle</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld';
import Login from "./components/Login.vue";

export default {
  name: "App",

  components: {
    // HelloWorld,
    Login,
  },
  computed: {
    //设置公共消息提示
    snackbarInfo() {
      // this.$store.state.全局数据名称，组件访问state中数据的第一种方式
      return this.$store.state.snackbarInfo;
    },
    promptMessage: {
      get() {
        return this.$store.state.promptMessage;
      },
      set() {
        this.$store.state.promptMessage = false;
        this.$store.state.num = 0;
        // clearTimeout() 方法可取消由 setTimeout() 方法设置的 timeout。
        clearTimeout(this.$store.state.time);
      },
    },
  },

  data: () => ({
    //
  }),
};
</script>

