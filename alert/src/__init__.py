from .server import app

# Register routes and serializers.
with app.app_context():
    from .api import endpoints  # NOQA

__all__ = ("app",)
