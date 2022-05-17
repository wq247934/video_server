<template>
  <v-container fill-height>
    <v-row justify="center" align="center" class="mt-12">
      <v-col cols="12">
        <v-row class="text-center px-3">
          <!-- <v-col cols="12">
            <v-img
              :src="require('../assets/yigrow-tools-logo.png')"
              class="my-3"
              contain
              height="80"
            ></v-img>
            <h2 style="font-family: fz">易种小助</h2>
          </v-col> -->
          <v-col cols="12" class="mt-6">
            <v-text-field
                v-model="LoginInfo.phone"
                label="账号"
                required
                color="#6cbdff"
                prepend-inner-icon="mdi-cellphone"
            ></v-text-field>
            <!-- :counter="7" -->
            <v-text-field
                v-model="LoginInfo.password"
                v-show="!loginStyle"
                label="密码"
                required
                color="#6cbdff"
                class="mt-4"
                prepend-inner-icon="mdi-lock"
                @keyup.enter="UserLogin()"
                type="password"
            ></v-text-field>
            <!-- 验证码登录 -->
            <v-row v-show="loginStyle">
              <v-col cols="7">
                <v-text-field
                    v-model="LoginInfo.authCode"
                    :counter="6"
                    label="验证码"
                    required
                    color="#6cbdff"
                    class="mt-4"
                    @keyup.enter="UserLogin()"
                ></v-text-field>
                <!-- prepend-inner-icon="mdi-message-processing-outline" -->
              </v-col>
              <v-col cols="5" class="mt-7">
                <v-btn
                    v-show="sendAuthCode"
                    block
                    tile
                    dark
                    @click="getAuthCode"
                    color="#6cbdff"
                    class="elevation-0"
                >
                  获取验证码
                </v-btn>
                <v-btn
                    v-show="!sendAuthCode"
                    dark
                    tile
                    block
                    color="#BDBDBD"
                    class="elevation-0"
                >
                  {{ authTime }}秒后重新获取
                </v-btn>
              </v-col>
            </v-row>

            <v-row class="justify-space-between">
              <v-col cols="6" md="4">
                <v-checkbox
                    v-model="remenberPassword"
                    :label="`${!loginStyle ? '记住密码' : '保持登录状态'} `"
                    color="#6cbdff"
                ></v-checkbox>
              </v-col>
              <v-col cols="6" md="4" class="mt-3 text-right">
                <v-btn text @click="loginStyle = !loginStyle">
                  <span dark v-show="!loginStyle" class="subtitle-1"
                  >验证码登录
                  </span>
                  <span dark v-show="loginStyle" class="subtitle-1"
                  >账号密码登录
                  </span>
                </v-btn>
              </v-col>
            </v-row>
            <v-btn
                block
                large
                dark
                color="#6cbdff"
                class="elevation-0"
                @click="UserLogin"
            >
              <span class="subtitle-1">登 录</span>
            </v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import axios from "axios";

export default {
  name: "Login",
  data: () => ({
    remenberPassword: false, //是否记住密码复选框,默认为false

    LoginInfo: {
      //用户登录信息
      phone: "", //手机号
      password: "", //密码
      authCode: "", //验证码
    },
    loginStyle: false, //登录方式
    sendAuthCode: true, //布尔值，通过v-show控制显示‘获取验证码’还是‘倒计时’
    authTime: 0 /*倒计时 计数器*/,
  }),
  created() {
  },
  mounted() {
  },
  methods: {
    /**
     * 登录
     * 验证手机号
     * 1.正则判断输入是否正确
     * 2.手机号登录框不能为空
     */
    UserLogin() {
      // console.log(this.$router)
      // this.$router.replace("/play")
      // let data = {};
      if (!this.loginStyle) {
        //账号密码登录
        if (this.LoginInfo.password === "") {
          this.$store.commit("showsnackbar", "密码不能为空！");
        } else {
          axios.post("/api/user_login/login", {
            "username": this.LoginInfo.phone,
            "password": this.LoginInfo.password
          }).then((resp) => {
            if (resp.data.errno !== 0) {
              this.$store.commit("showsnackbar", "用户名或密码错误！");
              console.log(resp.data)
            } else {
              this.userToken = "Bearer " + resp.data.data.token;
              // 将用户token保存到vuex中
              this.$store.commit("changeLogin", {
                Authorization: this.userToken,
                userNme: resp.data.data.username,
                // userID: resp.data.userID,
              });
              this.$router.replace("/play")
            }
          }).catch((err) => {
            console.log(err)
            this.$store.commit("showsnackbar", err);

          })
        }
      } else {
        //验证码登录
        if (
            this.LoginInfo.authCode.length < 6 ||
            this.LoginInfo.authCode.length > 6
        ) {
          this.$store.commit("showsnackbar", "验证码长度不符合！");
        } else {
          console.log("验证码登录");
          this.$router.replace('/')
        }
      }
    },
    //获取验证码
    getAuthCode() {
      let _this = this;
      let re = new RegExp(
          /^((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)$/
      );
      let retu = this.LoginInfo.phone.match(re);
      if (this.LoginInfo.phone == "") {
        _this.$store.commit("showsnackbar", "手机号不能为空！");
      } else if (!retu) {
        _this.$store.commit("showsnackbar", "输入的手机号有误，请重新输入");
      } else {
        //判断验证码倒计时
        this.sendAuthCode = false;
        this.authTime = 60;
        var authTimer = setInterval(() => {
          this.authTime--;
          if (this.authTime <= 0) {
            this.sendAuthCode = true;
            clearInterval(authTimer); //停止计时器
          }
        }, 1000);
      }
    },
  },
};
</script>
