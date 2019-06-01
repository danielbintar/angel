extends Node2D

var username_form
var password_form
var flash_label
var last_hide_flash
var http_request
var config

const HIDE_FLASH_TIMER = 4
const HIDE_FLASH_MIN_TIME = 3800

var requesting = false

func _ready():
	flash_label = get_node("Flash")
	flash_label.hide()
	username_form = get_node("Username")
	password_form = get_node("Password")
	config = get_node("/root/Config")
	http_request = get_node("HTTPRequest")
	get_node("Register").connect("pressed", self, "register")
	http_request.connect("request_completed", self, "on_HTTPRequest_request_completed")

func register():
	if requesting:
		return

	show_flash("Requesting login...")
	requesting = true
	var data = {
		"username": username_form.get_text(),
		"password": password_form.get_text()
	}
	var headers = ["Content-Type: application/json"]
	var use_ssl = false

	http_request.request(config.API_URL + "users", headers, use_ssl, HTTPClient.METHOD_POST, JSON.print(data))

func on_HTTPRequest_request_completed(_result, response_code, _headers, body):
	requesting = false
	match response_code:
		200:
			show_flash("success")
		422:
			show_flash(body.get_string_from_utf8())
		_:
			show_flash("try again")

func show_flash(message):
	flash_label.text = message
	flash_label.show()
	last_hide_flash = OS.get_ticks_msec()
	yield(get_tree().create_timer(HIDE_FLASH_TIMER), "timeout")
	hide_flash()

func hide_flash():
	if OS.get_ticks_msec() - last_hide_flash >= HIDE_FLASH_MIN_TIME:
		flash_label.hide()
