import { defineStore } from 'pinia'
import { store } from '@/pinia';
import type { DialogStatus } from '#/store'

interface GlobalState {
    dialogStatus: DialogStatus
}

export const useGlobalStore = defineStore('global', {
    state: (): GlobalState => ({
        dialogStatus: {
            mode: 'Login',
            visible: false,
        }
    }),
    getters: {},
    actions: {
        changeDialogStatus(dialogStatus: DialogStatus) {
            console.log(dialogStatus)
            if (dialogStatus.mode !== undefined) {
                this.dialogStatus.mode = dialogStatus.mode
            }
            if (dialogStatus.visible !== undefined) {
                this.dialogStatus.visible = dialogStatus.visible
            }
        },
    },
})

// Need to be used outside the setup
export function useGlobalStoreWithOut() {
  return useGlobalStore(store);
}