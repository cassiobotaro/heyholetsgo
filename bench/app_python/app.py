from flask import Flask, jsonify

app = Flask('app')


@app.route('/')
def hello():
    return jsonify({'hello': 'world'})


app.run(port='8080')
