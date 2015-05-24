var ws;

function onLoad() {
	setStatus('Connecting...', false);
	doConnect();
}

function doConnect() {
	var proto = location.protocol === 'https:' ? 'wss' : 'ws';
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

function onKeyPress(event, func) {
	if (event.keyCode === 13) {
		func();
		return false;
	} else {
		return true;
	}
}

function doLogin() {
	var name = document.getElementById('txt-name').value.trim();
	if (name.length > 0) {
		var msg = { "command": "login", "name": name };
		ws.send(JSON.stringify(msg));
	}
}

function setDisplay(id, flag) {
	document.getElementById(id).style.display = flag ? 'inline' : 'none';
}

function setStatus(status, connected, msg) {
	document.getElementById('lbl-status').textContent = status;
	setDisplay('btn-connect', ! connected);
	setDisplay('btn-close', connected);

	setDisplay('div-login', msg && msg.status === 'login');
	setDisplay('div-players', msg && msg.status === 'idle');

	if (msg && msg.status === 'idle') {
		var list = document.getElementById('lst-players');
		while (list.firstChild) { list.removeChild(list.firstChild); }
		var cnt = 0;
		for (var i = 0; i < msg.players.length; i++) {
			var p = msg.players[i];
			if (p.id !== msg.id && p.status === 'idle') {
				var item = document.createElement('li');
				item.textContent = p.name;
				if (p.id == msg.opponent) { item.className = 'choice'; }
				if (p.opponent == msg.id) { item.className = 'chosen'; }
				item.onclick = onItemClick.bind(p);
				list.appendChild(item);
				cnt++;
			}
		}
		if (cnt == 0) {
			var item = document.createElement('li');
			item.textContent = 'no players...';
			list.appendChild(item);
		}
	}
}

function onItemClick() {
	var msg = { "command": "play", "opponent": this.id };
	ws.send(JSON.stringify(msg));
}

function onOpen(event) {
	setStatus('Please login.', true, { "status": "login" });
}

function onMessage(event) {
	var msg = JSON.parse(event.data);
	if (msg.status === 'login') {
		setStatus('Please login.', true, msg);
	} else if (msg.status === 'idle') {
		setStatus('Welcome, ' + msg.name + '!', true, msg);
	} else if (msg.status === 'play') {
		setStatus(msg.name + ' vs ' + msg.opponent, true, msg);
	} else {
		setStatus('Unknown status.', true);
	}
}

function onClose(event) {
	setStatus('Disconnected.', false);
}

function onError(event) {
	setStatus('ERROR!!!', false);
}
