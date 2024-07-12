import {ElectronAPI} from '@electron-toolkit/preload'
import type {Api} from './api.d.ts'

declare global {
  interface Window {
    api: Api
  }
}
