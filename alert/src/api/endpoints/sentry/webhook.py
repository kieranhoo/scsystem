from __future__ import annotations

from typing import Any, Mapping

from src import app
from src.api.middleware import verify_sentry_signature
# from src.models import SentryInstallation, Organization

from flask import request, Response


@app.route("/api/sentry/webhook/", methods=["POST"])
# @verify_sentry_signature()
def webhook_index():
    # Parse the JSON body fields off of the request
    action = request.json.get("action")
    data = request.json.get("data")
    actor = request.json.get("actor")
    uuid = request.json.get("installation", {}).get("uuid")
    # Identify the resource triggering the webhook in Sentry
    resource = request.headers.get("sentry-hook-resource")
    if not (action and data and uuid and resource):
        return Response("no data", 400)
    app.logger.info(f"Received '{resource}.{action}' webhook from Sentry")

    # We can monitor what status codes we're sending from the integration dashboard
    return Response("", 200)
