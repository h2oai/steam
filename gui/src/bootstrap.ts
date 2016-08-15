/**
 * Created by justin on 8/14/16.
 */
const app =  require('electron').app;
const browserWindow = require('electron').BrowserWindow;

let win;

function createWindow() {
  win = new browserWindow({
    width: 600,
    height: 400
  });
  win.loadURL(`http://localhost:9000/index.html`);
  win.webContents.openDevTools();

  win.on('close', () => {
    win = null;
  });
}

app.on('ready', createWindow);

app.on('window-all-close', () => {
  if (process.platform !== 'darwin') {
    // app.;
  }
});

app.on('activate', () => {
  if (win === null) {
    createWindow();
  }
});

app.on('login', function(event, webContents, request, authInfo, callback) {
  // event.preventDefault();
  callback('superuser', 'superuser');
});
