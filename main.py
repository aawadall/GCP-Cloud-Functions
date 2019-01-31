import base64


def hello_pubsub(event, context):
    """Triggered from a message on a Cloud Pub/Sub topic.
    Args:
         event (dict): Event payload.
         context (google.cloud.functions.Context): Metadata for the event.
    """
    msg_details = event['attributes']

    pubsub_message = base64.b64decode(event['data']).decode('utf-8')
    print(pubsub_message)
    print("Message Details:")
    for det in msg_details:
        print(det.decode('utf-8'))

