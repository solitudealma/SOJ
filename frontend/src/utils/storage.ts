const localStorage = window.localStorage

export default {
    name: 'storage',

    /**
     * save value(Object) to key
     * @param {string} key 键
     * @param {Object} value 值
     */
    set(key: string, value: any) {
        localStorage.setItem(key, JSON.stringify(value))
    },

    /**
     * get value(Object) by key
     * @param string key 键
     * @returns object | string
     */
    get(key: string) {
        let value = localStorage.getItem(key)

        if(value == null) {
            return null
        } else if(typeof value === 'string') {
            return JSON.parse(value)
        }
    },

    /**
     * remove key from localStorage
     * @param {string} key 键
     */
    remove(key: string) {
        localStorage.removeItem(key)
    },
    /**
     * clear all
     */
    clear() {
        localStorage.clear()
    },
}
