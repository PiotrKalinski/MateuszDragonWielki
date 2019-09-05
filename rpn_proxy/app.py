from flask import Flask, jsonify, json
from calculator_service import calculate_expression

app = Flask(__name__)


@app.route('/rpn/', methods=['POST'])
def reverse_polish_notation():
    expression = calculate_expression()

    response = app.response_class(
        response= json.dumps([{'processed': k, 'expression': v} for k,v in expression.items()], indent=4),
        status=200,
        mimetype='application/json'
    )

    print(response)
    return response


if __name__ == '__main__':
    app.run()