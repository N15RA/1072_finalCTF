import flask
import os


app = flask.Flask(__name__)
app.config['FLAG'] = os.environ.pop('FLAG')
app.config['FLASK_DEEPER'] = os.environ.pop('FLASK_DEEPER')

@app.route('/')
def index():
    return flask.render_template('index.html', content=open(__file__).read())

@app.route('/dalun/<path:dalun>')
def dalun(dalun):
    def waf(s):
        s = s.replace('(', '').replace(')', '')
        blacklist = ['config', 'self','url_for','request','session','get_flashed_messages']
        return ''.join(['{{% set {}=None%}}'.format(c) for c in blacklist])+s
    return flask.render_template_string(waf(dalun))

if __name__ == '__main__':
    app.run(debug=True)
