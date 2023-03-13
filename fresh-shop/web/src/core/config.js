/**
 * 网站配置文件
 */

const config = {
  appName: '启运冷鲜',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用 启运冷鲜后台管理系统！`
      )
    )
    console.log('\n')
  }
}

export default config
