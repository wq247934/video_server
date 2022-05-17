const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  publicPath: '/',
  outputDir: 'dist',
  devServer: {
    open: false,
    host: 'localhost',
    port: '8080',
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:9001', // 要请求的地址
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
})
