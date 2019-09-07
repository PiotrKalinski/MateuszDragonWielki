from flask import Flask
from calculator_service import calculate_expression

app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Flask Dockerized'


@app.route('/rpn/', methods=['POST'])
def reverse_polish_notation():
    return app.response_class(
        response=calculate_expression(),
        status=200,
        mimetype='application/json'
    )


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')