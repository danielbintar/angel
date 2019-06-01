extends Node

var API_URL

func _ready():
	var result = {}
	var f = File.new()
	f.open(".env", File.READ)
	while not f.eof_reached():
		var line = f.get_line()
		var k = line.split("=")
		if k.size() == 2:
			result[k[0]] = k[1]

	f.close()
	API_URL = result["API_URL"]
