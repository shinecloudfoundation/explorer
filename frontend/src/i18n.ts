import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

const messages = {
    'en': require('./locales/en-us.json'),
    'zhs': require('./locales/zh-hans.json'),
    'zht': require('./locales/zh-hant.json'),
}

const i18n = new VueI18n({
    locale: 'en',
    messages
})

export default i18n
