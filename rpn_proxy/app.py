from datetime import datetime
from flask import Flask, request, jsonify
from subprocess import Popen, PIPE
import json


app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Hello World!'


@app.route('/rpn/', methods=['POST'])
def reverse_polish_notation():
    print(request.data)
    result = request.data.decode("utf-8")
    print("result" + result)
    expression = json.loads(result)['expression']
    print(expression)
    cmd = "ruby -r \"./calculator.rb\" -e \"RPNParser.parse '{0}'\"".format(expression)
    p = Popen(cmd, shell=True, stdout=PIPE)
    output, errors = p.communicate()
    result = output
    print(result)
    fw = open('logs/test-{0}.txt'.format(datetime.now()), 'w')
    fw.write('Result: ' + result + '')
    fw.write('Above should be result from performing ruby script')
    fw.close()
    return jsonify({
        'result': result
    })

if __name__ == '__main__':
    app.run()
