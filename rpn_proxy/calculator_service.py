from datetime import datetime
import json
import time
from subprocess import Popen, PIPE

from flask import request


def calculate_expression():
    request_data = request.data.decode("utf-8")
    expressions = json.loads(request_data)['expression']

    results = []
    for single_expression in expressions.replace("\r", "").split("\n"):
        start = time.time()
        cmd = "ruby -r \"./calculator.rb\" -e \"RPNParser.parse '{0}'\"".format(single_expression)
        end = time.time() - start
        p = Popen(cmd, shell=True, stdout=PIPE)
        output, errors = p.communicate()
        request_data = output.decode("utf-8").rstrip('\n')
        results.append({"expression": single_expression, "data": request_data, "time": end})

    response = json.dumps(results, indent=4)

    with open('logs/test-{0}.txt'.format(datetime.now()), 'w') as fw:
        print(response)
         
        fw.write(response)

    return response