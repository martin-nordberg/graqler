<script lang="ts">
  import Versions from './components/Versions.svelte'
  import electronLogo from './assets/electron.svg'
  import {pingMessage} from "../../../../graqler-core/src/corestuff";

  const sendPing = (): void => {
    window.api.ping(pingMessage)
  }

  const openFile = (): void => {
    window.api.openFile().then(fileName => {
      alert(fileName)
    })

  }
</script>

<img alt="logo" class="logo" src={electronLogo}/>
<div class="creator">Powered by electron-vite</div>
<div class="text">
  Build an Electron app with
  <span class="svelte">Svelte</span>
  and
  <span class="ts">TypeScript</span>
</div>
<p class="tip">Please press <code>F12</code> to open the devTool</p>
<div class="actions">
  <div class="action">
    <a href="https://electron-vite.org/" target="_blank" rel="noreferrer">Documentation</a>
  </div>
  <div class="action">
    <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions a11y-missing-attribute-->
    <a target="_blank" rel="noreferrer" on:click={sendPing}>Send IPC</a>
  </div>
  <div class="action">
    <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions a11y-missing-attribute-->
    <a target="_blank" rel="noreferrer" on:click={openFile}>Open File</a>
  </div>
  <div id="fileName"></div>
</div>
<Versions/>
