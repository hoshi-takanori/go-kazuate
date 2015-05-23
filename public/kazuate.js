var ws;

function onLoad() {
	setStatus('NOT CONNECTED', false);
}

function doConnect() {
	var proto = location.protocol == 'https:' ? 'wss' : 'ws';
	var index = location.pathname.lastIndexOf('/');
	var path = index > 0 ? location.pathname.substring(0, index + 1) : '/';
	var url = proto + '://' + location.host + path + 'ws';
	ws = new WebSocket(url);
	ws.onopen = onOpen;
	ws.onmessage = onMessage;
	ws.onclose = onClose;
	ws.onerror = onError;
}

function doClose() {
	ws.close();
}

function setStatus(status, connected) {
	document.getElementById('lbl-status').textContent = status;
	document.getElementById('btn-connect').style.display = connected ? 'none' : 'inline';
	document.getElementById('btn-close').style.display = connected ? 'inline' : 'none';
}

function onOpen(event) {
	setStatus('Connected', true);
	ws.send("connected");
}

function onMessage(event) {
	alert('onmessage');
}

function onClose(event) {
	setStatus('Closed', false);
}

function onError(event) {
	alert('onerror');
}
