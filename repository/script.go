package repository

var getCoins = `SELECT coins FROM users WHERE username = $1`

var getInventory = `SELECT merch, quantity FROM purchases WHERE username = $1`

var getReceivedHistory = `SELECT sent_from, money FROM transfers WHERE sent_to = $1`

var getSentHistory = `SELECT sent_to, money FROM transfers WHERE sent_from = $1`

var getUsername = "SELECT username FROM users WHERE username = $1"

var getPassword = "SELECT password FROM users WHERE password = $1"
