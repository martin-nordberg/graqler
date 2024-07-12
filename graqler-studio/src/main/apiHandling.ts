import {dialog} from "electron/main";
import {pongText} from "../../../graqler-core/src/corestuff";
import {IpcMainInvokeEvent, IpcMain} from "electron";

async function openFile(): Promise<string> {
  const {canceled, filePaths} = await dialog.showOpenDialog({})
  if (!canceled) {
    return filePaths[0]
  }
  return ""
}

export function handleApiCalls(ipcMain: IpcMain): void {
  ipcMain.handle('action:ping', (_: IpcMainInvokeEvent, msg: string) => console.log(pongText, msg))
  ipcMain.handle('dialog:openFile', () => openFile())
}
