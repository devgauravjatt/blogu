from flask import Flask
from flask import send_file
from flask import send_from_directory


app = Flask(__name__)


@app.route("/")
def index():
    return send_file("build/index.html")


@app.route("/")
def searching():
    return send_file("build/searching.html")


@app.route("/assets/<path:path>")
def send_asset(path):
    return send_from_directory("build/assets", path)


@app.route("/images/<path:path>")
def send_image(path):
    return send_from_directory("build/images", path)


@app.route("/blog/<title>")
def blogs(title):
    return send_file(f"build/blog/{title}.html")
