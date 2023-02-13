const { defineConfig } = require('@vue/cli-service')
module.exports = ({
  lintOnSave:false,//关闭语法检查
  devServer:{
        host: 'localhost', // 监听地址
        port: '7070', // 启动端口号
        open: true // 启动后是否自动打开网页
    },
})
