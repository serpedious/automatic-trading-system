const environment = process.env.NODE_ENV
const envSettings = require(`./env.${environment}.js`)
module.exports = {
  env: envSettings,
  /*
  ** Headers of the page
  */
  head: {
    title: 'front',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'automatic-trading-front' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  buildModules: [
    '@nuxtjs/vuetify'
  ],
 
  build: {
    /*
    ** Run ESLint on save
    */
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/,
        })
      }
    }
  },
  
  generate: {
    routes: [
      '/auth/signin',
      '/auth/signup',
      '/dashboard',
      '/dashboard/chart',
      '/dashboard/memo',
      '/dashboard/csv',
      '/dashboard/alert'
    ]
  },
  
  server: {
    host: '0.0.0.0'
  }
}

