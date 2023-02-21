import { createApp } from 'vue'
import App from '@/App.vue'

import ElementPlus from 'element-plus'

import router from '@/router'

async function bootstrap() {
    const app = createApp(App);
    app.use(ElementPlus, { size: 'small', zIndex: 3000 }).use(router).mount('#app')
}

bootstrap()
