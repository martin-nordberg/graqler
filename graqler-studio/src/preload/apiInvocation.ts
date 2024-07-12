
import {ipcRenderer} from 'electron'
import {electronAPI} from '@electron-toolkit/preload'

import type {Api} from './api'

// Custom APIs for renderer
export const api: Api = {
  openFile: async function (): Promise<string> {
    return ipcRenderer.invoke('dialog:openFile')
  },

  ping: function (msg: string): void {
    ipcRenderer.invoke('action:ping', msg)
  },

  versions: function (): { [key: string]: string | undefined } {
    return electronAPI.process.versions
  }
}
