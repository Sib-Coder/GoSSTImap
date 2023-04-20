from flask import Flask
from flask import render_template
from flask import render_template_string
from flask import request
import urllib
app = Flask(__name__)

@app.route('/')
def root():
    return render_template('index.html')

@app.errorhandler(404)
def page_not_found(error):
    return render_template('aerror.html', url = urllib.parse.unquote(request.url))



if __name__ == '__main__':
    app.run(debug=True)

