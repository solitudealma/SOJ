import { createApp } from 'vue';
import App from '@/App.vue';

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';

import 'animate.css';
import 'font-awesome/css/font-awesome.min.css';

import '@/permission';

/**
 * @description 导入加载进度条，防止首屏加载时间过长，用户等待
 *
 * */
import Nprogress from 'nprogress';
import 'nprogress/nprogress.css';
Nprogress.configure({ showSpinner: false, easing: "ease", speed: 500 });
Nprogress.start();

import '@/assets/markdown.css';

import './monaco-worker';
import router from '@/router';
import { setupStore } from '@/pinia';



async function bootstrap() {
    const app = createApp(App);

    setupStore(app);

    app.use(ElementPlus).use(router).mount('#app');
}

bootstrap()
