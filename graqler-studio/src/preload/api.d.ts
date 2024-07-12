export interface Api {
  openFile(): Promise<string>

  ping(msg: string): void

  versions(): { [key: string]: string | undefined }
}

