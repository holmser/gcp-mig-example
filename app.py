from flask import Flask
import requests

app = Flask(__name__)


@app.route("/")
def get_default():
    url = 'http://metadata.google.internal/computeMetadata/v1'
    headers = {'Metadata-Flavor': 'Google'}
    default_info = ['/instance/hostname', '/instance/machine-type', '/instance/zone']
    return_text = ''
    for subpath in default_info:
        info = subpath.split("/")[1]
        new_url = url + subpath
        r = requests.get(new_url, headers=headers)
        return_text += (info + ": " + r.text + "\n")
    return(return_text)


@app.route("/<path:subpath>")
def get_info(subpath):
    url = 'http://metadata.google.internal/computeMetadata/v1/'
    headers = {'Metadata-Flavor': 'Google'}
    url += subpath
    r = requests.get(url, headers=headers)
    return(r.text)