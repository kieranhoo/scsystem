from flask_cors import CORS
import os
from dotenv import load_dotenv
from flask import Flask

load_dotenv()
FLASK_ENV = os.getenv("FLASK_ENV")


def create_app(config=None):
    flask_app = Flask(__name__, instance_relative_config=True)
    CORS(flask_app)
    return flask_app


app = create_app()

__all__ = ("app",)
