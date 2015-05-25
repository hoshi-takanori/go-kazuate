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

function onItemClick() {
	var msg = { "command": "opponent", "opponent": this.id };
	ws.send(JSON.stringify(msg));
}

function doRandom() {
	var a = Math.floor(Math.random() * 10);
	var b, c;
	do {
		b = Math.floor(Math.random() * 10);
	} while (b == a);
	do {
		c = Math.floor(Math.random() * 10);
	} while (c == a || c == b);
	document.getElementById('txt-number').value = '' + a + b + c;
}

function doNumber() {
	var num = document.getElementById('txt-number').value.trim();
	if (num.length == 3) {
		var a = num.charCodeAt(0) - 48;
		var b = num.charCodeAt(1) - 48;
		var c = num.charCodeAt(2) - 48;
		if (a >= 0 && a < 10 && b >= 0 && b < 10 && c >= 0 && c < 10 &&
				a != b && b != c && c != a) {
			var msg = { "command": "number", "number": a * 100 + b * 10 + c };
			ws.send(JSON.stringify(msg));
		} else {
			alert('Duplicate digits.');
		}
	} else {
		alert('Bad length.');
	}
}

function setDisplay(id, flag) {
	document.getElementById(id).style.display = flag ? 'inline' : 'none';
}

function setEnabled(id, flag) {
	document.getElementById(id).disabled = ! flag;
}

function setStatus(status, connected, msg) {
	document.getElementById('lbl-status').textContent = status;
	setDisplay('btn-connect', ! connected);
	setDisplay('btn-close', connected);

	setDisplay('div-login', msg && msg.status === 'login');
	setDisplay('div-players', msg && msg.status === 'idle');
	setDisplay('div-number', msg && (msg.status === 'num1' || msg.status === 'num2'));
	setEnabled('txt-number', msg && msg.status === 'num1');
	setEnabled('btn-random', msg && msg.status === 'num1');
	setEnabled('btn-number', msg && msg.status === 'num1');
	setDisplay('div-play', msg && msg.status === 'play');

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

function onOpen(event) {
	setStatus('Please login.', true, { "status": "login" });
}

function onMessage(event) {
	var msg = JSON.parse(event.data);
	if (msg.status === 'login') {
		setStatus('Please login.', true, msg);
	} else if (msg.status === 'idle') {
		setStatus('Welcome, ' + msg.name + '!', true, msg);
	} else if (msg.status === 'num1' || msg.status === 'num2' || msg.status == 'play') {
		setStatus(msg.name + ' vs ' + msg.opp_name, true, msg);
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
