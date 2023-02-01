import {UserConfigExport, ConfigEnv, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path'

function pathResolve(dir: string) {
    return resolve(process.cwd(), '.', dir)
}

// https://vitejs.dev/config/
export default ({command, mode}: ConfigEnv): UserConfigExport => {
    const env = loadEnv(mode, process.cwd())

    return {
        plugins: [vue()],
        resolve: {
            alias: [
                {
                    find: '@',
                    replacement: pathResolve('src') + '/',
                },
                {
                    find: '#',
                    replacement: pathResolve('types') + '/',
                },
            ],
        },
        server: {
            host: '0.0.0.0',
            hmr: true,
            port: 8000,
            proxy: {
                '/api': {
                    target: '127.0.0.1:8888',
                    changeOrigin: true,
                    secure: false,
                    rewrite: (path) => path.replace('', '/')
                },
            },
        },
    }
}
