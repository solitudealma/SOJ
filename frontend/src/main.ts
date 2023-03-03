import { createApp } from 'vue'
import App from '@/App.vue'

import ElementPlus from 'element-plus'

import router from '@/router'
import { setupStore } from '@/pinia'

async function bootstrap() {
    const app = createApp(App);

    setupStore(app);

    app.use(ElementPlus).use(router).mount('#app')
}

bootstrap()
