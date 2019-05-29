extends Node2D

var username_form
var password_form
var error_label
var last_hide_error
var http_request

const HIDE_ERROR_TIMER = 4
const HIDE_ERROR_MIN_TIME = 3800

var requesting = false

func _ready():
	username_form = get_node("Username")
	password_form = get_node("Password")
	error_label = get_node("ErrorLabel")
	error_label.hide()
	http_request = get_node("HTTPRequest")
	get_node("Register").connect("pressed", self, "register")
	http_request.connect("request_completed", self, "on_HTTPRequest_request_completed")

func register():
	if requesting:
		return

	requesting = true
	var data = {
		"username": username_form.get_text(),
		"password": password_form.get_text()
	}
	var headers = ["Content-Type: application/json"]
	var use_ssl = false

	http_request.request("http://api.local.host:8089/users", headers, use_ssl, HTTPClient.METHOD_POST, JSON.print(data))

func on_HTTPRequest_request_completed(_result, response_code, _headers, body):
	requesting = false
	match response_code:
		200:
			show_error("success")
		422:
			show_error(body.get_string_from_utf8())
		_:
			show_error("try again")

func show_error(error):
	error_label.text = error
	error_label.show()
	last_hide_error = OS.get_ticks_msec()
	yield(get_tree().create_timer(HIDE_ERROR_TIMER), "timeout")
	hide_error()

func hide_error():
	if OS.get_ticks_msec() - last_hide_error >= HIDE_ERROR_MIN_TIME:
		error_label.hide()
