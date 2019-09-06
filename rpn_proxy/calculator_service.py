from datetime import datetime
import json
import time
from subprocess import Popen, PIPE

from flask import request

from ExpressionResult import ExpressionResult


def calculate_expression():
    start = time.clock()
    request_data = request.data.decode("utf-8")
    expression = json.loads(request_data)['expression']
    cmd = "ruby -r \"./calculator.rb\" -e \"RPNParser.parse '{0}'\"".format(expression)
    p = Popen(cmd, shell=True, stdout=PIPE)
    output, errors = p.communicate()
    request_data = format(output).decode("utf-8").rstrip('\n')
    fw = open('logs/test-{0}.txt'.format(datetime.now()), 'w')
    end = str(time.clock() - start)
    dad = dictionary = dict(zip(request_data.split(), expression.replace("\r","").split("\n")))
    try:
        fw.write(request_data + ', ' + end)
    except:
        fw.write("An error occured:")
    finally:
        fw.close()
    # res = ExpressionResult(request_data, end)
    return dad
