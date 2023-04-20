from flask import Flask
from flask import render_template
from flask import render_template_string
from flask import request
import urllib
app = Flask(__name__)

@app.route('/')
def root():
    return render_template('index.html')
    
@app.route('/sib_coder')
def page_ric():
	return render_template('aerror.html')

@app.route('/robot.txt')
def page_robot():
	return render_template('robot.html')

	
@app.errorhandler(404)
def page_not_found(error):
    error_text = '''
    <h1>This page doesn't exist!</h1>
    <pre>%s</pre>
    ''' % (urllib.parse.unquote(request.url))

    return render_template_string(error_text)


if __name__ == '__main__':
    app.run(host='0.0.0.0')

