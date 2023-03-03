import { defineStore } from 'pinia'
import { store } from '@/pinia';
import type { ModalStatus } from '#/store'

interface GlobalState {
    modalStatus: {
        mode: string 
        visible: boolean
    }
}

export const useGlobalStore = defineStore('global', {
    state: (): GlobalState => ({
        modalStatus: {
            mode: 'LoginComponent',
            visible: false,
        }
    }),
    getters: {},
    actions: {
        changeModalStatus(modalStatus: ModalStatus) {
            if (modalStatus.mode !== undefined) {
                this.modalStatus.mode = modalStatus.mode
            }
            if (modalStatus.visible !== undefined) {
                this.modalStatus.visible = modalStatus.visible
            }
        },
    },
})

// Need to be used outside the setup
export function useGlobalStoreWithOut() {
  return useGlobalStore(store);
}
